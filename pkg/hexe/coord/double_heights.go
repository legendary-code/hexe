package coord

import "slices"

type DoubleHeights []DoubleHeight

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

func (d DoubleHeights) Sorted() DoubleHeights {
	return d.Axials().Sorted().DoubleHeights()
}

func (d DoubleHeights) UnionWith(other DoubleHeights) DoubleHeights {
	return d.Axials().UnionWith(other.Axials()).DoubleHeights()
}

func (d DoubleHeights) IntersectWith(other DoubleHeights) DoubleHeights {
	return d.Axials().IntersectWith(other.Axials()).DoubleHeights()
}

func (d DoubleHeights) DifferenceWith(other DoubleHeights) DoubleHeights {
	return d.Axials().DifferenceWith(other.Axials()).DoubleHeights()
}
