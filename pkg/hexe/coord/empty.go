package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type Empty []Coord

func (e Empty) Type() consts.CoordType {
	return consts.Unknown
}

func (e Empty) Convert(typ consts.CoordType) Coords {
	return convertCoords(e, typ)
}

func (e Empty) ToSlice() []Coord {
	return []Coord{}
}

func (e Empty) Axials() Axials {
	return Axials{}
}

func (e Empty) Cubes() Cubes {
	return Cubes{}
}

func (e Empty) DoubleWidths() DoubleWidths {
	return DoubleWidths{}
}

func (e Empty) DoubleHeights() DoubleHeights {
	return DoubleHeights{}
}

func (e Empty) EvenQs() EvenQs {
	return EvenQs{}
}

func (e Empty) EvenRs() EvenRs {
	return EvenRs{}
}

func (e Empty) OddQs() OddQs {
	return OddQs{}
}

func (e Empty) OddRs() OddRs {
	return OddRs{}
}

func (e Empty) Copy() Empty {
	return e
}

func (e Empty) Sort() Empty {
	return e
}

func (e Empty) UnionWith(_ Empty) Empty {
	return e
}

func (e Empty) IntersectWith(_ Empty) Empty {
	return e
}

func (e Empty) DifferenceWith(_ Empty) Empty {
	return e
}

func (e Empty) Rotate(_ Coord, _ int) Empty {
	return e
}
