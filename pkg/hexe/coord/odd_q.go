package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type OddQ [2]int

func NewOddQ(q int, r int) OddQ {
	return OddQ{q, r}
}

func (o OddQ) Axial() Axial {
	return NewAxial(o[0], o[1]-(o[0]-(o[0]&1))/2)
}

func (o OddQ) Cube() Cube {
	return o.Axial().Cube()
}

func (o OddQ) OddR() OddR {
	return o.Axial().OddR()
}

func (o OddQ) EvenR() EvenR {
	return o.Axial().EvenR()
}

func (o OddQ) OddQ() OddQ {
	return o
}

func (o OddQ) EvenQ() EvenQ {
	return o.Axial().EvenQ()
}

func (o OddQ) DoubleWidth() DoubleWidth {
	return o.Axial().DoubleWidth()
}

func (o OddQ) DoubleHeight() DoubleHeight {
	return o.Axial().DoubleHeight()
}

func (o OddQ) Q() int {
	return o[0]
}

func (o OddQ) R() int {
	return o[1]
}

func (o OddQ) Unpack() (int, int) {
	return o[0], o[1]
}

func (o OddQ) Neighbors() OddQs {
	neighbors := make(OddQs, consts.Sides)
	for i, neighbor := range o.Axial().Neighbors() {
		neighbors[i] = neighbor.OddQ()
	}
	return neighbors
}

func (o OddQ) DiagonalNeighbors() OddQs {
	neighbors := make(OddQs, consts.Sides)
	for i, neighbor := range o.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.OddQ()
	}
	return neighbors
}

func (o OddQ) DistanceTo(other OddQ) int {
	return o.Cube().DistanceTo(other.Cube())
}

func (o OddQ) LineTo(other OddQ) OddQs {
	return o.Cube().LineTo(other.Cube()).OddQs()
}

func (o OddQ) MovementRange(n int) OddQs {
	return o.Cube().MovementRange(n).OddQs()
}
