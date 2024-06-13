package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

//go:generate go run ../../../internal/hexe/gen/coords
//go:generate go fmt .

// Coord represents an untyped coordinate
type Coord interface {
	fmt.Stringer
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

// CCoord represents a constrained untyped coordinate, to ensure that it is a specific type and comparable
type CCoord interface {
	Coord
	comparable
	Axial | Cube | DoubleHeight | DoubleWidth | EvenQ | EvenR | OddQ | OddR
}

// TCoord represents a typed coordinate
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

// Predicate represents a function that returns true or false for a given coordinate
type Predicate[T Coord] func(coord T) bool

// QRCoord represents a generic coordinate that has Q and R components
type QRCoord interface {
	Coord
	// Q returns the Q component of the coordinate
	Q() int

	// R returns the R component of the coordinate
	R() int

	// Unpack returns all coordinate components for ease of multiple assignments
	Unpack() (int, int)
}

// TQRCoord represents a typed coordinate with Q and R components
type TQRCoord[T CCoord, TS CCoords] interface {
	QRCoord
	TCoord[T, TS]
}

// QRSCoord represents a generic coordinate that has Q, R and S components
type QRSCoord interface {
	Coord
	// Q returns the Q component of the coordinate
	Q() int

	// R returns the R component of the coordinate
	R() int

	// S returns the S component of the coordinate
	S() int

	// Unpack returns all coordinate components for ease of multiple assignments
	Unpack() (int, int, int)
}

// TQRSCoord represents a typed hex coordinate with Q, R and S components
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
