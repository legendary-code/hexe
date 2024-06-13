package coord

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCoordinates = map[consts.CoordType][]Coord{
	consts.Axial: {
		NewAxial(0, 0),
		NewAxial(-1, 0),
		NewAxial(0, -1),
		NewAxial(1, -1),
		NewAxial(-1, 1),
		NewAxial(0, 1),
		NewAxial(1, 0),
	},
	consts.Cube: {
		NewCube(0, 0, 0),
		NewCube(-1, 0, 1),
		NewCube(0, -1, 1),
		NewCube(1, -1, 0),
		NewCube(-1, 1, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
	},
	consts.DoubleHeight: {
		NewDoubleHeight(0, 0),
		NewDoubleHeight(-1, -1),
		NewDoubleHeight(0, -2),
		NewDoubleHeight(1, -1),
		NewDoubleHeight(-1, 1),
		NewDoubleHeight(0, 2),
		NewDoubleHeight(1, 1),
	},
	consts.DoubleWidth: {
		NewDoubleWidth(0, 0),
		NewDoubleWidth(-2, 0),
		NewDoubleWidth(-1, -1),
		NewDoubleWidth(1, -1),
		NewDoubleWidth(-1, 1),
		NewDoubleWidth(1, 1),
		NewDoubleWidth(2, 0),
	},
	consts.EvenQ: {
		NewEvenQ(0, 0),
		NewEvenQ(-1, 0),
		NewEvenQ(0, -1),
		NewEvenQ(1, 0),
		NewEvenQ(-1, 1),
		NewEvenQ(0, 1),
		NewEvenQ(1, 1),
	},
	consts.EvenR: {
		NewEvenR(0, 0),
		NewEvenR(-1, 0),
		NewEvenR(0, -1),
		NewEvenR(1, -1),
		NewEvenR(0, 1),
		NewEvenR(1, 1),
		NewEvenR(1, 0),
	},
	consts.OddQ: {
		NewOddQ(0, 0),
		NewOddQ(-1, -1),
		NewOddQ(0, -1),
		NewOddQ(1, -1),
		NewOddQ(-1, 0),
		NewOddQ(0, 1),
		NewOddQ(1, 0),
	},
	consts.OddR: {
		NewOddR(0, 0),
		NewOddR(-1, 0),
		NewOddR(-1, -1),
		NewOddR(0, -1),
		NewOddR(-1, 1),
		NewOddR(0, 1),
		NewOddR(1, 0),
	},
}

func TestCoord_Coord(t *testing.T) {
	for _, typeA := range consts.CoordinateTypes() {
		for _, typeB := range consts.CoordinateTypes() {
			if typeA == typeB {
				continue
			}

			t.Run(fmt.Sprintf("Test%s_%s", typeA.Name(), typeB.Name()), func(t2 *testing.T) {
				typeACoords := testCoordinates[typeA]
				typeBCoords := testCoordinates[typeB]

				for i, typeACoord := range typeACoords {
					typeBCoord := typeBCoords[i]
					assert.Equal(t2, typeBCoord, typeACoord.Convert(typeB))
				}
			})
		}
	}
}

func TestCoord_Type(t *testing.T) {
	for _, coordType := range consts.CoordinateTypes() {
		t.Run(fmt.Sprintf("Test%s_Type", coordType.Name()), func(t2 *testing.T) {
			assert.Equal(t2, coordType, testCoordinates[coordType][0].Type())
		})
	}
}

func TestCoord_Neighbors(t *testing.T) {
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]*Cubes, len(cubeCoordinates))
	for i, cubeCoord := range cubeCoordinates {
		expecteds[i] = cubeCoord.Cube().Neighbors()
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_Neighbors", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				var actual *Cubes
				expected := expecteds[i]

				switch coordType {
				case consts.Axial:
					actual = coord.Axial().Neighbors().Cubes()
				case consts.DoubleHeight:
					actual = coord.DoubleHeight().Neighbors().Cubes()
				case consts.DoubleWidth:
					actual = coord.DoubleWidth().Neighbors().Cubes()
				case consts.EvenQ:
					actual = coord.EvenQ().Neighbors().Cubes()
				case consts.EvenR:
					actual = coord.EvenR().Neighbors().Cubes()
				case consts.OddQ:
					actual = coord.OddQ().Neighbors().Cubes()
				case consts.OddR:
					actual = coord.OddR().Neighbors().Cubes()
				default:
					assert.Fail(t2, "unexpected coord type")
				}

				assert.True(t2, expected.Equal(actual))
			}
		})
	}
}

