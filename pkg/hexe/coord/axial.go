package coord

func (a Axial) Axial() Axial {
	return a
}

func (a Axial) Cube() Cube {
	return NewCube(a[0], a[1], -a[0]-a[1])
}

func (a Axial) OddR() OddR {
	return NewOddR(a[0]+(a[1]-a[1]&1)/2, a[1])
}

func (a Axial) EvenR() EvenR {
	return NewEvenR(a[0]+(a[1]+a[1]&1)/2, a[1])
}

func (a Axial) OddQ() OddQ {
	return NewOddQ(a[0], a[1]+(a[0]-(a[0]&1))/2)
}

func (a Axial) EvenQ() EvenQ {
	return NewEvenQ(a[0], a[1]+(a[0]+(a[0]&1))/2)
}

func (a Axial) DoubleWidth() DoubleWidth {
	return NewDoubleWidth(2*a[0]+a[1], a[1])
}

func (a Axial) DoubleHeight() DoubleHeight {
	return NewDoubleHeight(a[0], 2*a[1]+a[0])
}
