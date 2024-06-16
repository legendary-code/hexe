package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type DoubleWidth [2]int

func NewDoubleWidth(q int, r int) DoubleWidth {
	return DoubleWidth{q, r}
}

func (d DoubleWidth) Type() consts.CoordType {
	return consts.DoubleWidth
}

func (d DoubleWidth) Convert(typ consts.CoordType) Coord {
	return convertCoord(d, typ)
}

func (d DoubleWidth) Axial() Axial {
	return NewAxial((d[0]-d[1])/2, d[1])
}

func (d DoubleWidth) Cube() Cube {
	return d.Axial().Cube()
}

func (d DoubleWidth) OddR() OddR {
	return d.Axial().OddR()
}

func (d DoubleWidth) EvenR() EvenR {
	return d.Axial().EvenR()
}

func (d DoubleWidth) OddQ() OddQ {
	return d.Axial().OddQ()
}

func (d DoubleWidth) EvenQ() EvenQ {
	return d.Axial().EvenQ()
}

func (d DoubleWidth) DoubleWidth() DoubleWidth {
	return d
}

func (d DoubleWidth) DoubleHeight() DoubleHeight {
	return d.Axial().DoubleHeight()
}

func (d DoubleWidth) Q() int {
	return d[0]
}

func (d DoubleWidth) R() int {
	return d[1]
}

func (d DoubleWidth) Unpack() (int, int) {
	return d[0], d[1]
}

func (d DoubleWidth) Neighbors() DoubleWidths {
	neighbors := make(DoubleWidths, consts.Sides)
	for i, neighbor := range d.Axial().Neighbors() {
		neighbors[i] = neighbor.DoubleWidth()
	}
	return neighbors
}

func (d DoubleWidth) DiagonalNeighbors() DoubleWidths {
	neighbors := make(DoubleWidths, consts.Sides)
	for i, neighbor := range d.Axial().DiagonalNeighbors() {
		neighbors[i] = neighbor.DoubleWidth()
	}
	return neighbors
}

func (d DoubleWidth) DistanceTo(other DoubleWidth) int {
	return d.Cube().DistanceTo(other.Cube())
}

func (d DoubleWidth) LineTo(other DoubleWidth) DoubleWidths {
	return d.Cube().LineTo(other.Cube()).DoubleWidths()
}

func (d DoubleWidth) MovementRange(n int) DoubleWidths {
	return d.Cube().MovementRange(n).DoubleWidths()
}

func (d DoubleWidth) FloodFill(n int, blocked Predicate[DoubleWidth]) DoubleWidths {
	return d.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.DoubleWidth())
	}).DoubleWidths()
}

func (d DoubleWidth) Rotate(center DoubleWidth, angle int) DoubleWidth {
	return d.Cube().Rotate(center.Cube(), angle).DoubleWidth()
}
