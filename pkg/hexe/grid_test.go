package hexe

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
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

	neighbors := g.IndexCoords(coord.NewAxial(0, 0).Neighbors())
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

	neighbors := g.IndexCoords(coord.NewAxial(0, 0).DiagonalNeighbors())
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

func TestGrid_IndexLine(t *testing.T) {
	g := NewGrid[string]()
	g.Set(0, 0, "foo")
	g.Set(-1, -1, "bar")
	g.Set(-1, 2, "baz")
	g.Set(2, -1, "qux")
	g.Set(-1, -2, "quux")
	g.Set(2, 2, "quuux")

	is := g.IndexCoords(coord.NewAxial(-1, -1).LineTo(coord.NewAxial(2, 1)))
	names := make([]string, 0)
	coords := make([][]int, 0)

	for _, i := range is {
		names = append(names, i.Value())

		q, r := i.Index().Unpack()
		coords = append(coords, []int{q, r})
	}

	assert.Equal(
		t,
		[]string{"bar", "foo"},
		names,
	)

	assert.Equal(
		t,
		[][]int{
			{-1, -1},
			{0, 0},
		},
		coords,
	)
}

func TestQrGrid_IndexMovementRange(t *testing.T) {
	g := NewGrid[string]()
	g.Set(0, 0, "foo")
	g.Set(-1, -1, "bar")
	g.Set(-1, 2, "baz")
	g.Set(2, -1, "qux")
	g.Set(-1, -2, "quux")
	g.Set(2, 2, "quuux")

	is := g.IndexCoords(coord.NewAxial(1, 1).MovementRange(2))
	names := make([]string, 0)
	coords := make([][]int, 0)

	for _, i := range is {
		names = append(names, i.Value())

		q, r := i.Index().Unpack()
		coords = append(coords, []int{q, r})
	}

	assert.Equal(
		t,
		[]string{"baz", "foo", "qux", "quuux"},
		names,
	)

	assert.Equal(
		t,
		[][]int{
			{-1, 2},
			{0, 0},
			{2, -1},
			{2, 2},
		},
		coords,
	)
}
