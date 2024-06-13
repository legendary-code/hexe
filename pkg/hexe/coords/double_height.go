package coords

type doubleHeight [2]int

func DoubleHeight(q int, r int) CoordQR {
	return doubleHeight{q, r}
}

func (d doubleHeight) Axial() CoordQR {
	return Axial(d[0], (d[1]-d[0])/2)
}

func (d doubleHeight) Cube() CoordQRS {
	return d.Axial().Cube()
}

func (d doubleHeight) OddR() CoordQR {
	return d.Axial().OddR()
}

func (d doubleHeight) EvenR() CoordQR {
	return d.Axial().EvenR()
}

func (d doubleHeight) OddQ() CoordQR {
	return d.Axial().OddQ()
}

func (d doubleHeight) EvenQ() CoordQR {
	return d.Axial().EvenQ()
}

func (d doubleHeight) DoubleWidth() CoordQR {
	return d.Axial().DoubleWidth()
}

func (d doubleHeight) DoubleHeight() CoordQR {
	return d
}

func (d doubleHeight) Q() int {
	return d[0]
}

func (d doubleHeight) R() int {
	return d[1]
}

func (d doubleHeight) Unpack() (int, int) {
	return d[0], d[1]
}