func TestCoord_DiagonalNeighbors(t *testing.T) {
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]*Cubes, len(cubeCoordinates))
	for i, cubeCoord := range cubeCoordinates {
		expecteds[i] = cubeCoord.Cube().DiagonalNeighbors()
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_DiagonalNeighbors", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				var actual *Cubes
				expected := expecteds[i]

				switch coordType {
				case consts.Axial:
					actual = coord.Axial().DiagonalNeighbors().Cubes()
				case consts.DoubleHeight:
					actual = coord.DoubleHeight().DiagonalNeighbors().Cubes()
				case consts.DoubleWidth:
					actual = coord.DoubleWidth().DiagonalNeighbors().Cubes()
				case consts.EvenQ:
					actual = coord.EvenQ().DiagonalNeighbors().Cubes()
				case consts.EvenR:
					actual = coord.EvenR().DiagonalNeighbors().Cubes()
				case consts.OddQ:
					actual = coord.OddQ().DiagonalNeighbors().Cubes()
				case consts.OddR:
					actual = coord.OddR().DiagonalNeighbors().Cubes()
				default:
					assert.Fail(t2, "unexpected coord type")
				}

				assert.True(t2, expected.Equal(actual))
			}
		})
	}
}

func TestCoord_DistanceTo(t *testing.T) {
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]int, len(cubeCoordinates)*len(cubeCoordinates))
	for i, cubeCoord := range cubeCoordinates {
		for j, otherCoord := range cubeCoordinates {
			expecteds[i*len(cubeCoordinates)+j] = cubeCoord.Cube().DistanceTo(otherCoord.Cube())
		}
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_DistanceTo", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				for j, otherCoord := range testCoordinates[coordType] {
					var actual int
					expected := expecteds[i*len(cubeCoordinates)+j]

					switch coordType {
					case consts.Axial:
						actual = coord.Axial().DistanceTo(otherCoord.Axial())
					case consts.DoubleHeight:
						actual = coord.DoubleHeight().DistanceTo(otherCoord.DoubleHeight())
					case consts.DoubleWidth:
						actual = coord.DoubleWidth().DistanceTo(otherCoord.DoubleWidth())
					case consts.EvenQ:
						actual = coord.EvenQ().DistanceTo(otherCoord.EvenQ())
					case consts.EvenR:
						actual = coord.EvenR().DistanceTo(otherCoord.EvenR())
					case consts.OddQ:
						actual = coord.OddQ().DistanceTo(otherCoord.OddQ())
					case consts.OddR:
						actual = coord.OddR().DistanceTo(otherCoord.OddR())
					default:
						assert.Fail(t2, "unexpected coord type")
					}

					assert.Equal(t2, expected, actual)
				}
			}
		})
	}
}

func TestCoord_LineTo(t *testing.T) {
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]*Cubes, len(cubeCoordinates)*len(cubeCoordinates))
	for i, cubeCoord := range cubeCoordinates {
		for j, otherCoord := range cubeCoordinates {
			expecteds[i*len(cubeCoordinates)+j] = cubeCoord.Cube().LineTo(otherCoord.Cube())
		}
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_LineTo", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				for j, otherCoord := range testCoordinates[coordType] {
					var actual *Cubes
					expected := expecteds[i*len(cubeCoordinates)+j]

					switch coordType {
					case consts.Axial:
						actual = coord.Axial().LineTo(otherCoord.Axial()).Cubes()
					case consts.DoubleHeight:
						actual = coord.DoubleHeight().LineTo(otherCoord.DoubleHeight()).Cubes()
					case consts.DoubleWidth:
						actual = coord.DoubleWidth().LineTo(otherCoord.DoubleWidth()).Cubes()
					case consts.EvenQ:
						actual = coord.EvenQ().LineTo(otherCoord.EvenQ()).Cubes()
					case consts.EvenR:
						actual = coord.EvenR().LineTo(otherCoord.EvenR()).Cubes()
					case consts.OddQ:
						actual = coord.OddQ().LineTo(otherCoord.OddQ()).Cubes()
					case consts.OddR:
						actual = coord.OddR().LineTo(otherCoord.OddR()).Cubes()
					default:
						assert.Fail(t2, "unexpected coord type")
					}

					assert.True(t2, expected.Equal(actual))
				}
			}
		})
	}
}

