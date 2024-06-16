package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type EvenQs []EvenQ

func (e EvenQs) Type() consts.CoordType {
	return consts.EvenQ
}

func (e EvenQs) Convert(typ consts.CoordType) Coords {
	return convertCoords(e, typ)
}

func (e EvenQs) ToSlice() []Coord {
	return toCoords(e)
}

func (e EvenQs) Axials() Axials {
	return castAs(e, EvenQ.Axial)
}

func (e EvenQs) Cubes() Cubes {
	return castAs(e, EvenQ.Cube)
}

func (e EvenQs) DoubleWidths() DoubleWidths {
	return castAs(e, EvenQ.DoubleWidth)
}

func (e EvenQs) DoubleHeights() DoubleHeights {
	return castAs(e, EvenQ.DoubleHeight)
}

func (e EvenQs) EvenQs() EvenQs {
	return e
}

func (e EvenQs) EvenRs() EvenRs {
	return castAs(e, EvenQ.EvenR)
}

func (e EvenQs) OddQs() OddQs {
	return castAs(e, EvenQ.OddQ)
}

func (e EvenQs) OddRs() OddRs {
	return castAs(e, EvenQ.OddR)
}

func (e EvenQs) Copy() EvenQs {
	return slices.Clone(e)
}

func (e EvenQs) Sort() EvenQs {
	return e.Cubes().Sort().EvenQs()
}

func (e EvenQs) UnionWith(other EvenQs) EvenQs {
	return e.Cubes().UnionWith(other.Cubes()).EvenQs()
}

func (e EvenQs) IntersectWith(other EvenQs) EvenQs {
	return e.Cubes().IntersectWith(other.Cubes()).EvenQs()
}

func (e EvenQs) DifferenceWith(other EvenQs) EvenQs {
	return e.Cubes().DifferenceWith(other.Cubes()).EvenQs()
}

func (e EvenQs) Rotate(center EvenQ, angle int) EvenQs {
	return e.Cubes().Rotate(center.Cube(), angle).EvenQs()
}

func (e EvenQs) ReflectQ() EvenQs {
	return e.Cubes().ReflectQ().EvenQs()
}

func (e EvenQs) ReflectR() EvenQs {
	return e.Cubes().ReflectR().EvenQs()
}

func (e EvenQs) ReflectS() EvenQs {
	return e.Cubes().ReflectS().EvenQs()
}
