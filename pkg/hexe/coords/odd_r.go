package coords

import "github.com/legendary-code/hexe/pkg/hexe/consts"

type oddR [2]int

func OddR(q int, r int) CoordQR {
	return oddR{q, r}
}

func (o oddR) Axial() CoordQR {
	return Axial(o[0]-(o[1]-(o[1]&1))/2, o[1])
}

func (o oddR) Cube() CoordQRS {
	return o.Axial().Cube()
}

func (o oddR) OddR() CoordQR {
	return o.Axial().OddR()
}

func (o oddR) EvenR() CoordQR {
	return o.Axial().EvenR()
}

func (o oddR) OddQ() CoordQR {
	return o.Axial().OddQ()
}

func (o oddR) EvenQ() CoordQR {
	return o.Axial().EvenQ()
}

func (o oddR) DoubleWidth() CoordQR {
	return o.Axial().DoubleWidth()
}

func (o oddR) DoubleHeight() CoordQR {
	return o.Axial().DoubleHeight()
}

func (o oddR) Q() int {
	return o[0]
}

func (o oddR) R() int {
	return o[1]
}

func (o oddR) Unpack() (int, int) {
	return o[0], o[1]
}

func (o oddR) Neighbors() [consts.Sides]CoordQR {
	neighbors := o.Axial().Neighbors()
	for i, neighbor := range neighbors {
		neighbors[i] = neighbor.OddR()
	}
	return neighbors
}
