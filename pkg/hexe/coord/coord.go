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

type TCoord[T Coord, TS TCoords[T, TS]] interface {
	Coord
	Neighbors() TS
	DiagonalNeighbors() TS
	DistanceTo(other T) int
	LineTo(other T) TS
	MovementRange(n int) TS
	FloodFill(n int, blocked Predicate[T]) TS
	Rotate(center T, angle int) T
}

type Predicate[T Coord] func(coord T) bool

type QR interface {
	Coord
	Q() int
	R() int
	Unpack() (int, int)
}

type QRCoord[T Coord, TS TCoords[T, TS]] interface {
	QR
	TCoord[T, TS]
}

type QRS interface {
	Coord
	Q() int
	R() int
	S() int
	Unpack() (int, int, int)
}

type QRSCoord[T Coord, TS TCoords[T, TS]] interface {
	QRS
	TCoord[T, TS]
}

func convertCoord[F Coord](value F, typ consts.CoordType) Coord {
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
	default:
		panic(fmt.Sprintf("unsupported coord type: %+v", typ))
	}
}
