package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid_Size(t *testing.T) {
	t.Run("size wrong", func(t *testing.T) {
		rows := uint64(100)
		cols := uint64(50)
		grid, err := NewGrid(rows, cols)
		assert.Nil(t, err)
		sizeRows, sizeCols := grid.Size()
		assert.Equal(t, rows, sizeRows)
		assert.Equal(t, cols, sizeCols)
	})

	t.Run("size 0", func(t *testing.T) {
		_, err := NewGrid(0, 20)
		assert.NotNil(t, err)
	})
}

func TestGrid_IsValid(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		rows := uint64(100)
		cols := uint64(50)
		grid, err := NewGrid(rows, cols)
		assert.Nil(t, err)
		assert.True(t, grid.IsValid(50, 20))
	})

	t.Run("not valid", func(t *testing.T) {
		rows := uint64(100)
		cols := uint64(50)
		grid, err := NewGrid(rows, cols)
		assert.Nil(t, err)
		assert.False(t, grid.IsValid(rows, 20))
	})
}

func TestGrid_Set(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		rows := uint64(100)
		cols := uint64(50)
		grid, err := NewGrid(rows, cols)
		assert.Nil(t, err)
		grid.Set(10, 20, true)
		res, err := grid.Get(10, 20)
		assert.Nil(t, err)
		assert.True(t, res)
	})

	t.Run("index out of bounds", func(t *testing.T) {
		rows := uint64(100)
		cols := uint64(50)
		grid, err := NewGrid(rows, cols)
		assert.Nil(t, err)
		grid.Set(rows, 20, true)
		res, err := grid.Get(10, 20)
		assert.Nil(t, err)
		assert.False(t, res)
	})
}
