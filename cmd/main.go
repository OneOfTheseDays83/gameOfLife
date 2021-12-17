package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gol/api"
	"gol/cmd/core"
	"gol/cmd/data"
	"gol/cmd/publish"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var publisher publish.Publisher

// init is the reserved golang function that will initialize all components once.
func init() {
	publisher = publish.NewConsolePublisher()
}

func main() {
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, os.Interrupt)
	res := make(chan error, 1)
	defer close(res)

	port := os.Getenv("SERVICE_PORT")
	log.Printf("Listening in port %s", port)

	s := http.Server{
		Addr:    ":" + port,
		Handler: createRootHandler(),
	}
	go func() {
		res <- s.ListenAndServe()
	}()

	select {
	case <-quit:
		log.Println("user initiated termination of server started")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := s.Shutdown(ctx)
		if err != nil {
			log.Println("graceful shutdown failed: ", err)
		}
	case err := <-res:
		log.Println("server stopped with error: ", err)
	}
}

func createRootHandler() http.Handler {
	r := mux.NewRouter()
	restApi := r.PathPrefix("/v1").Subrouter()
	restApi.HandleFunc("/gol", playTheGame).Methods(http.MethodPost)
	return r
}

func playTheGame(rw http.ResponseWriter, r *http.Request) {
	var request api.Request
	if err := decode(r, &request); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	var grid data.Grid
	var err error

	if len(request.Grid) > 0 {
		// user provided a grid
		grid, err = data.CreateFromContent(request.Grid)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if request.Rows > 0 && request.Cols > 0 {
		// user just gave rows and cols -> rand start
		grid, err = data.NewGrid(request.Rows, request.Cols)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		grid.Random()
	} else {
		// user didn't provide anything stop here
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO currently we use only console as publisher
	// but we can easily create a new publisher that constructs a http response and inject it here
	gol := core.NewGameOfLife(publisher, grid)

	for i := uint64(0); i < request.Iterations; i++ {
		gol.Continue()
	}

	rw.WriteHeader(http.StatusOK)
	return
}

func decode(r *http.Request, data interface{}) (err error) {
	if data == nil {
		return errors.New("is nil")
	}
	err = json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Println("json decoding failed: ", err.Error())
	}
	return
}
