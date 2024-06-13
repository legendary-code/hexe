package coords

type axial [2]int

func Axial(q int, r int) CoordQR {
	return axial{q, r}
}

func (a axial) Axial() CoordQR {
	return a
}

func (a axial) Cube() CoordQRS {
	return Cube(a[0], a[1], -a[0]-a[1])
}

func (a axial) OddR() CoordQR {
	return OddR(a[0]+(a[1]-a[1]&1)/2, a[1])
}

func (a axial) EvenR() CoordQR {
	return EvenR(a[0]+(a[1]+a[1]&1)/2, a[1])
}

func (a axial) OddQ() CoordQR {
	return OddQ(a[0], a[1]+(a[0]-(a[0]&1))/2)
}

func (a axial) EvenQ() CoordQR {
	return EvenQ(a[0], a[1]+(a[0]+(a[0]&1))/2)
}

func (a axial) DoubleWidth() CoordQR {
	return DoubleWidth(2*a[0]+a[1], a[1])
}

func (a axial) DoubleHeight() CoordQR {
	return DoubleHeight(a[0], 2*a[1]+a[0])
}

func (a axial) Q() int {
	return a[0]
}

func (a axial) R() int {
	return a[1]
}

func (a axial) Unpack() (int, int) {
	return a[0], a[1]
}
