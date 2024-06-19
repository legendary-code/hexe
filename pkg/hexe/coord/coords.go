package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

// Coords represents an untyped set of coordinates
type Coords interface {
	// Type returns the coordinate system type for the set of coordinates
	Type() consts.CoordType

	// Convert converts the set of coordinates to another coordinate system
	Convert(typ consts.CoordType) Coords

	// Axials converts the set of coordinates to a set of axial coordinates
	Axials() *Axials

	// Cubes converts the set of coordinates to a set of cube coordinates
	Cubes() *Cubes

	// OddRs converts the set of coordinates to a set of odd-r coordinates
	OddRs() *OddRs

	// EvenRs converts the set of coordinates to a set of even-r coordinates
	EvenRs() *EvenRs

	// OddQs converts the set of coordinates to a set of odd-q coordinates
	OddQs() *OddQs

	// EvenQs converts the set of coordinates to a set of even-q coordinates
	EvenQs() *EvenQs

	// DoubleWidths converts the set of coordinates to a set of double-width coordinates
	DoubleWidths() *DoubleWidths

	// DoubleHeights converts the set of coordinates to a set of double-height coordinates
	DoubleHeights() *DoubleHeights
}

// CCoords represents a constrained untyped set of coordinates, to ensure that they are a specific type
type CCoords interface {
	Coords
	*Axials | *Cubes | *DoubleHeights | *DoubleWidths | *EvenQs | *EvenRs | *OddQs | *OddRs
}

// TCoords represents a typed set of coordinates
type TCoords[T CCoord, TS CCoords] interface {
	Coords
	OrderedSet[T, TS]

	// Rotate rotates the set of coordinates around a given center by multiples of 60 degrees. A negative angle
	// indicates counter-clockwise.
	Rotate(center T, angle int) TS

	// ReflectQ returns the mirror set of coordinates around the Q-axis
	ReflectQ() TS

	// ReflectR returns the mirror set of coordinates around the R-axis
	ReflectR() TS

	// ReflectS returns the mirror set of coordinates around the S-axis
	ReflectS() TS
}

func castAs[F comparable, FS any, T comparable, TS any](
	values OrderedSet[F, FS], convertFunc func(F) T, ctorFunc func(...T) TS,
) TS {
	result := make([]T, values.Size())
	for i := values.Iterator(); i.Next(); {
		result[i.Index()] = convertFunc(i.Item())
	}
	return ctorFunc(result...)
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
