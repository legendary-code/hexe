package coord

import (
	"github.com/legendary-code/hexe/internal/hexe/check"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

var cubeNeighborCoords = [consts.Sides][3]int{
	{1, 0, -1},
	{0, 1, -1},
	{-1, 1, 0},
	{-1, 0, 1},
	{0, -1, 1},
	{1, -1, 0},
}

var cubeDiagonalNeighborCoords = [consts.Sides][3]int{
	{1, 1, -2},
	{-1, 2, -1},
	{-2, 1, 1},
	{-1, -1, 2},
	{1, -2, 1},
	{2, -1, -1},
}

func NewCube(q int, r int, s int) Cube {
	check.That(q+r+s == 0, "cube coordinates (q + r + s) must equal 0")
	return Cube{q, r, s}
}

func (c Cube) Axial() Axial {
	return NewAxial(c[0], c[1])
}

func (c Cube) Cube() Cube {
	return c
}

func (c Cube) OddR() OddR {
	return NewOddR(c[0]+(c[1]-c[1]&1)/2, c[1])
}

func (c Cube) EvenR() EvenR {
	return NewEvenR(c[0]+(c[1]+c[1]&1)/2, c[1])
}

func (c Cube) OddQ() OddQ {
	return NewOddQ(c[0], c[1]+(c[0]-(c[0]&1))/2)
}

func (c Cube) EvenQ() EvenQ {
	return NewEvenQ(c[0], c[1]+(c[0]+(c[0]&1))/2)
}

func (c Cube) DoubleWidth() DoubleWidth {
	return NewDoubleWidth(2*c[0]+c[1], c[1])
}

func (c Cube) DoubleHeight() DoubleHeight {
	return NewDoubleHeight(c[0], 2*c[1]+c[0])
}

func (c Cube) Q() int {
	return c[0]
}

func (c Cube) R() int {
	return c[1]
}

func (c Cube) S() int {
	return c[2]
}

func (c Cube) Unpack() (int, int, int) {
	return c[0], c[1], c[2]
}

func (c Cube) Neighbors() [consts.Sides]Cube {
	neighbors := [consts.Sides]Cube{}
	for i, neighborCoord := range cubeNeighborCoords {
		neighbors[i] = NewCube(c[0]+neighborCoord[0], c[1]+neighborCoord[1], c[2]+neighborCoord[2])
	}
	return neighbors
}

func (c Cube) DiagonalNeighbors() [consts.Sides]Cube {
	neighbors := [consts.Sides]Cube{}
	for i, neighborCoord := range cubeDiagonalNeighborCoords {
		neighbors[i] = NewCube(c[0]+neighborCoord[0], c[1]+neighborCoord[1], c[2]+neighborCoord[2])
	}
	return neighbors
}

func (c Cube) DistanceTo(other Cube) int {
	aq, ar, as := c.Unpack()
	bq, br, bs := other.Unpack()
	return math.CubeDistance(aq, ar, as, bq, br, bs)
}
