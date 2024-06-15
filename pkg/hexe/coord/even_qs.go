package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type EvenQs []EvenQ

func (e EvenQs) Type() consts.CoordType {
	return consts.EvenQ
}

func (e EvenQs) Coords() []Coord {
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

func (e EvenQs) Sorted() EvenQs {
	return e.Axials().Sorted().EvenQs()
}

func (e EvenQs) UnionWith(other EvenQs) EvenQs {
	return e.Axials().UnionWith(other.Axials()).EvenQs()
}

func (e EvenQs) IntersectWith(other EvenQs) EvenQs {
	return e.Axials().IntersectWith(other.Axials()).EvenQs()
}

func (e EvenQs) DifferenceWith(other EvenQs) EvenQs {
	return e.Axials().DifferenceWith(other.Axials()).EvenQs()
}
