package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type Axial [2]int

func NewAxial(q int, r int) Axial {
	return Axial{q, r}
}

func (a Axial) Type() consts.CoordType {
	return consts.Axial
}

func (a Axial) Convert(typ consts.CoordType) Coord {
	return convertCoord(a, typ)
}

func (a Axial) Axial() Axial {
	return a
}

func (a Axial) Cube() Cube {
	return NewCube(a[0], a[1], -a[0]-a[1])
}

func (a Axial) OddR() OddR {
	return NewOddR(a[0]+(a[1]-a[1]&1)/2, a[1])
}

func (a Axial) EvenR() EvenR {
	return NewEvenR(a[0]+(a[1]+a[1]&1)/2, a[1])
}

func (a Axial) OddQ() OddQ {
	return NewOddQ(a[0], a[1]+(a[0]-(a[0]&1))/2)
}

func (a Axial) EvenQ() EvenQ {
	return NewEvenQ(a[0], a[1]+(a[0]+(a[0]&1))/2)
}

func (a Axial) DoubleWidth() DoubleWidth {
	return NewDoubleWidth(2*a[0]+a[1], a[1])
}

func (a Axial) DoubleHeight() DoubleHeight {
	return NewDoubleHeight(a[0], 2*a[1]+a[0])
}

func (a Axial) Q() int {
	return a[0]
}

func (a Axial) R() int {
	return a[1]
}

func (a Axial) Unpack() (int, int) {
	return a[0], a[1]
}

func (a Axial) Neighbors() Axials {
	return a.Cube().Neighbors().Axials()
}

func (a Axial) DiagonalNeighbors() Axials {
	return a.Cube().DiagonalNeighbors().Axials()
}

func (a Axial) DistanceTo(other Axial) int {
	return a.Cube().DistanceTo(other.Cube())
}

func (a Axial) LineTo(other Axial) Axials {
	return a.Cube().LineTo(other.Cube()).Axials()
}

func (a Axial) MovementRange(n int) Axials {
	return a.Cube().MovementRange(n).Axials()
}

func (a Axial) FloodFill(n int, blocked Predicate[Axial]) Axials {
	return a.Cube().FloodFill(n, func(coord Cube) bool {
		return blocked(coord.Axial())
	}).Axials()
}

func (a Axial) Rotate(center Axial, angle int) Axial {
	return a.Cube().Rotate(center.Cube(), angle).Axial()
}
