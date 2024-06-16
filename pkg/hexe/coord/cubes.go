package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"golang.org/x/exp/maps"
	"slices"
	"sort"
)

type Cubes []Cube

func (c Cubes) Type() consts.CoordType {
	return consts.Cube
}

func (c Cubes) Convert(typ consts.CoordType) Coords {
	return convertCoords(c, typ)
}

func (c Cubes) ToSlice() []Coord {
	return toCoords(c)
}

func (c Cubes) Axials() Axials {
	return castAs(c, Cube.Axial)
}

func (c Cubes) Cubes() Cubes {
	return c
}

func (c Cubes) DoubleWidths() DoubleWidths {
	return castAs(c, Cube.DoubleWidth)
}

func (c Cubes) DoubleHeights() DoubleHeights {
	return castAs(c, Cube.DoubleHeight)
}

func (c Cubes) EvenQs() EvenQs {
	return castAs(c, Cube.EvenQ)
}

func (c Cubes) EvenRs() EvenRs {
	return castAs(c, Cube.EvenR)
}

func (c Cubes) OddQs() OddQs {
	return castAs(c, Cube.OddQ)
}

func (c Cubes) OddRs() OddRs {
	return castAs(c, Cube.OddR)
}

func (c Cubes) Copy() Cubes {
	return slices.Clone(c)
}

func (c Cubes) Sort() Cubes {
	sorted := c.Copy()
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Q() == sorted[j].Q() {
			return sorted[i].R() < sorted[j].R()
		}
		return sorted[i].Q() < sorted[j].Q()
	})
	return sorted
}

func (c Cubes) UnionWith(other Cubes) Cubes {
	coords := make(map[Cube]bool)
	for _, i := range c {
		coords[i] = true
	}
	for _, i := range other {
		coords[i] = true
	}
	return maps.Keys(coords)

}

func (c Cubes) IntersectWith(other Cubes) Cubes {
	coords := make(map[Cube]bool)
	intersection := make(map[Cube]bool)

	for _, i := range c {
		coords[i] = true
	}

	for _, i := range other {
		if _, ok := coords[i]; ok {
			intersection[i] = true
		}
	}

	return maps.Keys(intersection)
}

func (c Cubes) DifferenceWith(other Cubes) Cubes {
	coords := make(map[Cube]bool)

	for _, i := range c {
		coords[i] = true
	}

	for _, i := range other {
		delete(coords, i)
	}

	return maps.Keys(coords)
}

func (c Cubes) Rotate(center Cube, angle int) Cubes {
	coords := make(Cubes, len(c))
	for i, coord := range c {
		coords[i] = coord.Rotate(center, angle)
	}
	return coords
}
