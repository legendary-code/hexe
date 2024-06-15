package coord

type EvenQs []EvenQ

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
