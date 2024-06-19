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
	// Type returns the coordinate system type for the coordinate
	Type() consts.CoordType

	// Convert converts the coordinate to another coordinate system
	Convert(typ consts.CoordType) Coord

	// Axial converts the coordinate to an axial coordinate
	Axial() Axial

	// Cube converts the coordinate to a cube coordinate
	Cube() Cube

	// OddR converts the coordinate to an odd-r coordinate
	OddR() OddR

	// EvenR converts the coordinate to an even-r coordinate
	EvenR() EvenR

	// OddQ converts the coordinate to an odd-q coordinate
	OddQ() OddQ

	// EvenQ converts the coordinate to an even-q coordinate
	EvenQ() EvenQ

	// DoubleWidth converts the coordinate to a double-width coordinate
	DoubleWidth() DoubleWidth

	// DoubleHeight converts the coordinate to a double-height coordinate
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
	// Neighbor returns a neighboring coordinate for a specific angle [0..6]
	Neighbor(angle int) T

	// Add returns the sum of the current coordinate and another coordinate
	Add(other T) T

	// Scale multiplies the current coordinate by some factor and returns the result
	Scale(factor int) T

	// Neighbors returns all neighboring coordinates for the current coordinate
	Neighbors() TS

	// DiagonalNeighbors returns all diagonally neighboring coordinates for the current coordinate
	DiagonalNeighbors() TS

	// DistanceTo returns the distance from the current coordinate to another coordinate
	DistanceTo(other T) int

	// LineTo returns a set of coordinates that form a line from the current coordinate to the given coordinate
	LineTo(other T) TS

	// TraceTo returns a set of coordinates that form a line with collision detection from the current coordinate to the
	// given coordinate
	TraceTo(other T, blocked Predicate[T]) TS

	// MovementRange returns all the coordinates reachable from the current coordinate within n steps
	MovementRange(n int) TS

	// FloodFill returns all the coordinates reachable from the current coordinate within n steps with collision taken
	// into account
	FloodFill(n int, blocked Predicate[T]) TS

	// Rotate rotates the current coordinate around a given center by multiples of 60 degrees. A negative angle
	// indicates counter-clockwise.
	Rotate(center T, angle int) T

	// ReflectQ returns the mirror coordinate around the Q-axis
	ReflectQ() T

	// ReflectR returns the mirror coordinate around the R-axis
	ReflectR() T

	// ReflectS returns the mirror coordinate around the S-axis
	ReflectS() T

	// Ring returns a set of coordinates that form a ring of size n around the current coordinate
	Ring(radius int) TS

	// FieldOfView does simple ray tracing to return set of coordinates visible from the current coordinate, at most
	// some given radius away and taking into account collision
	FieldOfView(radius int, blocked Predicate[T]) TS

	// FindPathBFS returns the shortest path to the target, taking into account collision and not exceeding the specified
	// maximum distance
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
