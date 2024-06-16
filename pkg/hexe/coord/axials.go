package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type Axials []Axial

func (a Axials) Type() consts.CoordType {
	return consts.Axial
}

func (a Axials) Convert(typ consts.CoordType) Coords {
	return convertCoords(a, typ)
}

func (a Axials) ToSlice() []Coord {
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

func (a Axials) Sort() Axials {
	return a.Cubes().Sort().Axials()
}

func (a Axials) UnionWith(other Axials) Axials {
	return a.Cubes().UnionWith(other.Cubes()).Axials()
}

func (a Axials) IntersectWith(other Axials) Axials {
	return a.Cubes().IntersectWith(other.Cubes()).Axials()
}

func (a Axials) DifferenceWith(other Axials) Axials {
	return a.Cubes().DifferenceWith(other.Cubes()).Axials()
}

func (a Axials) Rotate(center Axial, angle int) Axials {
	return a.Cubes().Rotate(center.Cube(), angle).Axials()
}
