package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type QRSGrid[T any, C coord.QRSCoord[C, CS], CS coord.Coords[C, CS]] interface {
	Grid[T]

	// Get returns item at given coordinate
	Get(q int, r int, s int) T

	// Index returns item at given coordinate and a bool of whether it exists, similar to indexing into a map
	Index(q int, r int, s int) (T, bool)

	// Set sets an item at the given coordinate
	Set(q int, r int, s int, value T)

	// Delete deletes an item at the given coordinate
	Delete(q int, r int, s int)

	// IndexCoords returns items at the given coordinates
	IndexCoords(coords CS) Items[T, C, CS]
}

type qrsGrid[T any, C coord.QRSCoord[C, CS], CS coord.Coords[C, CS]] struct {
	*grid[T]
	toAxial   func(q int, r int, s int) (int, int)
	fromAxial func(q int, r int) C
}

func (g *qrsGrid[T, C, CS]) Get(q int, r int, s int) T {
	q, r = g.toAxial(q, r, s)
	return g.grid.get(q, r)
}

func (g *qrsGrid[T, C, CS]) Index(q int, r int, s int) (T, bool) {
	q, r = g.toAxial(q, r, s)
	return g.grid.index(q, r)
}

func (g *qrsGrid[T, C, CS]) Set(q int, r int, s int, value T) {
	q, r = g.toAxial(q, r, s)
	g.grid.set(q, r, value)
}

func (g *qrsGrid[T, C, CS]) Delete(q int, r int, s int) {
	q, r = g.toAxial(q, r, s)
	g.grid.delete(q, r)
}

func (g *qrsGrid[T, C, CS]) indexCoords(coords coord.Axials) Items[T, C, CS] {
	items := make(Items[T, C, CS], 0)
	for _, c := range coords {
		value, ok := g.grid.index(c.Q(), c.R())
		if ok {
			typedCoord := g.fromAxial(c.Q(), c.R())
			items = append(items, newItem[T, C](typedCoord, value))
		}
	}

	return items
}

func (g *qrsGrid[T, C, CS]) IndexCoords(coords CS) Items[T, C, CS] {
	items := make(Items[T, C, CS], 0)
	for _, c := range coords {
		value, ok := g.Index(c.Q(), c.R(), c.S())
		if ok {
			items = append(items, newItem[T, C](c, value))
		}
	}

	return items
}

func (g *qrsGrid[T, C, CS]) Neighbors(q int, r int, s int) Items[T, C, CS] {
	q, r = g.toAxial(q, r, s)
	coords := coord.NewAxial(q, r).Neighbors()
	return g.indexCoords(coords)
}

func (g *qrsGrid[T, C, CS]) DiagonalNeighbors(q int, r int, s int) Items[T, C, CS] {
	q, r = g.toAxial(q, r, s)
	coords := coord.NewAxial(q, r).DiagonalNeighbors()
	return g.indexCoords(coords)
}

func (g *qrsGrid[T, C, CS]) IndexLine(aq int, ar int, as int, bq int, br int, bs int) Items[T, C, CS] {
	aq, ar = g.toAxial(aq, ar, as)
	bq, br = g.toAxial(bq, br, bs)
	coords := coord.NewAxial(aq, ar).LineTo(coord.NewAxial(bq, br))
	return g.indexCoords(coords)
}

func (g *qrsGrid[T, C, CS]) IndexMovementRange(q int, r int, s int, n int) Items[T, C, CS] {
	q, r = g.toAxial(q, r, s)
	coords := coord.NewAxial(q, r).MovementRange(n)
	return g.indexCoords(coords)
}
