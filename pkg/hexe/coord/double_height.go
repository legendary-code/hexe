package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type DoubleHeight [2]int

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

func (d DoubleHeight) Neighbors() DoubleHeights {
	neighbors := make(DoubleHeights, consts.Sides)
	for i, neighbor := range d.Axial().Neighbors() {
		neighbors[i] = neighbor.DoubleHeight()
	}
	return neighbors
}

func (d DoubleHeight) DiagonalNeighbors() DoubleHeights {
	neighbors := make(DoubleHeights, consts.Sides)
	for i, neighbor := range d.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.DoubleHeight()
	}
	return neighbors
}

func (d DoubleHeight) DistanceTo(other DoubleHeight) int {
	return d.Cube().DistanceTo(other.Cube())
}

func (d DoubleHeight) LineTo(other DoubleHeight) DoubleHeights {
	return d.Cube().LineTo(other.Cube()).DoubleHeights()
}

func (d DoubleHeight) MovementRange(n int) DoubleHeights {
	return d.Cube().MovementRange(n).DoubleHeights()
}
