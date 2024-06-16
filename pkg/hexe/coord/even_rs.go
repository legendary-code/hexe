package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type EvenRs []EvenR

func (e EvenRs) Type() consts.CoordType {
	return consts.EvenR
}

func (e EvenRs) Convert(typ consts.CoordType) Coords {
	return convertCoords(e, typ)
}

func (e EvenRs) ToSlice() []Coord {
	return toCoords(e)
}

func (e EvenRs) Axials() Axials {
	return castAs(e, EvenR.Axial)
}

func (e EvenRs) Cubes() Cubes {
	return castAs(e, EvenR.Cube)
}

func (e EvenRs) DoubleWidths() DoubleWidths {
	return castAs(e, EvenR.DoubleWidth)
}

func (e EvenRs) DoubleHeights() DoubleHeights {
	return castAs(e, EvenR.DoubleHeight)
}

func (e EvenRs) EvenQs() EvenQs {
	return castAs(e, EvenR.EvenQ)
}

func (e EvenRs) EvenRs() EvenRs {
	return e
}

func (e EvenRs) OddQs() OddQs {
	return castAs(e, EvenR.OddQ)
}

func (e EvenRs) OddRs() OddRs {
	return castAs(e, EvenR.OddR)
}

func (e EvenRs) Copy() EvenRs {
	return slices.Clone(e)
}

func (e EvenRs) Sort() EvenRs {
	return e.Cubes().Sort().EvenRs()
}

func (e EvenRs) UnionWith(other EvenRs) EvenRs {
	return e.Cubes().UnionWith(other.Cubes()).EvenRs()
}

func (e EvenRs) IntersectWith(other EvenRs) EvenRs {
	return e.Cubes().IntersectWith(other.Cubes()).EvenRs()
}

func (e EvenRs) DifferenceWith(other EvenRs) EvenRs {
	return e.Cubes().DifferenceWith(other.Cubes()).EvenRs()
}

func (e EvenRs) Rotate(center EvenR, angle int) EvenRs {
	return e.Cubes().Rotate(center.Cube(), angle).EvenRs()
}
