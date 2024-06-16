package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type EvenR [2]int

func NewEvenR(q int, r int) EvenR {
	return EvenR{q, r}
}

func (e EvenR) Type() consts.CoordType {
	return consts.EvenR
}

func (e EvenR) Convert(typ consts.CoordType) Coord {
	return convertCoord(e, typ)
}

func (e EvenR) Axial() Axial {
	return NewAxial(e[0]-(e[1]+(e[1]&1))/2, e[1])
}

func (e EvenR) Cube() Cube {
	return e.Axial().Cube()
}

func (e EvenR) OddR() OddR {
	return e.Axial().OddR()
}

func (e EvenR) EvenR() EvenR {
	return e
}

func (e EvenR) OddQ() OddQ {
	return e.Axial().OddQ()
}

func (e EvenR) EvenQ() EvenQ {
	return e.Axial().EvenQ()
}

func (e EvenR) DoubleWidth() DoubleWidth {
	return e.Axial().DoubleWidth()
}

func (e EvenR) DoubleHeight() DoubleHeight {
	return e.Axial().DoubleHeight()
}

func (e EvenR) Q() int {
	return e[0]
}

func (e EvenR) R() int {
	return e[1]
}

func (e EvenR) Unpack() (int, int) {
	return e[0], e[1]
}

func (e EvenR) Neighbors() EvenRs {
	neighbors := make(EvenRs, consts.Sides)
	for i, neighbor := range e.Axial().Neighbors() {
		neighbors[i] = neighbor.EvenR()
	}
	return neighbors
}

func (e EvenR) DiagonalNeighbors() EvenRs {
	neighbors := make(EvenRs, consts.Sides)
	for i, neighbor := range e.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.EvenR()
	}
	return neighbors
}

func (e EvenR) DistanceTo(other EvenR) int {
	return e.Cube().DistanceTo(other.Cube())
}

func (e EvenR) LineTo(other EvenR) EvenRs {
	return e.Cube().LineTo(other.Cube()).EvenRs()
}

func (e EvenR) MovementRange(n int) EvenRs {
	return e.Cube().MovementRange(n).EvenRs()
}

func (e EvenR) FloodFill(n int, blocked Predicate[EvenR]) EvenRs {
	return e.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.EvenR())
	}).EvenRs()
}

func (e EvenR) Rotate(center EvenR, angle int) EvenR {
	return e.Cube().Rotate(center.Cube(), angle).EvenR()
}

func (e EvenR) ReflectQ() EvenR {
	return e.Cube().ReflectQ().EvenR()
}

func (e EvenR) ReflectR() EvenR {
	return e.Cube().ReflectR().EvenR()
}

func (e EvenR) ReflectS() EvenR {
	return e.Cube().ReflectS().EvenR()
}
