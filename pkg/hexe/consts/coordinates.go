package consts

import "fmt"

// CoordType represents an enum of supported hexagonal coordinate system types
type CoordType int

const (
	Axial CoordType = iota
	Cube
	DoubleHeight
	DoubleWidth
	EvenQ
	EvenR
	OddQ
	OddR
	Unknown CoordType = -1
)

// Name returns the display name for the coordinate system
func (c CoordType) Name() string {
	switch c {
	case Axial:
		return "Axial"
	case Cube:
		return "Cube"
	case DoubleHeight:
		return "DoubleHeight"
	case DoubleWidth:
		return "DoubleWidth"
	case EvenQ:
		return "EvenQ"
	case EvenR:
		return "EvenR"
	case OddQ:
		return "OddQ"
	case OddR:
		return "OddR"
	case Unknown:
		return "Unknown"
	default:
		panic(fmt.Sprintf("unknown coord type: %+v", c))
	}
}

// CoordinateTypes returns a slice of supported coordinate types
func CoordinateTypes() []CoordType {
	return []CoordType{
		Axial,
		Cube,
		DoubleHeight,
		DoubleWidth,
		EvenQ,
		EvenR,
		OddQ,
		OddR,
	}
}
