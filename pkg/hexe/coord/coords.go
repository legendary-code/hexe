package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type Coords interface {
	Type() consts.CoordType
	Convert(typ consts.CoordType) Coords
	ToSlice() []Coord
	Axials() Axials
	Cubes() Cubes
	OddRs() OddRs
	EvenRs() EvenRs
	OddQs() OddQs
	EvenQs() EvenQs
	DoubleWidths() DoubleWidths
	DoubleHeights() DoubleHeights
}

type TCoordSlice[T Coord] interface {
	~[]T
}

type TCoords[T Coord, TS TCoordSlice[T]] interface {
	~[]T
	Copy() TS
	Sort() TS
	UnionWith(other TS) TS
	IntersectWith(other TS) TS
	DifferenceWith(other TS) TS
	Rotate(center T, angle int) TS
}

func castAs[F Coord, T Coord](values []F, convertFunc func(F) T) []T {
	result := make([]T, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = convertFunc(values[i])
	}
	return result
}

func toCoords[T Coord, TS TCoords[T, TS]](coords TS) []Coord {
	cs := make([]Coord, len(coords))
	for i, c := range coords {
		cs[i] = c
	}
	return cs
}

func convertCoords[F Coords](values F, typ consts.CoordType) Coords {
	switch typ {
	case consts.Axial:
		return values.Axials()
	case consts.Cube:
		return values.Cubes()
	case consts.DoubleHeight:
		return values.DoubleHeights()
	case consts.DoubleWidth:
		return values.DoubleWidths()
	case consts.EvenQ:
		return values.EvenQs()
	case consts.EvenR:
		return values.EvenRs()
	case consts.OddQ:
		return values.OddQs()
	case consts.OddR:
		return values.OddRs()
	default:
		panic(fmt.Sprintf("unsupported coord type: %+v", typ))
	}
}

func NewCoords(coords []Coord) Coords {
	if len(coords) == 0 {
		return Empty{}
	}

	a := make(Axials, len(coords))
	typ := coords[0].Type()

	for i, coord := range coords {
		if coord.Type() != typ {
			panic(fmt.Sprintf("cannot create Coords from mixed Coord types: %s and %s", coord.Type().Name(), typ.Name()))
		}
		a[i] = coord.Axial()
		typ = coord.Type()
	}

	return a.Convert(typ)
}
