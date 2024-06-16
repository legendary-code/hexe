package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoords_Type(t *testing.T) {
	t.Parallel()

	for coordType, coords := range testCoordinates {
		t.Run(fmt.Sprintf("Test%ss_Type", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()
			assert.Equal(t2, coordType, NewCoords(coords).Type())
		})
	}
}

func TestCoords_Coords(t *testing.T) {
	t.Parallel()

	for _, typeA := range consts.CoordinateTypes() {
		for _, typeB := range consts.CoordinateTypes() {
			if typeA == typeB {
				continue
			}

			t.Run(fmt.Sprintf("Test%ss_%ss", typeA.Name(), typeB.Name()), func(t2 *testing.T) {
				t2.Parallel()
				typeACoords := NewCoords(testCoordinates[typeA])
				typeBCoords := NewCoords(testCoordinates[typeB])

				assert.Equal(t2, typeBCoords, typeACoords.Convert(typeB))
			})
		}
	}
}

func TestCoords_ToSlice(t *testing.T) {
	t.Parallel()

	for _, coordType := range consts.CoordinateTypes() {
		t.Run(fmt.Sprintf("Test%ss_ToSlice", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()
			coords := NewCoords(testCoordinates[coordType])
			assert.Equal(t2, coords, NewCoords(coords.ToSlice()))
		})
	}
}

func TestCoords_Copy(t *testing.T) {
	t.Parallel()

	for _, coordType := range consts.CoordinateTypes() {
		t.Run(fmt.Sprintf("Test%ss_ToSlice", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()
			coords := NewCoords(testCoordinates[coordType])
			var cpy Coords

			switch coordType {
			case consts.Axial:
				cpy = coords.Axials().Copy()
			case consts.Cube:
				cpy = coords.Cubes().Copy()
			case consts.DoubleHeight:
				cpy = coords.DoubleHeights().Copy()
			case consts.DoubleWidth:
				cpy = coords.DoubleWidths().Copy()
			case consts.EvenQ:
				cpy = coords.EvenQs().Copy()
			case consts.EvenR:
				cpy = coords.EvenRs().Copy()
			case consts.OddQ:
				cpy = coords.OddQs().Copy()
			case consts.OddR:
				cpy = coords.OddRs().Copy()
			default:
				panic("unknown coord type")
			}

			assert.NotEqual(t2, fmt.Sprintf("%p", &coords), fmt.Sprintf("%p", &cpy))
		})
	}
}

func TestCoords_UnionWith(t *testing.T) {
	t.Parallel()

	a := Cubes{
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
	}

	b := Cubes{
		NewCube(0, 0, 0),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
	}

	expected := Cubes{
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
	}

	for _, coordType := range consts.CoordinateTypes() {
		t.Run(fmt.Sprintf("Test%ss_UnionWith", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()
			var actual Coords

			switch coordType {
			case consts.Axial:
				actual = a.Axials().UnionWith(b.Axials())
			case consts.Cube:
				actual = a.Cubes().UnionWith(b.Cubes())
			case consts.DoubleHeight:
				actual = a.DoubleHeights().UnionWith(b.DoubleHeights())
			case consts.DoubleWidth:
				actual = a.DoubleWidths().UnionWith(b.DoubleWidths())
			case consts.EvenQ:
				actual = a.EvenQs().UnionWith(b.EvenQs())
			case consts.EvenR:
				actual = a.EvenRs().UnionWith(b.EvenRs())
			case consts.OddQ:
				actual = a.OddQs().UnionWith(b.OddQs())
			case consts.OddR:
				actual = a.OddRs().UnionWith(b.OddRs())
			default:
				panic("unknown coord type")
			}

			assert.Equal(t2, expected, actual.Cubes().Sort())
		})
	}
}