func TestCoord_MovementRange(t *testing.T) {
	n := 3
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]*Cubes, len(cubeCoordinates)*n)
	for i, cubeCoord := range cubeCoordinates {
		for j := 0; j < n; j++ {
			expecteds[i*n+j] = cubeCoord.Cube().MovementRange(j)
		}
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_MovementRange", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				for j := 0; j < n; j++ {
					var actual *Cubes
					expected := expecteds[i*n+j]

					switch coordType {
					case consts.Axial:
						actual = coord.Axial().MovementRange(j).Cubes()
					case consts.DoubleHeight:
						actual = coord.DoubleHeight().MovementRange(j).Cubes()
					case consts.DoubleWidth:
						actual = coord.DoubleWidth().MovementRange(j).Cubes()
					case consts.EvenQ:
						actual = coord.EvenQ().MovementRange(j).Cubes()
					case consts.EvenR:
						actual = coord.EvenR().MovementRange(j).Cubes()
					case consts.OddQ:
						actual = coord.OddQ().MovementRange(j).Cubes()
					case consts.OddR:
						actual = coord.OddR().MovementRange(j).Cubes()
					default:
						assert.Fail(t2, "unexpected coord type")
					}

					assert.True(t2, expected.Equal(actual))
				}
			}
		})
	}
}

func TestCoord_FloodFill(t *testing.T) {
	n := 3
	isBlocked := func(coord Cube) bool {
		return coord.Axial().Q() > coord.Axial().R()
	}

	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]*Cubes, len(cubeCoordinates)*n)
	for i, cubeCoord := range cubeCoordinates {
		for j := 0; j < n; j++ {
			expecteds[i*n+j] = cubeCoord.Cube().FloodFill(j, isBlocked)
		}
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_FloodFill", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				for j := 0; j < n; j++ {
					var actual *Cubes
					expected := expecteds[i*n+j]

					switch coordType {
					case consts.Axial:
						actual = coord.Axial().FloodFill(j, func(coord Axial) bool {
							return coord.Q() > coord.R()
						}).Cubes()
					case consts.DoubleHeight:
						actual = coord.DoubleHeight().FloodFill(j, func(coord DoubleHeight) bool {
							return coord.Axial().Q() > coord.Axial().R()
						}).Cubes()
					case consts.DoubleWidth:
						actual = coord.DoubleWidth().FloodFill(j, func(coord DoubleWidth) bool {
							return coord.Axial().Q() > coord.Axial().R()
						}).Cubes()
					case consts.EvenQ:
						actual = coord.EvenQ().FloodFill(j, func(coord EvenQ) bool {
							return coord.Axial().Q() > coord.Axial().R()
						}).Cubes()
					case consts.EvenR:
						actual = coord.EvenR().FloodFill(j, func(coord EvenR) bool {
							return coord.Axial().Q() > coord.Axial().R()
						}).Cubes()
					case consts.OddQ:
						actual = coord.OddQ().FloodFill(j, func(coord OddQ) bool {
							return coord.Axial().Q() > coord.Axial().R()
						}).Cubes()
					case consts.OddR:
						actual = coord.OddR().FloodFill(j, func(coord OddR) bool {
							return coord.Axial().Q() > coord.Axial().R()
						}).Cubes()
					default:
						assert.Fail(t2, "unexpected coord type")
					}

					assert.True(t2, expected.SetEqual(actual))
				}
			}
		})
	}
}

func TestCoord_Rotate(t *testing.T) {
	minAngle := -12
	maxAngle := 12
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]Cube, len(cubeCoordinates)*len(cubeCoordinates)*25)
	n := 0
	for _, cubeCoord := range cubeCoordinates {
		for _, otherCoord := range cubeCoordinates {
			for angle := minAngle; angle <= maxAngle; angle++ {
				expecteds[n] = cubeCoord.Cube().Rotate(otherCoord.Cube(), angle)
				n++
			}
		}
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_Rotate", coordType.Name()), func(t2 *testing.T) {
			n := 0

			for _, coord := range testCoordinates[coordType] {
				for _, otherCoord := range testCoordinates[coordType] {
					for angle := minAngle; angle <= maxAngle; angle++ {
						var actual Cube
						expected := expecteds[n]

						switch coordType {
						case consts.Axial:
							actual = coord.Axial().Rotate(otherCoord.Axial(), angle).Cube()
						case consts.DoubleHeight:
							actual = coord.DoubleHeight().Rotate(otherCoord.DoubleHeight(), angle).Cube()
						case consts.DoubleWidth:
							actual = coord.DoubleWidth().Rotate(otherCoord.DoubleWidth(), angle).Cube()
						case consts.EvenQ:
							actual = coord.EvenQ().Rotate(otherCoord.EvenQ(), angle).Cube()
						case consts.EvenR:
							actual = coord.EvenR().Rotate(otherCoord.EvenR(), angle).Cube()
						case consts.OddQ:
							actual = coord.OddQ().Rotate(otherCoord.OddQ(), angle).Cube()
						case consts.OddR:
							actual = coord.OddR().Rotate(otherCoord.OddR(), angle).Cube()
						default:
							assert.Fail(t2, "unexpected coord type")
						}

						assert.Equal(t2, expected, actual)
						n++
					}
				}
			}
		})
	}
}

