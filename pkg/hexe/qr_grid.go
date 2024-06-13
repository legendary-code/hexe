package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type QRGrid[T any, C coords.CoordQR] interface {
	Grid[T]
	Get(q int, r int) T
	Index(q int, r int) (T, bool)
	Set(q int, r int, value T)
	Delete(q int, r int)
	Neighbors(q int, r int) []Item[T, C]
}

type qrGrid[T any, C any] struct {
	*grid[T]
	toAxial   func(q int, r int) (int, int)
	fromAxial func(q int, r int) (int, int)
}

func (g *qrGrid[T, C]) Get(q int, r int) T {
	q, r = g.toAxial(q, r)
	return g.grid.get(q, r)
}

func (g *qrGrid[T, C]) Index(q int, r int) (T, bool) {
	q, r = g.toAxial(q, r)
	return g.grid.index(q, r)
}

func (g *qrGrid[T, C]) Set(q int, r int, value T) {
	q, r = g.toAxial(q, r)
	g.grid.set(q, r, value)
}

func (g *qrGrid[T, C]) Delete(q int, r int) {
	q, r = g.toAxial(q, r)
	g.grid.delete(q, r)
}

func (g *qrGrid[T, C]) Neighbors(q int, r int) []Item[T, C] {
	neighbors := make([]Item[T, C], 0)
	neighborCoords := coords.Axial(q, r).Neighbors()

	for _, neighborCoord := range neighborCoords {
		value, ok := g.grid.index(neighborCoord.Q(), neighborCoord.R())
		if ok {
			neighbors = append(neighbors, newItem[T, C](neighborCoord.(C), value))
		}
	}

	return neighbors
}
