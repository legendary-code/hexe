package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"slices"
)

type Cubes []Cube

func (c Cubes) Type() consts.CoordType {
	return consts.Cube
}

func (c Cubes) Coords() []Coord {
	return toCoords(c)
}

func (c Cubes) Axials() Axials {
	return castAs(c, Cube.Axial)
}

func (c Cubes) Cubes() Cubes {
	return c
}

func (c Cubes) DoubleWidths() DoubleWidths {
	return castAs(c, Cube.DoubleWidth)
}

func (c Cubes) DoubleHeights() DoubleHeights {
	return castAs(c, Cube.DoubleHeight)
}

func (c Cubes) EvenQs() EvenQs {
	return castAs(c, Cube.EvenQ)
}

func (c Cubes) EvenRs() EvenRs {
	return castAs(c, Cube.EvenR)
}

func (c Cubes) OddQs() OddQs {
	return castAs(c, Cube.OddQ)
}

func (c Cubes) OddRs() OddRs {
	return castAs(c, Cube.OddR)
}

func (c Cubes) Copy() Cubes {
	return slices.Clone(c)
}

func (c Cubes) Sorted() Cubes {
	return c.Axials().Sorted().Cubes()
}

func (c Cubes) UnionWith(other Cubes) Cubes {
	return c.Axials().UnionWith(other.Axials()).Cubes()
}

func (c Cubes) IntersectWith(other Cubes) Cubes {
	return c.Axials().IntersectWith(other.Axials()).Cubes()
}

func (c Cubes) DifferenceWith(other Cubes) Cubes {
	return c.Axials().DifferenceWith(other.Axials()).Cubes()
}
