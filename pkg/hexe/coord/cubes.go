package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type Cubes struct {
	*orderedSet[Cube, *Cubes]
}

func NewCubes(values ...Cube) *Cubes {
	return &Cubes{
		orderedSet: newOrderedSet[Cube, *Cubes](func() *Cubes {
			return NewCubes()
		}, values...),
	}
}

func (c *Cubes) Type() consts.CoordType {
	return consts.Cube
}

func (c *Cubes) Convert(typ consts.CoordType) Coords {
	return convertCoords(c, typ)
}

func (c *Cubes) Axials() *Axials {
	return castAs[Cube, *Cubes, Axial, *Axials](c, Cube.Axial, NewAxials)
}

func (c *Cubes) Cubes() *Cubes {
	return c
}

func (c *Cubes) DoubleWidths() *DoubleWidths {
	return castAs[Cube, *Cubes, DoubleWidth, *DoubleWidths](c, Cube.DoubleWidth, NewDoubleWidths)
}

func (c *Cubes) DoubleHeights() *DoubleHeights {
	return castAs[Cube, *Cubes, DoubleHeight, *DoubleHeights](c, Cube.DoubleHeight, NewDoubleHeights)
}

func (c *Cubes) EvenQs() *EvenQs {
	return castAs[Cube, *Cubes, EvenQ, *EvenQs](c, Cube.EvenQ, NewEvenQs)
}

func (c *Cubes) EvenRs() *EvenRs {
	return castAs[Cube, *Cubes, EvenR, *EvenRs](c, Cube.EvenR, NewEvenRs)
}

func (c *Cubes) OddQs() *OddQs {
	return castAs[Cube, *Cubes, OddQ, *OddQs](c, Cube.OddQ, NewOddQs)
}

func (c *Cubes) OddRs() *OddRs {
	return castAs[Cube, *Cubes, OddR, *OddRs](c, Cube.OddR, NewOddRs)
}

func (c *Cubes) Rotate(center Cube, angle int) *Cubes {
	coords := make([]Cube, c.Size())
	for i := c.Iterator(); i.Next(); {
		coords[i.Index()] = i.Item().Rotate(center, angle)
	}
	return NewCubes(coords...)
}

func (c *Cubes) ReflectQ() *Cubes {
	coords := make([]Cube, c.Size())
	for i := c.Iterator(); i.Next(); {
		coords[i.Index()] = i.Item().ReflectQ()
	}
	return NewCubes(coords...)
}

func (c *Cubes) ReflectR() *Cubes {
	coords := make([]Cube, c.Size())
	for i := c.Iterator(); i.Next(); {
		coords[i.Index()] = i.Item().ReflectR()
	}
	return NewCubes(coords...)
}

func (c *Cubes) ReflectS() *Cubes {
	coords := make([]Cube, c.Size())
	for i := c.Iterator(); i.Next(); {
		coords[i.Index()] = i.Item().ReflectS()
	}
	return NewCubes(coords...)
}
