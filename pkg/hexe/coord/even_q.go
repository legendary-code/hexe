package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

func NewEvenQ(q int, r int) EvenQ {
	return EvenQ{q, r}
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

func (e EvenQ) Neighbors() [consts.Sides]EvenQ {
	neighbors := [consts.Sides]EvenQ{}
	for i, neighbor := range e.Axial().Neighbors() {
		neighbors[i] = neighbor.EvenQ()
	}
	return neighbors
}

func (e EvenQ) DiagonalNeighbors() [consts.Sides]EvenQ {
	neighbors := [consts.Sides]EvenQ{}
	for i, neighbor := range e.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.EvenQ()
	}
	return neighbors
}

func (e EvenQ) DistanceTo(other EvenQ) int {
	aq, ar, as := e.Cube().Unpack()
	bq, br, bs := other.Cube().Unpack()
	return math.CubeDistance(aq, ar, as, bq, br, bs)
}
