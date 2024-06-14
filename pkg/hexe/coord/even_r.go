package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

func NewEvenR(q int, r int) EvenR {
	return EvenR{q, r}
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

func (e EvenR) Neighbors() [consts.Sides]EvenR {
	neighbors := [consts.Sides]EvenR{}
	for i, neighbor := range e.Axial().Neighbors() {
		neighbors[i] = neighbor.EvenR()
	}
	return neighbors
}

func (e EvenR) DiagonalNeighbors() [consts.Sides]EvenR {
	neighbors := [consts.Sides]EvenR{}
	for i, neighbor := range e.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.EvenR()
	}
	return neighbors
}

func (e EvenR) DistanceTo(other EvenR) int {
	aq, ar, as := e.Cube().Unpack()
	bq, br, bs := other.Cube().Unpack()
	return math.CubeDistance(aq, ar, as, bq, br, bs)
}
