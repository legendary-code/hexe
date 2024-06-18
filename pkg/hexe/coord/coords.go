package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
)

// Coords represents an untyped set of coordinates
type Coords interface {
	Type() consts.CoordType
	Convert(typ consts.CoordType) Coords
	Axials() *Axials
	Cubes() *Cubes
	OddRs() *OddRs
	EvenRs() *EvenRs
	OddQs() *OddQs
	EvenQs() *EvenQs
	DoubleWidths() *DoubleWidths
	DoubleHeights() *DoubleHeights
}

type CCoords interface {
	Coords
	*Axials | *Cubes | *DoubleHeights | *DoubleWidths | *EvenQs | *EvenRs | *OddQs | *OddRs
}

type TCoords[T CCoord, TS CCoords] interface {
	Coords
	OrderedSet[T, TS]
	Rotate(center T, angle int) TS
	ReflectQ() TS
	ReflectR() TS
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

func NewCoords(coords []Coord) Coords {
	if len(coords) == 0 {
		return NewEmpty()
	}

	a := make([]Axial, len(coords))
	typ := coords[0].Type()

	for i, coord := range coords {
		if coord.Type() != typ {
			panic(fmt.Sprintf("cannot create Coords from mixed Coord types: %s and %s", coord.Type().Name(), typ.Name()))
		}
		a[i] = coord.Axial()
		typ = coord.Type()
	}

	return NewAxials(a...).Convert(typ)
}
