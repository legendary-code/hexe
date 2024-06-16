package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCubes_UnionWith(t *testing.T) {
	a := Cubes{
		NewCube(0, 0, 0),
		NewCube(1, 0, -1),
	}

	b := Cubes{
		NewCube(0, 1, -1),
		NewCube(1, 1, -2),
	}

	c := a.UnionWith(b).Sort()
	expected := Cubes{
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(1, 1, -2),
	}

	assert.Equal(t, expected, c)
}

func TestCubes_IntersectWith(t *testing.T) {
	a := Cubes{
		NewCube(0, 0, 0),
		NewCube(1, 0, -1),
		NewCube(0, 1, -1),
		NewCube(1, 1, -2),
	}

	b := Cubes{
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
	}

	c := a.IntersectWith(b).Sort()
	expected := Cubes{
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
	}

	assert.Equal(t, expected, c)
}

func TestCubes_DifferenceWith(t *testing.T) {
	a := Cubes{
		NewCube(0, 0, 0),
		NewCube(1, 0, -1),
		NewCube(0, 1, -1),
		NewCube(1, 1, -2),
	}

	b := Cubes{
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
	}

	c := a.DifferenceWith(b).Sort()
	expected := Cubes{
		NewCube(0, 0, 0),
		NewCube(1, 1, -2),
	}

	assert.Equal(t, expected, c)
}

func TestCubes_Copy(t *testing.T) {
	expected := Cubes{
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(1, 1, -2),
	}

	actual := expected.Copy()
	actual[0] = NewCube(2, 2, -4)

	assert.NotEqual(t, expected, actual)
}

func TestCubes_Sort(t *testing.T) {
	expected := Cubes{
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(1, 1, -2),
	}

	actual := Cubes{
		NewCube(1, 1, -2),
		NewCube(0, 0, 0),
		NewCube(1, 0, -1),
		NewCube(0, 1, -1),
	}

	actual = actual.Sort()
	assert.Equal(t, expected, actual)
}

func TestCubes_Rotate(t *testing.T) {
	center := NewCube(1, 2, -3)
	angle := -2

	coords := Cubes{
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
	}

	expected := Cubes{
		coords[0].Rotate(center, angle),
		coords[1].Rotate(center, angle),
		coords[2].Rotate(center, angle),
		coords[3].Rotate(center, angle),
		coords[4].Rotate(center, angle),
	}

	actual := coords.Rotate(center, angle)
	assert.Equal(t, expected.Sort(), actual.Sort())
}
