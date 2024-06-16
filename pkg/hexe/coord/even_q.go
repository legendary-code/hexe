package coord

func (e EvenQ) Axial() Axial {
	return NewAxial(e[0], e[1]-(e[0]+(e[0]&1))/2)
}

func (e EvenQ) Cube() Cube {
	return e.Axial().Cube()
}

func (e EvenQ) OddR() OddR {
	return e.Axial().OddR()
}

func (e EvenQ) EvenR() EvenR {
	return e.Axial().EvenR()
}

func (e EvenQ) OddQ() OddQ {
	return e.Axial().OddQ()
}

func (e EvenQ) EvenQ() EvenQ {
	return e
}

func (e EvenQ) DoubleWidth() DoubleWidth {
	return e.Axial().DoubleWidth()
}

func (e EvenQ) DoubleHeight() DoubleHeight {
	return e.Axial().DoubleHeight()
}
