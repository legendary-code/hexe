package coords

import "github.com/legendary-code/hexe/pkg/hexe/consts"

type evenQ [2]int

func EvenQ(q int, r int) CoordQR {
	return evenQ{q, r}
}

func (e evenQ) Axial() CoordQR {
	return Axial(e[0], e[1]-(e[0]+(e[0]&1))/2)
}

func (e evenQ) Cube() CoordQRS {
	return e.Axial().Cube()
}

func (e evenQ) OddR() CoordQR {
	return e.Axial().OddR()
}

func (e evenQ) EvenR() CoordQR {
	return e.Axial().EvenR()
}

func (e evenQ) OddQ() CoordQR {
	return e.Axial().OddQ()
}

func (e evenQ) EvenQ() CoordQR {
	return e.Axial().EvenQ()
}

func (e evenQ) DoubleWidth() CoordQR {
	return e.Axial().DoubleWidth()
}

func (e evenQ) DoubleHeight() CoordQR {
	return e.Axial().DoubleHeight()
}

func (e evenQ) Q() int {
	return e[0]
}

func (e evenQ) R() int {
	return e[1]
}

func (e evenQ) Unpack() (int, int) {
	return e[0], e[1]
}

func (e evenQ) Neighbors() [consts.Sides]CoordQR {
	neighbors := e.Axial().Neighbors()
	for i, neighbor := range neighbors {
		neighbors[i] = neighbor.EvenQ()
	}
	return neighbors
}
