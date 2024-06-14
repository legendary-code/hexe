package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

func NewDoubleHeight(q int, r int) DoubleHeight {
	return DoubleHeight{q, r}
}

func (d DoubleHeight) Axial() Axial {
	return NewAxial(d[0], (d[1]-d[0])/2)
}

func (d DoubleHeight) Cube() Cube {
	return d.Axial().Cube()
}

func (d DoubleHeight) OddR() OddR {
	return d.Axial().OddR()
}

func (d DoubleHeight) EvenR() EvenR {
	return d.Axial().EvenR()
}

func (d DoubleHeight) OddQ() OddQ {
	return d.Axial().OddQ()
}

func (d DoubleHeight) EvenQ() EvenQ {
	return d.Axial().EvenQ()
}

func (d DoubleHeight) DoubleWidth() DoubleWidth {
	return d.Axial().DoubleWidth()
}

func (d DoubleHeight) DoubleHeight() DoubleHeight {
	return d
}

func (d DoubleHeight) Q() int {
	return d[0]
}

func (d DoubleHeight) R() int {
	return d[1]
}

func (d DoubleHeight) Unpack() (int, int) {
	return d[0], d[1]
}

func (d DoubleHeight) Neighbors() [consts.Sides]DoubleHeight {
	neighbors := [consts.Sides]DoubleHeight{}
	for i, neighbor := range d.Axial().Neighbors() {
		neighbors[i] = neighbor.DoubleHeight()
	}
	return neighbors
}

func (d DoubleHeight) DiagonalNeighbors() [consts.Sides]DoubleHeight {
	neighbors := [consts.Sides]DoubleHeight{}
	for i, neighbor := range d.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.DoubleHeight()
	}
	return neighbors
}

func (d DoubleHeight) DistanceTo(other DoubleHeight) int {
	aq, ar, as := d.Cube().Unpack()
	bq, br, bs := other.Cube().Unpack()
	return math.CubeDistance(aq, ar, as, bq, br, bs)
}