func TestCoords_IntersectWith(t *testing.T) {
	t.Parallel()

	a := Cubes{
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(1, 2, -3),
	}

	b := Cubes{
		NewCube(0, 0, 0),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
	}

	expected := Cubes{
		NewCube(0, 0, 0),
		NewCube(1, 2, -3),
	}

	for _, coordType := range consts.CoordinateTypes() {
		t.Run(fmt.Sprintf("Test%ss_IntersectWith", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()
			var actual Coords

			switch coordType {
			case consts.Axial:
				actual = a.Axials().IntersectWith(b.Axials())
			case consts.Cube:
				actual = a.Cubes().IntersectWith(b.Cubes())
			case consts.DoubleHeight:
				actual = a.DoubleHeights().IntersectWith(b.DoubleHeights())
			case consts.DoubleWidth:
				actual = a.DoubleWidths().IntersectWith(b.DoubleWidths())
			case consts.EvenQ:
				actual = a.EvenQs().IntersectWith(b.EvenQs())
			case consts.EvenR:
				actual = a.EvenRs().IntersectWith(b.EvenRs())
			case consts.OddQ:
				actual = a.OddQs().IntersectWith(b.OddQs())
			case consts.OddR:
				actual = a.OddRs().IntersectWith(b.OddRs())
			default:
				panic("unknown coord type")
			}

			assert.Equal(t2, expected, actual.Cubes().Sort())
		})
	}
}

func TestCoords_DifferenceWith(t *testing.T) {
	t.Parallel()

	a := Cubes{
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(1, 2, -3),
	}

	b := Cubes{
		NewCube(0, 0, 0),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
	}

	expected := Cubes{
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
	}

	for _, coordType := range consts.CoordinateTypes() {
		t.Run(fmt.Sprintf("Test%ss_DifferenceWith", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()
			var actual Coords

			switch coordType {
			case consts.Axial:
				actual = a.Axials().DifferenceWith(b.Axials())
			case consts.Cube:
				actual = a.Cubes().DifferenceWith(b.Cubes())
			case consts.DoubleHeight:
				actual = a.DoubleHeights().DifferenceWith(b.DoubleHeights())
			case consts.DoubleWidth:
				actual = a.DoubleWidths().DifferenceWith(b.DoubleWidths())
			case consts.EvenQ:
				actual = a.EvenQs().DifferenceWith(b.EvenQs())
			case consts.EvenR:
				actual = a.EvenRs().DifferenceWith(b.EvenRs())
			case consts.OddQ:
				actual = a.OddQs().DifferenceWith(b.OddQs())
			case consts.OddR:
				actual = a.OddRs().DifferenceWith(b.OddRs())
			default:
				panic("unknown coord type")
			}

			assert.Equal(t2, expected, actual.Cubes().Sort())
		})
	}
}

func TestCoords_Rotate(t *testing.T) {
	t.Parallel()

	minAngle := -12
	maxAngle := 12
	cubeCoordinates := NewCoords(testCoordinates[consts.Cube]).Cubes()
	expecteds := make([]Cubes, len(cubeCoordinates)*25)
	n := 0
	for _, centerCoord := range cubeCoordinates {
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
			t2.Parallel()
			n := 0

			for _, centerCoord := range testCoordinates[coordType] {
				for angle := minAngle; angle <= maxAngle; angle++ {
					var actual Cubes
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

					assert.Equal(t2, expected, actual)
					n++

				}
			}
		})
	}
}

func TestCoords_ReflectQ(t *testing.T) {
	t.Parallel()

	cubeCoordinates := NewCoords(testCoordinates[consts.Cube]).Cubes()
	expected := cubeCoordinates.ReflectQ()

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%ss_ReflectQ", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()

			var actual Cubes

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

			assert.Equal(t2, expected, actual)
		})
	}
}

func TestCoords_ReflectR(t *testing.T) {
	t.Parallel()

	cubeCoordinates := NewCoords(testCoordinates[consts.Cube]).Cubes()
	expected := cubeCoordinates.ReflectR()

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%ss_ReflectR", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()

			var actual Cubes

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

			assert.Equal(t2, expected, actual)
		})
	}
}

func TestCoords_ReflectS(t *testing.T) {
	t.Parallel()

	cubeCoordinates := NewCoords(testCoordinates[consts.Cube]).Cubes()
	expected := cubeCoordinates.ReflectS()

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%ss_ReflectS", coordType.Name()), func(t2 *testing.T) {
			t2.Parallel()

			var actual Cubes

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

			assert.Equal(t2, expected, actual)
		})
	}
}
