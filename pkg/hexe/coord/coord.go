package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

type Coord interface {
	Type() consts.CoordType
	Convert(typ consts.CoordType) Coord
	Axial() Axial
	Cube() Cube
	OddR() OddR
	EvenR() EvenR
	OddQ() OddQ
	EvenQ() EvenQ
	DoubleWidth() DoubleWidth
	DoubleHeight() DoubleHeight
}

type TCoords[T Coord] interface {
}

type Predicate[T Coord] func(coord T) bool

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
	Sorted() TS
	UnionWith(other TS) TS
	IntersectWith(other TS) TS
	DifferenceWith(other TS) TS
}

type QR interface {
	Coord
	Q() int
	R() int
	Unpack() (int, int)
}

type QRCoord[T Coord, TS Coords[T, TS]] interface {
	QR
	Neighbors() TS
	DiagonalNeighbors() TS
	DistanceTo(other T) int
	LineTo(other T) TS
	MovementRange(n int) TS
	FloodFill(n int, blocked Predicate[T]) TS
}

type QRS interface {
	Coord
	Q() int
	R() int
	S() int
	Unpack() (int, int, int)
}

type QRSCoord[T Coord, TS Coords[T, TS]] interface {
	QRS
	Neighbors() TS
	DiagonalNeighbors() TS
	DistanceTo(other T) int
	LineTo(other T) TS
	MovementRange(n int) TS
	FloodFill(n int, blocked Predicate[T]) TS
}

func castAs[F Coord, T Coord](values []F, convertFunc func(F) T) []T {
	result := make([]T, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = convertFunc(values[i])
	}
	return result
}

func convert[F Coord](value F, typ consts.CoordType) Coord {
	switch typ {
	case consts.Axial:
		return value.Axial()
	case consts.Cube:
		return value.Cube()
	case consts.DoubleHeight:
		return value.DoubleHeight()
	case consts.DoubleWidth:
		return value.DoubleWidth()
	case consts.EvenQ:
		return value.EvenQ()
	case consts.EvenR:
		return value.EvenR()
	case consts.OddQ:
		return value.OddQ()
	case consts.OddR:
		return value.OddR()
	}

	panic(fmt.Sprintf("unsupported coord type: %+v", typ))
}

func toCoords[T Coord, TS Coords[T, TS]](coords TS) []Coord {
	cs := make([]Coord, len(coords))
	for i, c := range coords {
		cs[i] = c
	}
	return cs
}
