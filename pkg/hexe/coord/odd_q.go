package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

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

func (o OddQ) Neighbors() [consts.Sides]OddQ {
	neighbors := [consts.Sides]OddQ{}
	for i, neighbor := range o.Axial().Neighbors() {
		neighbors[i] = neighbor.OddQ()
	}
	return neighbors
}

func (o OddQ) DiagonalNeighbors() [consts.Sides]OddQ {
	neighbors := [consts.Sides]OddQ{}
	for i, neighbor := range o.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.OddQ()
	}
	return neighbors
}

func (o OddQ) DistanceTo(other OddQ) int {
	aq, ar, as := o.Cube().Unpack()
	bq, br, bs := other.Cube().Unpack()
	return math.CubeDistance(aq, ar, as, bq, br, bs)
}
