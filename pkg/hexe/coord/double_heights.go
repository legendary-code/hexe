package coord

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
