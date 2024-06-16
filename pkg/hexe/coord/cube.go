package coord

import (
	"github.com/legendary-code/hexe/internal/hexe/check"
	hm "github.com/legendary-code/hexe/internal/hexe/math"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/math"
	"golang.org/x/exp/maps"
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

func ZeroCube() Cube {
	return NewCube(0, 0, 0)
}

func (c Cube) Type() consts.CoordType {
	return consts.Cube
}

func (c Cube) Convert(typ consts.CoordType) Coord {
	return convertCoord(c, typ)
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

func (c Cube) Neighbor(angle int) Cube {
	angle = math.ClampHexAngle(angle)
	dir := cubeNeighborCoords[angle]
	return NewCube(c[0]+dir[0], c[1]+dir[1], c[2]+dir[2])
}

func (c Cube) Add(other Cube) Cube {
	return NewCube(c[0]+other[0], c[1]+other[1], c[2]+other[2])
}

func (c Cube) Scale(factor int) Cube {
	return NewCube(c[0]*factor, c[1]*factor, c[2]*factor)
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
	visited := make(map[Cube]bool)
	visited[c] = true

	fringes := make([]Cubes, 0)
	fringes = append(fringes, Cubes{c})

	for k := 1; k <= n; k++ {
		fringes = append(fringes, Cubes{})
		for _, coord := range fringes[k-1] {
			for _, neighbor := range coord.Neighbors() {
				if _, ok := visited[neighbor]; ok {
					continue
				}

				if blocked(neighbor) {
					continue
				}

				visited[neighbor] = true
				fringes[k] = append(fringes[k], neighbor)
			}
		}
	}

	return maps.Keys(visited)
}

func (c Cube) Rotate(center Cube, angle int) Cube {
	q, r, s := c.Unpack()
	cq, cr, cs := center.Unpack()
	dq, dr, ds := q-cq, r-cr, s-cs
	angle = math.ClampHexAngle(angle)

	for i := 0; i < angle; i++ {
		dq, dr, ds = -dr, -ds, -dq
	}

	return NewCube(cq+dq, cr+dr, cs+ds)
}

func (c Cube) ReflectQ() Cube {
	return NewCube(c[0], c[2], c[1])
}

func (c Cube) ReflectR() Cube {
	return NewCube(c[2], c[1], c[0])
}

func (c Cube) ReflectS() Cube {
	return NewCube(c[1], c[0], c[2])
}

func (c Cube) Ring(radius int) Cubes {
	if radius < 1 {
		return Cubes{c}
	}

	results := make(Cubes, 0)
	coord := c.Add(ZeroCube().Neighbor(4).Scale(radius))

	for i := 0; i < consts.Sides; i++ {
		for j := 0; j < radius; j++ {
			results = append(results, coord)
			coord = coord.Neighbor(i)
		}
	}

	return results
}
