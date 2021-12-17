package publish

import "gol/cmd/data"

type Publisher interface {
	// Print outputs the current iteration
	Print(grid data.Grid, iteration uint64)
}
