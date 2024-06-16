package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type DoubleHeights []DoubleHeight

func (d DoubleHeights) Type() consts.CoordType {
	return consts.DoubleHeight
}

func (d DoubleHeights) Convert(typ consts.CoordType) Coords {
	return convertCoords(d, typ)
}

func (d DoubleHeights) ToSlice() []Coord {
	return toCoords(d)
}

func (d DoubleHeights) Axials() Axials {
	return castAs(d, DoubleHeight.Axial)
}

func (d DoubleHeights) Cubes() Cubes {
	return castAs(d, DoubleHeight.Cube)
}

func (d DoubleHeights) DoubleWidths() DoubleWidths {
	return castAs(d, DoubleHeight.DoubleWidth)
}

func (d DoubleHeights) DoubleHeights() DoubleHeights {
	return d
}

func (d DoubleHeights) EvenQs() EvenQs {
	return castAs(d, DoubleHeight.EvenQ)
}

func (d DoubleHeights) EvenRs() EvenRs {
	return castAs(d, DoubleHeight.EvenR)
}

func (d DoubleHeights) OddQs() OddQs {
	return castAs(d, DoubleHeight.OddQ)
}

func (d DoubleHeights) OddRs() OddRs {
	return castAs(d, DoubleHeight.OddR)
}

func (d DoubleHeights) Copy() DoubleHeights {
	return slices.Clone(d)
}

func (d DoubleHeights) Sort() DoubleHeights {
	return d.Cubes().Sort().DoubleHeights()
}

func (d DoubleHeights) UnionWith(other DoubleHeights) DoubleHeights {
	return d.Cubes().UnionWith(other.Cubes()).DoubleHeights()
}

func (d DoubleHeights) IntersectWith(other DoubleHeights) DoubleHeights {
	return d.Cubes().IntersectWith(other.Cubes()).DoubleHeights()
}

func (d DoubleHeights) DifferenceWith(other DoubleHeights) DoubleHeights {
	return d.Cubes().DifferenceWith(other.Cubes()).DoubleHeights()
}

func (d DoubleHeights) Rotate(center DoubleHeight, angle int) DoubleHeights {
	return d.Cubes().Rotate(center.Cube(), angle).DoubleHeights()
}

func (d DoubleHeights) ReflectQ() DoubleHeights {
	return d.Cubes().ReflectQ().DoubleHeights()
}

func (d DoubleHeights) ReflectR() DoubleHeights {
	return d.Cubes().ReflectR().DoubleHeights()
}

func (d DoubleHeights) ReflectS() DoubleHeights {
	return d.Cubes().ReflectS().DoubleHeights()
}
