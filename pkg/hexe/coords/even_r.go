package coords

type evenR [2]int

func EvenR(q int, r int) CoordQR {
	return evenR{q, r}
}

func (e evenR) Axial() CoordQR {
	return Axial(e[0]-(e[1]+(e[1]&1))/2, e[1])
}

func (e evenR) Cube() CoordQRS {
	return e.Axial().Cube()
}

func (e evenR) OddR() CoordQR {
	return e.Axial().OddR()
}

func (e evenR) EvenR() CoordQR {
	return e.Axial().EvenR()
}

func (e evenR) OddQ() CoordQR {
	return e.Axial().OddQ()
}

func (e evenR) EvenQ() CoordQR {
	return e.Axial().EvenQ()
}

func (e evenR) DoubleWidth() CoordQR {
	return e.Axial().DoubleWidth()
}

func (e evenR) DoubleHeight() CoordQR {
	return e.Axial().DoubleHeight()
}

func (e evenR) Q() int {
	return e[0]
}

func (e evenR) R() int {
	return e[1]
}

func (e evenR) Unpack() (int, int) {
	return e[0], e[1]
}
