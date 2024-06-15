package coord

type Axials []Axial

func (a Axials) Axials() Axials {
	return a
}

func (a Axials) Cubes() Cubes {
	return castAs(a, Axial.Cube)
}

func (a Axials) DoubleWidths() DoubleWidths {
	return castAs(a, Axial.DoubleWidth)
}

func (a Axials) DoubleHeights() DoubleHeights {
	return castAs(a, Axial.DoubleHeight)
}

func (a Axials) EvenQs() EvenQs {
	return castAs(a, Axial.EvenQ)
}

func (a Axials) EvenRs() EvenRs {
	return castAs(a, Axial.EvenR)
}

func (a Axials) OddQs() OddQs {
	return castAs(a, Axial.OddQ)
}

func (a Axials) OddRs() OddRs {
	return castAs(a, Axial.OddR)
}
