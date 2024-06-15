package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type QRGrid[T any, C coord.QR[C, CS], CS coord.Coords[C, CS]] interface {
	Grid[T]

	// Get returns item at given coordinate
	Get(q int, r int) T

	// Index returns item at given coordinate and a bool of whether it exists, similar to indexing into a map
	Index(q int, r int) (T, bool)

	// Set sets an item at the given coordinate
	Set(q int, r int, value T)

	// Delete deletes an item at the given coordinate
	Delete(q int, r int)

	// Neighbors returns items that neighbor the given coordinate
	Neighbors(q int, r int) Items[T, C, CS]

	// DiagonalNeighbors returns items that diagonally neighbor the given coordinate
	DiagonalNeighbors(q int, r int) Items[T, C, CS]

	// IndexCoords returns items at the given coordinates
	IndexCoords(coords CS) Items[T, C, CS]

	// IndexLine returns items that fall on the line between the two coordinates
	IndexLine(aq int, ar int, bq int, br int) Items[T, C, CS]

	// IndexMovementRange returns items within the given radius and the given coordinates
	IndexMovementRange(q int, r int, n int) Items[T, C, CS]
}

type qrGrid[T any, C coord.QR[C, CS], CS coord.Coords[C, CS]] struct {
	*grid[T]
	toAxial   func(q int, r int) (int, int)
	fromAxial func(q int, r int) C
}

func (g *qrGrid[T, C, CS]) Get(q int, r int) T {
	q, r = g.toAxial(q, r)
	return g.grid.get(q, r)
}

func (g *qrGrid[T, C, CS]) Index(q int, r int) (T, bool) {
	q, r = g.toAxial(q, r)
	return g.grid.index(q, r)
}

func (g *qrGrid[T, C, CS]) Set(q int, r int, value T) {
	q, r = g.toAxial(q, r)
	g.grid.set(q, r, value)
}

func (g *qrGrid[T, C, CS]) Delete(q int, r int) {
	q, r = g.toAxial(q, r)
	g.grid.delete(q, r)
}

func (g *qrGrid[T, C, CS]) indexCoords(coords coord.Axials) Items[T, C, CS] {
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

func (g *qrGrid[T, C, CS]) IndexCoords(coords CS) Items[T, C, CS] {
	items := make(Items[T, C, CS], 0)
	for _, c := range coords {
		value, ok := g.Index(c.Q(), c.R())
		if ok {
			items = append(items, newItem[T, C](c, value))
		}
	}

	return items
}

func (g *qrGrid[T, C, CS]) Neighbors(q int, r int) Items[T, C, CS] {
	q, r = g.toAxial(q, r)
	coords := coord.NewAxial(q, r).Neighbors()
	return g.indexCoords(coords)
}

func (g *qrGrid[T, C, CS]) DiagonalNeighbors(q int, r int) Items[T, C, CS] {
	q, r = g.toAxial(q, r)
	coords := coord.NewAxial(q, r).DiagonalNeighbors()
	return g.indexCoords(coords)
}

func (g *qrGrid[T, C, CS]) IndexLine(aq int, ar int, bq int, br int) Items[T, C, CS] {
	aq, ar = g.toAxial(aq, ar)
	bq, br = g.toAxial(bq, br)
	coords := coord.NewAxial(aq, ar).LineTo(coord.NewAxial(bq, br))
	return g.indexCoords(coords)
}

func (g *qrGrid[T, C, CS]) IndexMovementRange(q int, r int, n int) Items[T, C, CS] {
	q, r = g.toAxial(q, r)
	coords := coord.NewAxial(q, r).MovementRange(n)
	return g.indexCoords(coords)
}
