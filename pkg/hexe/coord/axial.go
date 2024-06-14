package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

var axialNeighborCoords = [consts.Sides][2]int{
	{1, 0},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{0, -1},
	{1, -1},
}

var axialDiagonalNeighborCoords = [consts.Sides][2]int{
	{1, 1},
	{-1, 2},
	{-2, 1},
	{-1, -1},
	{1, -2},
	{2, -1},
}

func NewAxial(q int, r int) Axial {
	return Axial{q, r}
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

func (a Axial) Neighbors() [consts.Sides]Axial {
	neighbors := [consts.Sides]Axial{}
	for i, neighborCoord := range axialNeighborCoords {
		neighbors[i] = NewAxial(a[0]+neighborCoord[0], a[1]+neighborCoord[1])
	}
	return neighbors
}

func (a Axial) DiagonalNeighbors() [consts.Sides]Axial {
	neighbors := [consts.Sides]Axial{}
	for i, neighborCoord := range axialDiagonalNeighborCoords {
		neighbors[i] = NewAxial(a[0]+neighborCoord[0], a[1]+neighborCoord[1])
	}
	return neighbors
}

func (a Axial) DistanceTo(other Axial) int {
	aq, ar, as := a.Cube().Unpack()
	bq, br, bs := other.Cube().Unpack()
	return math.CubeDistance(aq, ar, as, bq, br, bs)
}
