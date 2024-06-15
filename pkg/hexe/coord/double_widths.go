package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type DoubleWidths []DoubleWidth

func (d DoubleWidths) Type() consts.CoordType {
	return consts.DoubleWidth
}

func (d DoubleWidths) Coords() []Coord {
	return toCoords(d)
}

func (d DoubleWidths) Axials() Axials {
	return castAs(d, DoubleWidth.Axial)
}

func (d DoubleWidths) Cubes() Cubes {
	return castAs(d, DoubleWidth.Cube)
}

func (d DoubleWidths) DoubleWidths() DoubleWidths {
	return d
}

func (d DoubleWidths) DoubleHeights() DoubleHeights {
	return castAs(d, DoubleWidth.DoubleHeight)
}

func (d DoubleWidths) EvenQs() EvenQs {
	return castAs(d, DoubleWidth.EvenQ)
}

func (d DoubleWidths) EvenRs() EvenRs {
	return castAs(d, DoubleWidth.EvenR)
}

func (d DoubleWidths) OddQs() OddQs {
	return castAs(d, DoubleWidth.OddQ)
}

func (d DoubleWidths) OddRs() OddRs {
	return castAs(d, DoubleWidth.OddR)
}

func (d DoubleWidths) Copy() DoubleWidths {
	return slices.Clone(d)
}

func (d DoubleWidths) Sorted() DoubleWidths {
	return d.Axials().Sorted().DoubleWidths()
}

func (d DoubleWidths) UnionWith(other DoubleWidths) DoubleWidths {
	return d.Axials().UnionWith(other.Axials()).DoubleWidths()
}

func (d DoubleWidths) IntersectWith(other DoubleWidths) DoubleWidths {
	return d.Axials().IntersectWith(other.Axials()).DoubleWidths()
}

func (d DoubleWidths) DifferenceWith(other DoubleWidths) DoubleWidths {
	return d.Axials().DifferenceWith(other.Axials()).DoubleWidths()
}
