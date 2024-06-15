package coord

import "slices"

type EvenRs []EvenR

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

func (e EvenRs) Sorted() EvenRs {
	return e.Axials().Sorted().EvenRs()
}

func (e EvenRs) UnionWith(other EvenRs) EvenRs {
	return e.Axials().UnionWith(other.Axials()).EvenRs()
}

func (e EvenRs) IntersectWith(other EvenRs) EvenRs {
	return e.Axials().IntersectWith(other.Axials()).EvenRs()
}

func (e EvenRs) DifferenceWith(other EvenRs) EvenRs {
	return e.Axials().DifferenceWith(other.Axials()).EvenRs()
}
