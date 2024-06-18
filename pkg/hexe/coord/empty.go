package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type Empty struct {
	*orderedSet[Coord, *Empty]
}

func NewEmpty() *Empty {
	return &Empty{
		orderedSet: newOrderedSet[Coord, *Empty](NewEmpty),
	}
}

func (e *Empty) Type() consts.CoordType {
	return consts.Unknown
}

func (e *Empty) Convert(_ consts.CoordType) Coords {
	return e
}

func (e *Empty) ToSlice() []Coord {
	return []Coord{}
}

func (e *Empty) Axials() *Axials {
	return NewAxials()
}

func (e *Empty) Cubes() *Cubes {
	return NewCubes()
}

func (e *Empty) DoubleWidths() *DoubleWidths {
	return NewDoubleWidths()
}

func (e *Empty) DoubleHeights() *DoubleHeights {
	return NewDoubleHeights()
}

func (e *Empty) EvenQs() *EvenQs {
	return NewEvenQs()
}

func (e *Empty) EvenRs() *EvenRs {
	return NewEvenRs()
}

func (e *Empty) OddQs() *OddQs {
	return NewOddQs()
}

func (e *Empty) OddRs() *OddRs {
	return NewOddRs()
}

func (e *Empty) Rotate(_ Coord, _ int) *Empty {
	return e
}

func (e *Empty) ReflectQ() *Empty {
	return e
}

func (e *Empty) ReflectR() *Empty {
	return e
}

func (e *Empty) ReflectS() *Empty {
	return e
}
