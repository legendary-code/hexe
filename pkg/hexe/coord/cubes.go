package coord

type Cubes []Cube

func (c Cubes) Axials() Axials {
	return castAs(c, Cube.Axial)
}

func (c Cubes) Cubes() Cubes {
	return c
}

func (c Cubes) DoubleWidths() DoubleWidths {
	return castAs(c, Cube.DoubleWidth)
}

func (c Cubes) DoubleHeights() DoubleHeights {
	return castAs(c, Cube.DoubleHeight)
}

func (c Cubes) EvenQs() EvenQs {
	return castAs(c, Cube.EvenQ)
}

func (c Cubes) EvenRs() EvenRs {
	return castAs(c, Cube.EvenR)
}

func (c Cubes) OddQs() OddQs {
	return castAs(c, Cube.OddQ)
}

func (c Cubes) OddRs() OddRs {
	return castAs(c, Cube.OddR)
}
