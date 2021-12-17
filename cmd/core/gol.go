package core

import (
	"gol/cmd/data"
	"gol/cmd/publish"
	"log"
)

func NewGameOfLife(pub publish.Publisher, initialGrid data.Grid) Game {
	if pub != nil {
		pub.Print(initialGrid, 0)
	}
	return Game{pub: pub, currentGrid: initialGrid}
}

type Game struct {
	pub         publish.Publisher
	currentGrid data.Grid
	iteration   uint64
}

func (g *Game) Continue() {
	g.iteration++
	rows, cols := g.currentGrid.Size()

	// create the grid for the new iteration
	res, err := data.NewGrid(rows, cols)
	if err != nil {
		log.Printf("aborting Continue do to error: %s", err.Error())
		return
	}

	// loop through each cell in the matrix
	for r := uint64(0); r < rows; r++ {
		for c := uint64(0); c < cols; c++ {
			res.Set(r, c, g.isAlife(r, c))
		}
	}

	// store the new iteration as current one
	g.currentGrid = res

	if g.pub != nil {
		g.pub.Print(g.currentGrid, g.iteration)
	}
}

func (g *Game) isAlife(row, col uint64) bool {
	alife, err := g.currentGrid.Get(row, col)
	if err != nil {
		log.Printf("aborting isAlife due to error: %s", err.Error())
	}

	neighbours := g.countOfAliveNeighbours(row, col)

	// Any life cell with fewer than two life neighbours dies, as if caused by underpopulation
	if alife && neighbours < 2 {
		return false
	}

	// Any life cell with more than three life neighbours dies, as if by overcrowding.
	if alife && neighbours > 3 {
		return false
	}

	// Any life cell with two or three life neighbours lives on to the next generation.
	if alife && neighbours >= 2 && neighbours <= 3 {
		return true
	}

	// Any dead cell with exactly three life neighbours becomes a live cell.
	if !alife && neighbours == 3 {
		return true
	}

	return false
}

func (g *Game) countOfAliveNeighbours(row, col uint64) (res uint64) {
	indexOffset := []int{-1, 0, 1}

	for _, r := range indexOffset {
		for _, c := range indexOffset {

			if (int(row)+r) < 0 || (int(col)+c) < 0 {
				// we are at the upper and/or left edge
				continue
			}

			if r == 0 && c == 0 {
				// this is the point itself
				continue
			}

			alife, err := g.currentGrid.Get(uint64(int(row)+r), uint64(int(col)+c))
			if err != nil {
				continue
			}

			if alife {
				res++
			}
		}
	}
	return
}
