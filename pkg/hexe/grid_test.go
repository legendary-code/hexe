package hexe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGridGetSet(t *testing.T) {
	g := NewGrid[string]()
	c := g.Cube()

	s, ok := c.Index(1, 0, -1)
	assert.False(t, ok)
	assert.Equal(t, "", s)

	c.Set(1, 0, -1, "foo")
	s, ok = c.Index(1, 0, -1)
	assert.True(t, ok)
	assert.Equal(t, "foo", s)

	s = c.Get(1, 0, -1)
	assert.Equal(t, "foo", s)

	a := c.Axial()
	s = a.Get(1, 0)
	assert.Equal(t, "foo", s)

	a.Delete(1, 0)
	s, ok = a.Index(1, 0)
	assert.False(t, ok)
	assert.Equal(t, "", s)
}

func TestGrid_Neighbors(t *testing.T) {
	g := NewGrid[string]()
	g.Set(0, 0, "foo")
	g.Set(0, 1, "bar")
	g.Set(1, 0, "baz")
	g.Set(1, -1, "qux")
	g.Set(-1, -1, "quux")
	g.Set(2, 2, "quuux")

	neighbors := g.Neighbors(0, 0)
	names := make([]string, 0)
	coords := make([][]int, 0)
	for _, neighbor := range neighbors {
		names = append(names, neighbor.Value())

		q, r := neighbor.Index().Unpack()
		coords = append(coords, []int{q, r})
	}

	assert.Equal(
		t,
		[]string{"baz", "bar", "qux"},
		names,
	)

	assert.Equal(
		t,
		[][]int{
			{1, 0},
			{0, 1},
			{1, -1},
		},
		coords,
	)
}

func TestGrid_DiagonalNeighbors(t *testing.T) {
	g := NewGrid[string]()
	g.Set(0, 0, "foo")
	g.Set(-1, -1, "bar")
	g.Set(-1, 2, "baz")
	g.Set(2, -1, "qux")
	g.Set(-1, -2, "quux")
	g.Set(2, 2, "quuux")

	neighbors := g.DiagonalNeighbors(0, 0)
	names := make([]string, 0)
	coords := make([][]int, 0)
	for _, neighbor := range neighbors {
		names = append(names, neighbor.Value())

		q, r := neighbor.Index().Unpack()
		coords = append(coords, []int{q, r})
	}

	assert.Equal(
		t,
		[]string{"baz", "bar", "qux"},
		names,
	)

	assert.Equal(
		t,
		[][]int{
			{-1, 2},
			{-1, -1},
			{2, -1},
		},
		coords,
	)
}
