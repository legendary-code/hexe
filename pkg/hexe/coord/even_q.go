package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type EvenQ [2]int

func NewEvenQ(q int, r int) EvenQ {
	return EvenQ{q, r}
}

func (e EvenQ) Type() consts.CoordType {
	return consts.EvenQ
}

func (e EvenQ) Convert(typ consts.CoordType) Coord {
	return convert(e, typ)
}

func (e EvenQ) Axial() Axial {
	return NewAxial(e[0], e[1]-(e[0]+(e[0]&1))/2)
}

func (e EvenQ) Cube() Cube {
	return e.Axial().Cube()
}

func (e EvenQ) OddR() OddR {
	return e.Axial().OddR()
}

func (e EvenQ) EvenR() EvenR {
	return e.Axial().EvenR()
}

func (e EvenQ) OddQ() OddQ {
	return e.Axial().OddQ()
}

func (e EvenQ) EvenQ() EvenQ {
	return e
}

func (e EvenQ) DoubleWidth() DoubleWidth {
	return e.Axial().DoubleWidth()
}

func (e EvenQ) DoubleHeight() DoubleHeight {
	return e.Axial().DoubleHeight()
}

func (e EvenQ) Q() int {
	return e[0]
}

func (e EvenQ) R() int {
	return e[1]
}

func (e EvenQ) Unpack() (int, int) {
	return e[0], e[1]
}

func (e EvenQ) Neighbors() EvenQs {
	neighbors := make(EvenQs, consts.Sides)
	for i, neighbor := range e.Axial().Neighbors() {
		neighbors[i] = neighbor.EvenQ()
	}
	return neighbors
}

func (e EvenQ) DiagonalNeighbors() EvenQs {
	neighbors := make(EvenQs, consts.Sides)
	for i, neighbor := range e.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.EvenQ()
	}
	return neighbors
}

func (e EvenQ) DistanceTo(other EvenQ) int {
	return e.Cube().DistanceTo(other.Cube())
}

func (e EvenQ) LineTo(other EvenQ) EvenQs {
	return e.Cube().LineTo(other.Cube()).EvenQs()
}

func (e EvenQ) MovementRange(n int) EvenQs {
	return e.Cube().MovementRange(n).EvenQs()
}

func (e EvenQ) FloodFill(n int, blocked Predicate[EvenQ]) EvenQs {
	return e.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.EvenQ())
	}).EvenQs()
}

func (e EvenQ) Rotate(center EvenQ, angle int) EvenQ {
	return e.Cube().Rotate(center.Cube(), angle).EvenQ()
}
