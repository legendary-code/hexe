package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

func NewOddR(q int, r int) OddR {
	return OddR{q, r}
}

func (o OddR) Axial() Axial {
	return NewAxial(o[0]-(o[1]-(o[1]&1))/2, o[1])
}

func (o OddR) Cube() Cube {
	return o.Axial().Cube()
}

func (o OddR) OddR() OddR {
	return o
}

func (o OddR) EvenR() EvenR {
	return o.Axial().EvenR()
}

func (o OddR) OddQ() OddQ {
	return o.Axial().OddQ()
}

func (o OddR) EvenQ() EvenQ {
	return o.Axial().EvenQ()
}

func (o OddR) DoubleWidth() DoubleWidth {
	return o.Axial().DoubleWidth()
}

func (o OddR) DoubleHeight() DoubleHeight {
	return o.Axial().DoubleHeight()
}

func (o OddR) Q() int {
	return o[0]
}

func (o OddR) R() int {
	return o[1]
}

func (o OddR) Unpack() (int, int) {
	return o[0], o[1]
}

func (o OddR) Neighbors() [consts.Sides]OddR {
	neighbors := [consts.Sides]OddR{}
	for i, neighbor := range o.Axial().Neighbors() {
		neighbors[i] = neighbor.OddR()
	}
	return neighbors
}

func (o OddR) DiagonalNeighbors() [consts.Sides]OddR {
	neighbors := [consts.Sides]OddR{}
	for i, neighbor := range o.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.OddR()
	}
	return neighbors
}

func (o OddR) DistanceTo(other OddR) int {
	aq, ar, as := o.Cube().Unpack()
	bq, br, bs := other.Cube().Unpack()
	return math.CubeDistance(aq, ar, as, bq, br, bs)
}
