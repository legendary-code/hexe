package consts

import "fmt"

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
)

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
	default:
		panic(fmt.Sprintf("unknown coord type: %+v", c))
	}
}

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
