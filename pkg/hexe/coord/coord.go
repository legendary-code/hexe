package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

//go:generate go run ../../../internal/hexe/gen/coords
//go:generate go fmt .

// Coord represents an untyped hex coordinate
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

type CCoord interface {
	Coord
	comparable
	Axial | Cube | DoubleHeight | DoubleWidth | EvenQ | EvenR | OddQ | OddR
}

// TCoord represents a type hex coordinate
type TCoord[T CCoord, TS CCoords] interface {
	Coord
	Neighbor(angle int) T
	Add(other T) T
	Scale(factor int) T
	Neighbors() TS
	DiagonalNeighbors() TS
	DistanceTo(other T) int
	LineTo(other T) TS
	TraceTo(other T, blocked Predicate[T]) TS
	MovementRange(n int) TS
	FloodFill(n int, blocked Predicate[T]) TS
	Rotate(center T, angle int) T
	ReflectQ() T
	ReflectR() T
	ReflectS() T
	Ring(radius int) TS
	FieldOfView(radius int, blocked Predicate[T]) TS
	FindPathBFS(target T, maxDistance int, blocked Predicate[T]) TS
}

type Predicate[T Coord] func(coord T) bool

type QRCoord interface {
	Coord
	Q() int
	R() int
	Unpack() (int, int)
}

type TQRCoord[T CCoord, TS CCoords] interface {
	QRCoord
	TCoord[T, TS]
}

type QRSCoord interface {
	Coord
	Q() int
	R() int
	S() int
	Unpack() (int, int, int)
}

type TQRSCoord[T CCoord, TS CCoords] interface {
	QRSCoord
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
