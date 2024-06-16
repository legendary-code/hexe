package coord

func (o OddQ) Axial() Axial {
	return NewAxial(o[0], o[1]-(o[0]-(o[0]&1))/2)
}

func (o OddQ) Cube() Cube {
	return o.Axial().Cube()
}

func (o OddQ) OddR() OddR {
	return o.Axial().OddR()
}

func (o OddQ) EvenR() EvenR {
	return o.Axial().EvenR()
}

func (o OddQ) OddQ() OddQ {
	return o
}

func (o OddQ) EvenQ() EvenQ {
	return o.Axial().EvenQ()
}

func (o OddQ) DoubleWidth() DoubleWidth {
	return o.Axial().DoubleWidth()
}

func (o OddQ) DoubleHeight() DoubleHeight {
	return o.Axial().DoubleHeight()
}
