package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"golang.org/x/exp/maps"
	"slices"
	"sort"
)

type Axials []Axial

func (a Axials) Type() consts.CoordType {
	return consts.Axial
}

func (a Axials) Coords() []Coord {
	return toCoords(a)
}

func (a Axials) Axials() Axials {
	return a
}

func (a Axials) Cubes() Cubes {
	return castAs(a, Axial.Cube)
}

func (a Axials) DoubleWidths() DoubleWidths {
	return castAs(a, Axial.DoubleWidth)
}

func (a Axials) DoubleHeights() DoubleHeights {
	return castAs(a, Axial.DoubleHeight)
}

func (a Axials) EvenQs() EvenQs {
	return castAs(a, Axial.EvenQ)
}

func (a Axials) EvenRs() EvenRs {
	return castAs(a, Axial.EvenR)
}

func (a Axials) OddQs() OddQs {
	return castAs(a, Axial.OddQ)
}

func (a Axials) OddRs() OddRs {
	return castAs(a, Axial.OddR)
}

func (a Axials) Copy() Axials {
	return slices.Clone(a)
}

func (a Axials) Sorted() Axials {
	sorted := a.Copy()
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Q() == sorted[j].Q() {
			return sorted[i].R() < sorted[j].R()
		}
		return sorted[i].Q() < sorted[j].Q()
	})
	return sorted
}

func (a Axials) UnionWith(other Axials) Axials {
	coords := make(map[Axial]bool)
	for _, c := range a {
		coords[c] = true
	}
	for _, c := range other {
		coords[c] = true
	}
	return maps.Keys(coords)
}

func (a Axials) IntersectWith(other Axials) Axials {
	coords := make(map[Axial]bool)
	intersection := make(map[Axial]bool)

	for _, c := range a {
		coords[c] = true
	}

	for _, c := range other {
		if _, ok := coords[c]; ok {
			intersection[c] = true
		}
	}

	return maps.Keys(intersection)
}

func (a Axials) DifferenceWith(other Axials) Axials {
	coords := make(map[Axial]bool)

	for _, c := range a {
		coords[c] = true
	}

	for _, c := range other {
		delete(coords, c)
	}

	return maps.Keys(coords)
}
