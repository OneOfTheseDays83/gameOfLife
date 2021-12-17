package publish

import (
	"fmt"
	"gol/cmd/data"
	"log"
)

var deadOrAliveToString = map[bool]string{
	true:  fmt.Sprintf("%2s", "o"),
	false: fmt.Sprintf("%2s", "-"),
}

func NewConsolePublisher() Publisher {
	return &console{}
}

type console struct{}

func (c *console) Print(grid data.Grid, iteration uint64) {

	log.Printf("\n\n Iteration: %d", iteration)
	rows, _ := grid.Size()

	for r := uint64(0); r < rows; r++ {
		row, err := grid.GetRow(r)
		if err == nil {
			log.Println(getRowOutput(row))
		}
	}
}

func getRowOutput(row []bool) (out []string) {
	for _, alife := range row {
		out = append(out, deadOrAliveToString[alife])
	}

	return
}
