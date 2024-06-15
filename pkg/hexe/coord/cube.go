package coord

import (
	"github.com/legendary-code/hexe/internal/hexe/check"
	hm "github.com/legendary-code/hexe/internal/hexe/math"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

type Cube [3]int

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

func (c Cube) Type() consts.CoordType {
	return consts.Cube
}

func (c Cube) Convert(typ consts.CoordType) Coord {
	return convert(c, typ)
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

func (c Cube) Neighbors() Cubes {
	neighbors := make(Cubes, consts.Sides)
	for i, neighborCoord := range cubeNeighborCoords {
		neighbors[i] = NewCube(c[0]+neighborCoord[0], c[1]+neighborCoord[1], c[2]+neighborCoord[2])
	}
	return neighbors
}

func (c Cube) DiagonalNeighbors() Cubes {
	neighbors := make(Cubes, consts.Sides)
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

func (c Cube) LineTo(other Cube) Cubes {
	coords := make([]Cube, 0)
	for _, coord := range math.CubeLineDraw(c.Q(), c.R(), c.S(), other.Q(), other.R(), other.S()) {
		coords = append(coords, NewCube(coord[0], coord[1], coord[2]))
	}
	return coords
}

func (c Cube) MovementRange(n int) Cubes {
	results := make(Cubes, 0)
	for q := -n; q <= n; q++ {
		for r := hm.Maxi(-n, -q-n); r <= hm.Mini(n, -q+n); r++ {
			s := -q - r
			results = append(results, NewCube(c.Q()+q, c.R()+r, c.S()+s))
		}
	}
	return results
}

func (c Cube) FloodFill(n int, blocked Predicate[Cube]) Cubes {
	return c.Axial().FloodFill(n, func(coord Axial) bool {
		return blocked(coord.Cube())
	}).Cubes()
}
