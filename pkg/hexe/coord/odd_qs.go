package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type OddQs []OddQ

func (o OddQs) Type() consts.CoordType {
	return consts.OddQ
}

func (o OddQs) Coords() []Coord {
	return toCoords(o)
}

func (o OddQs) Axials() Axials {
	return castAs(o, OddQ.Axial)
}

func (o OddQs) Cubes() Cubes {
	return castAs(o, OddQ.Cube)
}

func (o OddQs) DoubleWidths() DoubleWidths {
	return castAs(o, OddQ.DoubleWidth)
}

func (o OddQs) DoubleHeights() DoubleHeights {
	return castAs(o, OddQ.DoubleHeight)
}

func (o OddQs) EvenQs() EvenQs {
	return castAs(o, OddQ.EvenQ)
}

func (o OddQs) EvenRs() EvenRs {
	return castAs(o, OddQ.EvenR)
}

func (o OddQs) OddQs() OddQs {
	return o
}

func (o OddQs) OddRs() OddRs {
	return castAs(o, OddQ.OddR)
}

func (o OddQs) Copy() OddQs {
	return slices.Clone(o)
}

func (o OddQs) Sorted() OddQs {
	return o.Axials().Sorted().OddQs()
}

func (o OddQs) UnionWith(other OddQs) OddQs {
	return o.Axials().UnionWith(other.Axials()).OddQs()
}

func (o OddQs) IntersectWith(other OddQs) OddQs {
	return o.Axials().IntersectWith(other.Axials()).OddQs()
}

func (o OddQs) DifferenceWith(other OddQs) OddQs {
	return o.Axials().DifferenceWith(other.Axials()).OddQs()
}