func TestCoord_ReflectQ(t *testing.T) {
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]Cube, len(cubeCoordinates))
	for i, cubeCoord := range cubeCoordinates {
		expecteds[i] = cubeCoord.Cube().ReflectQ()
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_ReflectQ", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				var actual Cube
				expected := expecteds[i]

				switch coordType {
				case consts.Axial:
					actual = coord.Axial().ReflectQ().Cube()
				case consts.DoubleHeight:
					actual = coord.DoubleHeight().ReflectQ().Cube()
				case consts.DoubleWidth:
					actual = coord.DoubleWidth().ReflectQ().Cube()
				case consts.EvenQ:
					actual = coord.EvenQ().ReflectQ().Cube()
				case consts.EvenR:
					actual = coord.EvenR().ReflectQ().Cube()
				case consts.OddQ:
					actual = coord.OddQ().ReflectQ().Cube()
				case consts.OddR:
					actual = coord.OddR().ReflectQ().Cube()
				default:
					assert.Fail(t2, "unexpected coord type")
				}

				assert.Equal(t2, expected, actual)
			}
		})
	}
}

func TestCoord_ReflectR(t *testing.T) {
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]Cube, len(cubeCoordinates))
	for i, cubeCoord := range cubeCoordinates {
		expecteds[i] = cubeCoord.Cube().ReflectR()
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_ReflectR", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				var actual Cube
				expected := expecteds[i]

				switch coordType {
				case consts.Axial:
					actual = coord.Axial().ReflectR().Cube()
				case consts.DoubleHeight:
					actual = coord.DoubleHeight().ReflectR().Cube()
				case consts.DoubleWidth:
					actual = coord.DoubleWidth().ReflectR().Cube()
				case consts.EvenQ:
					actual = coord.EvenQ().ReflectR().Cube()
				case consts.EvenR:
					actual = coord.EvenR().ReflectR().Cube()
				case consts.OddQ:
					actual = coord.OddQ().ReflectR().Cube()
				case consts.OddR:
					actual = coord.OddR().ReflectR().Cube()
				default:
					assert.Fail(t2, "unexpected coord type")
				}

				assert.Equal(t2, expected, actual)
			}
		})
	}
}

func TestCoord_ReflectS(t *testing.T) {
	cubeCoordinates := testCoordinates[consts.Cube]
	expecteds := make([]Cube, len(cubeCoordinates))
	for i, cubeCoord := range cubeCoordinates {
		expecteds[i] = cubeCoord.Cube().ReflectS()
	}

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			continue
		}

		t.Run(fmt.Sprintf("Test%s_ReflectS", coordType.Name()), func(t2 *testing.T) {
			for i, coord := range testCoordinates[coordType] {
				var actual Cube
				expected := expecteds[i]

				switch coordType {
				case consts.Axial:
					actual = coord.Axial().ReflectS().Cube()
				case consts.DoubleHeight:
					actual = coord.DoubleHeight().ReflectS().Cube()
				case consts.DoubleWidth:
					actual = coord.DoubleWidth().ReflectS().Cube()
				case consts.EvenQ:
					actual = coord.EvenQ().ReflectS().Cube()
				case consts.EvenR:
					actual = coord.EvenR().ReflectS().Cube()
				case consts.OddQ:
					actual = coord.OddQ().ReflectS().Cube()
				case consts.OddR:
					actual = coord.OddR().ReflectS().Cube()
				default:
					assert.Fail(t2, "unexpected coord type")
				}

				assert.Equal(t2, expected, actual)
			}
		})
	}
}
