package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type QRSGrid[T any, C coords.CoordQRS] interface {
	Grid[T]
	Get(q int, r int, s int) T
	Index(q int, r int, s int) (T, bool)
	Set(q int, r int, s int, value T)
	Delete(q int, r int, s int)
	Neighbors(q int, r int, s int) []Item[T, C]
}

type qrsGrid[T any, C any] struct {
	*grid[T]
	toAxial   func(q int, r int, s int) (int, int)
	fromAxial func(q int, r int) (int, int, int)
}

func (g *qrsGrid[T, C]) Get(q int, r int, s int) T {
	q, r = g.toAxial(q, r, s)
	return g.grid.get(q, r)
}

func (g *qrsGrid[T, C]) Index(q int, r int, s int) (T, bool) {
	q, r = g.toAxial(q, r, s)
	return g.grid.index(q, r)
}

func (g *qrsGrid[T, C]) Set(q int, r int, s int, value T) {
	q, r = g.toAxial(q, r, s)
	g.grid.set(q, r, value)
}

func (g *qrsGrid[T, C]) Delete(q int, r int, s int) {
	q, r = g.toAxial(q, r, s)
	g.grid.delete(q, r)
}

func (g *qrsGrid[T, C]) Neighbors(q int, r int, s int) []Item[T, C] {
	neighbors := make([]Item[T, C], 0)
	neighborCoords := coords.Cube(q, r, s).Neighbors()

	for _, neighborCoord := range neighborCoords {
		value, ok := g.grid.index(neighborCoord.Q(), neighborCoord.R())
		if ok {
			neighbors = append(neighbors, newItem[T, C](neighborCoord.(C), value))
		}
	}

	return neighbors
}
