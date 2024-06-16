package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type OddRs []OddR

func (o OddRs) Type() consts.CoordType {
	return consts.OddR
}

func (o OddRs) Convert(typ consts.CoordType) Coords {
	return convertCoords(o, typ)
}

func (o OddRs) ToSlice() []Coord {
	return toCoords(o)
}

func (o OddRs) Axials() Axials {
	return castAs(o, OddR.Axial)
}

func (o OddRs) Cubes() Cubes {
	return castAs(o, OddR.Cube)
}

func (o OddRs) DoubleWidths() DoubleWidths {
	return castAs(o, OddR.DoubleWidth)
}

func (o OddRs) DoubleHeights() DoubleHeights {
	return castAs(o, OddR.DoubleHeight)
}

func (o OddRs) EvenQs() EvenQs {
	return castAs(o, OddR.EvenQ)
}

func (o OddRs) EvenRs() EvenRs {
	return castAs(o, OddR.EvenR)
}

func (o OddRs) OddQs() OddQs {
	return castAs(o, OddR.OddQ)
}

func (o OddRs) OddRs() OddRs {
	return o
}

func (o OddRs) Copy() OddRs {
	return slices.Clone(o)
}

func (o OddRs) Sort() OddRs {
	return o.Cubes().Sort().OddRs()
}

func (o OddRs) UnionWith(other OddRs) OddRs {
	return o.Cubes().UnionWith(other.Cubes()).OddRs()
}

func (o OddRs) IntersectWith(other OddRs) OddRs {
	return o.Cubes().IntersectWith(other.Cubes()).OddRs()
}

func (o OddRs) DifferenceWith(other OddRs) OddRs {
	return o.Cubes().DifferenceWith(other.Cubes()).OddRs()
}

func (o OddRs) Rotate(center OddR, angle int) OddRs {
	return o.Cubes().Rotate(center.Cube(), angle).OddRs()
}
