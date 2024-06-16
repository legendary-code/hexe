package coord

func (d DoubleHeight) Axial() Axial {
	return NewAxial(d[0], (d[1]-d[0])/2)
}

func (d DoubleHeight) Cube() Cube {
	return d.Axial().Cube()
}

func (d DoubleHeight) OddR() OddR {
	return d.Axial().OddR()
}

func (d DoubleHeight) EvenR() EvenR {
	return d.Axial().EvenR()
}

func (d DoubleHeight) OddQ() OddQ {
	return d.Axial().OddQ()
}

func (d DoubleHeight) EvenQ() EvenQ {
	return d.Axial().EvenQ()
}

func (d DoubleHeight) DoubleWidth() DoubleWidth {
	return d.Axial().DoubleWidth()
}

func (d DoubleHeight) DoubleHeight() DoubleHeight {
	return d
}
