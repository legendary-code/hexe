package coord

import "github.com/legendary-code/hexe/pkg/hexe/consts"

type TCoords[T Coord] interface {
}

type Coords[T Coord, TS TCoords[T]] interface {
	~[]T
	Type() consts.CoordType
	Coords() []Coord
	Axials() Axials
	Cubes() Cubes
	OddRs() OddRs
	EvenRs() EvenRs
	OddQs() OddQs
	EvenQs() EvenQs
	DoubleWidths() DoubleWidths
	DoubleHeights() DoubleHeights
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

func toCoords[T Coord, TS Coords[T, TS]](coords TS) []Coord {
	cs := make([]Coord, len(coords))
	for i, c := range coords {
		cs[i] = c
	}
	return cs
}
