# Game of life
This implements the game of life as a backend service in GO. The requests to play the game can be send via REST API. The output/response is currently limited to console logging.

---
## Architecture
The application is set up as a microservice with a REST API to play the game. You can build it with your toolchain or use the docker variant.

## Build
### local
If you didn't change anything you don't need to build the application and can rather jump to "Start the service".

Download the dependencies first. This will download all the needed go modules needed to build the service.
```shell
make download-deps 
```
Now build the service.
```shell
make build-local
```

### Docker
```shell
make build
```

## Play Game of Life

### Start the service
If you build the service with your toolchain:
```shell
make start-local
```
If you used the docker variant:
```shell
make start
```

### Requests to play the game
The request are sent via the REST API of the service. Currently there is only one endpoint: `http://localhost:8000/v1/gol`.

The Request to play the game is a post request with the following JSON body:
```json
"iterations": 5,
"grid": [[]],
"rows": 2,
"cols": 5
```
- `iterations`: number of iterations (generations) to run the game
- `grid`: a matrix can be provided as a starting grid. `false` is for dead, `true` for alive. If not provided a random is used.
- `rows`: number of rows (can be empty if grid is given)
- `cols`: number of cols (can be empty if grid is given)

#### Example request with a grid
```shell
curl -v -X POST http://localhost:8000/v1/gol --data-raw '
{
    "iterations": 5,
    "grid": [
        [false, true, false, false, false, true, false, false],
        [false, false, false, true, false, false, false, false],
        [false, true, false, false, false, false, false, false],
        [false, false, false, true, false, false, false, false],
        [false, true, false, false, false, false, false, false]
     ]
}'
```
#### Example request with a random grid
```shell
curl -v -X POST http://localhost:8000/v1/gol --data-raw '
{
    "iterations": 5,
    "rows": 5,
    "cols": 8
}'
```

## Open points
* OpenApi spec missing in ./api
* Unit testing must be extended
* Improve grid json input to not be limited to `true` `false`  
* Return results of game via REST API (either one result for all generations or websocket)