package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newCoords(coords []Coord) Coords {
	if len(coords) == 0 {
		panic("unexpected empty coordinates")
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

func TestCoords_Type(t *testing.T) {
	for coordType, coords := range testCoordinates {
		t.Run(fmt.Sprintf("Test%ss_Type", coordType.Name()), func(t2 *testing.T) {
			assert.Equal(t2, coordType, newCoords(coords).Type())
		})
	}
}

func TestCoords_Coords(t *testing.T) {
	for _, typeA := range consts.CoordinateTypes() {
		for _, typeB := range consts.CoordinateTypes() {
			if typeA == typeB {
				continue
			}

			t.Run(fmt.Sprintf("Test%ss_%ss", typeA.Name(), typeB.Name()), func(t2 *testing.T) {
				typeACoords := newCoords(testCoordinates[typeA])
				typeBCoords := newCoords(testCoordinates[typeB])

				assert.True(t2, typeBCoords.Axials().Equal(typeACoords.Convert(typeB).Axials()))
			})
		}
	}
}

func TestCoords_Rotate(t *testing.T) {
	minAngle := -12
	maxAngle := 12
	cubeCoordinates := newCoords(testCoordinates[consts.Cube]).Cubes()
	expecteds := make([]*Cubes, cubeCoordinates.Size()*25)
	n := 0
	for i := cubeCoordinates.Iterator(); i.Next(); {
		centerCoord := i.Item()
		for angle := minAngle; angle <= maxAngle; angle++ {
			expecteds[n] = cubeCoordinates.Rotate(centerCoord.Cube(), angle)
			n++
		}
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%ss_Rotate", coordType.Name()), func(t2 *testing.T) {
			n := 0

			for _, centerCoord := range testCoordinates[coordType] {
				for angle := minAngle; angle <= maxAngle; angle++ {
					var actual *Cubes
					expected := expecteds[n]

					switch coordType {
					case consts.Axial:
						actual = cubeCoordinates.Axials().Rotate(centerCoord.Axial(), angle).Cubes()
					case consts.DoubleHeight:
						actual = cubeCoordinates.DoubleHeights().Rotate(centerCoord.DoubleHeight(), angle).Cubes()
					case consts.DoubleWidth:
						actual = cubeCoordinates.DoubleWidths().Rotate(centerCoord.DoubleWidth(), angle).Cubes()
					case consts.EvenQ:
						actual = cubeCoordinates.EvenQs().Rotate(centerCoord.EvenQ(), angle).Cubes()
					case consts.EvenR:
						actual = cubeCoordinates.EvenRs().Rotate(centerCoord.EvenR(), angle).Cubes()
					case consts.OddQ:
						actual = cubeCoordinates.OddQs().Rotate(centerCoord.OddQ(), angle).Cubes()
					case consts.OddR:
						actual = cubeCoordinates.OddRs().Rotate(centerCoord.OddR(), angle).Cubes()
					default:
						assert.Fail(t2, "unexpected coord type")
					}

					assert.True(t2, expected.Equal(actual))
					n++
				}
			}
		})
	}
}

func TestCoords_ReflectQ(t *testing.T) {
	cubeCoordinates := newCoords(testCoordinates[consts.Cube]).Cubes()
	expected := cubeCoordinates.ReflectQ()

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%ss_ReflectQ", coordType.Name()), func(t2 *testing.T) {
			var actual *Cubes

			switch coordType {
			case consts.Axial:
				actual = cubeCoordinates.Axials().ReflectQ().Cubes()
			case consts.DoubleHeight:
				actual = cubeCoordinates.DoubleHeights().ReflectQ().Cubes()
			case consts.DoubleWidth:
				actual = cubeCoordinates.DoubleWidths().ReflectQ().Cubes()
			case consts.EvenQ:
				actual = cubeCoordinates.EvenQs().ReflectQ().Cubes()
			case consts.EvenR:
				actual = cubeCoordinates.EvenRs().ReflectQ().Cubes()
			case consts.OddQ:
				actual = cubeCoordinates.OddQs().ReflectQ().Cubes()
			case consts.OddR:
				actual = cubeCoordinates.OddRs().ReflectQ().Cubes()
			default:
				assert.Fail(t2, "unexpected coord type")
			}

			assert.True(t2, expected.Equal(actual))
		})
	}
}

func TestCoords_ReflectR(t *testing.T) {
	cubeCoordinates := newCoords(testCoordinates[consts.Cube]).Cubes()
	expected := cubeCoordinates.ReflectR()

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%ss_ReflectR", coordType.Name()), func(t2 *testing.T) {
			var actual *Cubes

			switch coordType {
			case consts.Axial:
				actual = cubeCoordinates.Axials().ReflectR().Cubes()
			case consts.DoubleHeight:
				actual = cubeCoordinates.DoubleHeights().ReflectR().Cubes()
			case consts.DoubleWidth:
				actual = cubeCoordinates.DoubleWidths().ReflectR().Cubes()
			case consts.EvenQ:
				actual = cubeCoordinates.EvenQs().ReflectR().Cubes()
			case consts.EvenR:
				actual = cubeCoordinates.EvenRs().ReflectR().Cubes()
			case consts.OddQ:
				actual = cubeCoordinates.OddQs().ReflectR().Cubes()
			case consts.OddR:
				actual = cubeCoordinates.OddRs().ReflectR().Cubes()
			default:
				assert.Fail(t2, "unexpected coord type")
			}

			assert.True(t2, expected.Equal(actual))
		})
	}
}

func TestCoords_ReflectS(t *testing.T) {
	cubeCoordinates := newCoords(testCoordinates[consts.Cube]).Cubes()
	expected := cubeCoordinates.ReflectS()

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%ss_ReflectS", coordType.Name()), func(t2 *testing.T) {
			var actual *Cubes

			switch coordType {
			case consts.Axial:
				actual = cubeCoordinates.Axials().ReflectS().Cubes()
			case consts.DoubleHeight:
				actual = cubeCoordinates.DoubleHeights().ReflectS().Cubes()
			case consts.DoubleWidth:
				actual = cubeCoordinates.DoubleWidths().ReflectS().Cubes()
			case consts.EvenQ:
				actual = cubeCoordinates.EvenQs().ReflectS().Cubes()
			case consts.EvenR:
				actual = cubeCoordinates.EvenRs().ReflectS().Cubes()
			case consts.OddQ:
				actual = cubeCoordinates.OddQs().ReflectS().Cubes()
			case consts.OddR:
				actual = cubeCoordinates.OddRs().ReflectS().Cubes()
			default:
				assert.Fail(t2, "unexpected coord type")
			}

			assert.True(t2, expected.Equal(actual))
		})
	}
}
