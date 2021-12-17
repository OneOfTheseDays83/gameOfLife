package core

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gol/cmd/data"
	"gol/cmd/publish"
	"testing"
)

func TestGame_countOfAliveNeighbours(t *testing.T) {
	t.Run("upper left corner", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 4)
		assert.Nil(t, err)
		testGrid.Set(0, 1, true)
		testGrid.Set(1, 0, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, uint64(2), toTest.countOfAliveNeighbours(0, 0))
	})

	t.Run("upper right corner", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 4)
		assert.Nil(t, err)
		testGrid.Set(0, 2, true)
		testGrid.Set(1, 2, true)
		testGrid.Set(1, 3, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, uint64(3), toTest.countOfAliveNeighbours(0, 3))
	})

	t.Run("bottom left corner", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 4)
		assert.Nil(t, err)
		testGrid.Set(2, 1, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, uint64(1), toTest.countOfAliveNeighbours(3, 0))
	})

	t.Run("bottom right corner", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 4)
		assert.Nil(t, err)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, uint64(0), toTest.countOfAliveNeighbours(3, 3))
	})

	t.Run("index out of bounds", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 6)
		assert.Nil(t, err)
		testGrid.Set(0, 3, true)
		testGrid.Set(1, 2, true)
		testGrid.Set(1, 3, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, uint64(0), toTest.countOfAliveNeighbours(0, 7))
	})
}

func TestGame_isAlife(t *testing.T) {
	t.Run("underpopulation", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 6)
		assert.Nil(t, err)
		testGrid.Set(0, 1, true)
		testGrid.Set(1, 0, true)
		testGrid.Set(1, 1, true)
		testGrid.Set(1, 3, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, false, toTest.isAlife(1, 3))
	})
	t.Run("overcrowding", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 6)
		assert.Nil(t, err)
		testGrid.Set(0, 1, true)
		testGrid.Set(1, 0, true)
		testGrid.Set(1, 1, true)
		testGrid.Set(1, 2, true)
		testGrid.Set(2, 1, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, false, toTest.isAlife(1, 1))
	})
	t.Run("continues to life with 2", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 6)
		assert.Nil(t, err)
		testGrid.Set(0, 1, true)
		testGrid.Set(1, 0, true)
		testGrid.Set(1, 1, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, true, toTest.isAlife(1, 1))
	})
	t.Run("continues to life with 3", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 6)
		assert.Nil(t, err)
		testGrid.Set(0, 1, true)
		testGrid.Set(1, 0, true)
		testGrid.Set(1, 1, true)
		testGrid.Set(2, 1, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, true, toTest.isAlife(1, 1))
	})
	t.Run("dead to alive with 3", func(t *testing.T) {
		testGrid, err := data.NewGrid(4, 6)
		assert.Nil(t, err)
		testGrid.Set(0, 1, true)
		testGrid.Set(1, 0, true)
		testGrid.Set(2, 1, true)
		toTest := NewGameOfLife(nil, testGrid)
		assert.Equal(t, true, toTest.isAlife(1, 1))
	})
}

func TestGame_Continue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("test pattern 1", func(t *testing.T) {
		publisherMock := publish.NewMockPublisher(ctrl)
		testGrid, err := data.NewGrid(6, 6)
		assert.Nil(t, err)
		testGrid.Set(1, 1, true)
		testGrid.Set(2, 1, true)
		testGrid.Set(3, 2, true)

		publisherMock.EXPECT().Print(gomock.Any(), uint64(0))
		toTest := NewGameOfLife(publisherMock, testGrid)

		resultGrid, err := data.NewGrid(6, 6)
		assert.Nil(t, err)
		resultGrid.Set(2, 1, true)
		resultGrid.Set(2, 2, true)

		publisherMock.EXPECT().Print(resultGrid, uint64(1))
		toTest.Continue()
	})

	t.Run("test pattern 2", func(t *testing.T) {
		publisherMock := publish.NewMockPublisher(ctrl)
		testGrid, err := data.NewGrid(6, 6)
		assert.Nil(t, err)
		testGrid.Set(1, 3, true)
		testGrid.Set(2, 2, true)
		testGrid.Set(3, 1, true)

		publisherMock.EXPECT().Print(gomock.Any(), uint64(0))
		toTest := NewGameOfLife(publisherMock, testGrid)

		resultGrid, err := data.NewGrid(6, 6)
		assert.Nil(t, err)
		resultGrid.Set(2, 2, true)

		publisherMock.EXPECT().Print(resultGrid, uint64(1))
		toTest.Continue()
	})
}
