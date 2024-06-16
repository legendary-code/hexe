package coord

func (e EvenR) Axial() Axial {
	return NewAxial(e[0]-(e[1]+(e[1]&1))/2, e[1])
}

func (e EvenR) Cube() Cube {
	return e.Axial().Cube()
}

func (e EvenR) OddR() OddR {
	return e.Axial().OddR()
}

func (e EvenR) EvenR() EvenR {
	return e
}

func (e EvenR) OddQ() OddQ {
	return e.Axial().OddQ()
}

func (e EvenR) EvenQ() EvenQ {
	return e.Axial().EvenQ()
}

func (e EvenR) DoubleWidth() DoubleWidth {
	return e.Axial().DoubleWidth()
}

func (e EvenR) DoubleHeight() DoubleHeight {
	return e.Axial().DoubleHeight()
}
