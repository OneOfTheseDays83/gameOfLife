package data

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

type Grid struct {
	content [][]bool
}

func NewGrid(rows, cols uint64) (Grid, error) {
	if (rows == 0) || (cols == 0) {
		return Grid{}, errors.New("rows or cols is 0")
	}
	content := make([][]bool, rows)
	for i := 0; i < int(rows); i++ {
		content[i] = make([]bool, cols)
	}

	return Grid{content: content}, nil
}

func CreateFromContent(content [][]bool) (Grid, error) {
	return Grid{content: content}, nil
}

func (g *Grid) Get(row, col uint64) (entry bool, err error) {
	if !g.IsValid(row, col) {
		err = errors.New("row, col invalid")
		return
	}

	entry = g.content[row][col]
	return
}

func (g *Grid) GetRow(row uint64) (entry []bool, err error) {
	if !g.IsValid(row, 0) {
		err = errors.New("row, col invalid")
		return
	}

	entry = g.content[row]
	return
}

func (g *Grid) Set(row, col uint64, val bool) {
	if !g.IsValid(row, col) {
		log.Printf("tried to set value for index out of bounds, row %d, col %d", row, col)
		return
	}

	g.content[row][col] = val
	return
}

func (g *Grid) Size() (rows, cols uint64) {
	rows = uint64(len(g.content))

	if rows > 1 {
		cols = uint64(len(g.content[0]))
	}
	return
}

func (g *Grid) IsValid(row, col uint64) bool {
	rows, cols := g.Size()
	return (row < rows) && (col < cols)
}

func (g *Grid) Random() {
	rand.Seed(time.Now().UnixNano())

	rows, cols := g.Size()
	for r := uint64(0); r < rows; r++ {
		for c := uint64(0); c < cols; c++ {
			g.Set(r, c, rand.Float32() < 0.5)
		}
	}
}
