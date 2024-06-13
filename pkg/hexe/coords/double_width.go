package coords

import "github.com/legendary-code/hexe/pkg/hexe/consts"

type doubleWidth [2]int

func DoubleWidth(q int, r int) CoordQR {
	return doubleWidth{q, r}
}

func (d doubleWidth) Axial() CoordQR {
	return Axial((d[0]-d[1])/2, d[1])
}

func (d doubleWidth) Cube() CoordQRS {
	return d.Axial().Cube()
}

func (d doubleWidth) OddR() CoordQR {
	return d.Axial().OddR()
}

func (d doubleWidth) EvenR() CoordQR {
	return d.Axial().EvenR()
}

func (d doubleWidth) OddQ() CoordQR {
	return d.Axial().OddQ()
}

func (d doubleWidth) EvenQ() CoordQR {
	return d.Axial().EvenQ()
}

func (d doubleWidth) DoubleWidth() CoordQR {
	return d
}

func (d doubleWidth) DoubleHeight() CoordQR {
	return d.Axial().DoubleHeight()
}

func (d doubleWidth) Q() int {
	return d[0]
}

func (d doubleWidth) R() int {
	return d[1]
}

func (d doubleWidth) Unpack() (int, int) {
	return d[0], d[1]
}

func (d doubleWidth) Neighbors() [consts.Sides]CoordQR {
	neighbors := d.Axial().Neighbors()
	for i, neighbor := range neighbors {
		neighbors[i] = neighbor.DoubleWidth()
	}
	return neighbors
}
