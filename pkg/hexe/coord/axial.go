package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"golang.org/x/exp/maps"
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

type Axial [2]int

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

func (a Axial) Neighbors() Axials {
	neighbors := make(Axials, consts.Sides)
	for i, neighborCoord := range axialNeighborCoords {
		neighbors[i] = NewAxial(a[0]+neighborCoord[0], a[1]+neighborCoord[1])
	}
	return neighbors
}

func (a Axial) DiagonalNeighbors() Axials {
	neighbors := make(Axials, consts.Sides)
	for i, neighborCoord := range axialDiagonalNeighborCoords {
		neighbors[i] = NewAxial(a[0]+neighborCoord[0], a[1]+neighborCoord[1])
	}
	return neighbors
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

func (a Axial) FloodFill(n int, blocked CoordPredicate[Axial]) Axials {
	visited := make(map[Axial]bool)
	visited[a] = true

	fringes := make([]Axials, 0)
	fringes = append(fringes, Axials{a})

	for k := 1; k <= n; k++ {
		fringes = append(fringes, Axials{})
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
