package coord

type OddQs []OddQ

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
