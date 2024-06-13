package coords

import "github.com/legendary-code/hexe/pkg/hexe/consts"

type oddQ [2]int

func OddQ(q int, r int) CoordQR {
	return oddQ{q, r}
}

func (o oddQ) Axial() CoordQR {
	return Axial(o[0], o[1]-(o[0]-(o[0]&1))/2)
}

func (o oddQ) Cube() CoordQRS {
	return o.Axial().Cube()
}

func (o oddQ) OddR() CoordQR {
	return o.Axial().OddR()
}

func (o oddQ) EvenR() CoordQR {
	return o.Axial().EvenR()
}

func (o oddQ) OddQ() CoordQR {
	return o.Axial().OddQ()
}

func (o oddQ) EvenQ() CoordQR {
	return o.Axial().EvenQ()
}

func (o oddQ) DoubleWidth() CoordQR {
	return o.Axial().DoubleWidth()
}

func (o oddQ) DoubleHeight() CoordQR {
	return o.Axial().DoubleHeight()
}

func (o oddQ) Q() int {
	return o[0]
}

func (o oddQ) R() int {
	return o[1]
}

func (o oddQ) Unpack() (int, int) {
	return o[0], o[1]
}

func (o oddQ) Neighbors() [consts.Sides]CoordQR {
	neighbors := o.Axial().Neighbors()
	for i, neighbor := range neighbors {
		neighbors[i] = neighbor.OddQ()
	}
	return neighbors
}
