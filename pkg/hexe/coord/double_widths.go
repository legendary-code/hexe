package coord

type DoubleWidths []DoubleWidth

func (d DoubleWidths) Axials() Axials {
	return castAs(d, DoubleWidth.Axial)
}

func (d DoubleWidths) Cubes() Cubes {
	return castAs(d, DoubleWidth.Cube)
}

func (d DoubleWidths) DoubleWidths() DoubleWidths {
	return d
}

func (d DoubleWidths) DoubleHeights() DoubleHeights {
	return castAs(d, DoubleWidth.DoubleHeight)
}

func (d DoubleWidths) EvenQs() EvenQs {
	return castAs(d, DoubleWidth.EvenQ)
}

func (d DoubleWidths) EvenRs() EvenRs {
	return castAs(d, DoubleWidth.EvenR)
}

func (d DoubleWidths) OddQs() OddQs {
	return castAs(d, DoubleWidth.OddQ)
}

func (d DoubleWidths) OddRs() OddRs {
	return castAs(d, DoubleWidth.OddR)
}
