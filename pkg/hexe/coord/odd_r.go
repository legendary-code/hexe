package coord

func (o OddR) Axial() Axial {
	return NewAxial(o[0]-(o[1]-(o[1]&1))/2, o[1])
}

func (o OddR) Cube() Cube {
	return o.Axial().Cube()
}

func (o OddR) OddR() OddR {
	return o
}

func (o OddR) EvenR() EvenR {
	return o.Axial().EvenR()
}

func (o OddR) OddQ() OddQ {
	return o.Axial().OddQ()
}

func (o OddR) EvenQ() EvenQ {
	return o.Axial().EvenQ()
}

func (o OddR) DoubleWidth() DoubleWidth {
	return o.Axial().DoubleWidth()
}

func (o OddR) DoubleHeight() DoubleHeight {
	return o.Axial().DoubleHeight()
}
