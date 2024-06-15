package coord

type OddRs []OddR

func (o OddRs) Axials() Axials {
	return castAs(o, OddR.Axial)
}

func (o OddRs) Cubes() Cubes {
	return castAs(o, OddR.Cube)
}

func (o OddRs) DoubleWidths() DoubleWidths {
	return castAs(o, OddR.DoubleWidth)
}

func (o OddRs) DoubleHeights() DoubleHeights {
	return castAs(o, OddR.DoubleHeight)
}

func (o OddRs) EvenQs() EvenQs {
	return castAs(o, OddR.EvenQ)
}

func (o OddRs) EvenRs() EvenRs {
	return castAs(o, OddR.EvenR)
}

func (o OddRs) OddQs() OddQs {
	return castAs(o, OddR.OddQ)
}

func (o OddRs) OddRs() OddRs {
	return o
}
