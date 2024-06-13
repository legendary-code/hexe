package coord

func (d DoubleWidth) Axial() Axial {
	return NewAxial((d[0]-d[1])/2, d[1])
}

func (d DoubleWidth) Cube() Cube {
	return d.Axial().Cube()
}

func (d DoubleWidth) OddR() OddR {
	return d.Axial().OddR()
}

func (d DoubleWidth) EvenR() EvenR {
	return d.Axial().EvenR()
}

func (d DoubleWidth) OddQ() OddQ {
	return d.Axial().OddQ()
}

func (d DoubleWidth) EvenQ() EvenQ {
	return d.Axial().EvenQ()
}

func (d DoubleWidth) DoubleWidth() DoubleWidth {
	return d
}

func (d DoubleWidth) DoubleHeight() DoubleHeight {
	return d.Axial().DoubleHeight()
}
