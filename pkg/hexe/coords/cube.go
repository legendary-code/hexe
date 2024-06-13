package coords

import (
	"github.com/legendary-code/hexe/internal/hexe/check"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type cube [3]int

var cubeNeighborCoords = [consts.Sides][3]int{
	{1, 0, -1},
	{-1, 0, 1},
	{1, -1, 0},
	{-1, 1, 0},
	{0, -1, 1},
	{0, 1, -1},
}

func Cube(q int, r int, s int) CoordQRS {
	check.That(q+r+s == 0, "cube coordinates (q + r + s) must equal 0")
	return &cube{q, r, s}
}

func (c cube) Axial() CoordQR {
	return Axial(c[0], c[1])
}

func (c cube) Cube() CoordQRS {
	return c
}

func (c cube) OddR() CoordQR {
	return OddR(c[0]+(c[1]-c[1]&1)/2, c[1])
}

func (c cube) EvenR() CoordQR {
	return EvenR(c[0]+(c[1]+c[1]&1)/2, c[1])
}

func (c cube) OddQ() CoordQR {
	return OddQ(c[0], c[1]+(c[0]-(c[0]&1))/2)
}

func (c cube) EvenQ() CoordQR {
	return EvenQ(c[0], c[1]+(c[0]+(c[0]&1))/2)
}

func (c cube) DoubleWidth() CoordQR {
	return DoubleWidth(2*c[0]+c[1], c[1])
}

func (c cube) DoubleHeight() CoordQR {
	return DoubleHeight(c[0], 2*c[1]+c[0])
}

func (c cube) Q() int {
	return c[0]
}

func (c cube) R() int {
	return c[1]
}

func (c cube) S() int {
	return c[2]
}

func (c cube) Unpack() (int, int, int) {
	return c[0], c[1], c[2]
}

func (c cube) Neighbors() [consts.Sides]CoordQRS {
	neighbors := [consts.Sides]CoordQRS{}
	for i, neighborCoord := range cubeNeighborCoords {
		neighbors[i] = Cube(c[0]+neighborCoord[0], c[1]+neighborCoord[1], c[2]+neighborCoord[2])
	}
	return neighbors
}
