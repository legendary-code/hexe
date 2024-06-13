package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCubes_Rotate(t *testing.T) {
	center := NewCube(1, 2, -3)
	angle := -2

	coords := NewCubes(
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
	)

	expected := NewCubes()
	for i := coords.Iterator(); i.Next(); {
		expected.Add(i.Item().Rotate(center, angle))
	}

	actual := coords.Rotate(center, angle)
	assert.True(t, expected.Equal(actual))
}

func TestCubes_ReflectQ(t *testing.T) {
	coords := NewCubes(
		NewCube(0, 0, 0),
		NewCube(1, -1, 0),
		NewCube(-1, 2, -1),
	)

	expected := NewCubes(
		NewCube(0, 0, 0),
		NewCube(1, 0, -1),
		NewCube(-1, -1, 2),
	)

	assert.True(t, expected.Equal(coords.ReflectQ()))
}

func TestCubes_ReflectR(t *testing.T) {
	coords := NewCubes(
		NewCube(0, 0, 0),
		NewCube(1, -1, 0),
		NewCube(-1, 2, -1),
	)

	expected := NewCubes(
		NewCube(0, 0, 0),
		NewCube(0, -1, 1),
		NewCube(-1, 2, -1),
	)

	assert.True(t, expected.Equal(coords.ReflectR()))
}

func TestCubes_ReflectS(t *testing.T) {
	coords := NewCubes(
		NewCube(0, 0, 0),
		NewCube(1, -1, 0),
		NewCube(-1, 2, -1),
	)

	expected := NewCubes(
		NewCube(0, 0, 0),
		NewCube(-1, 1, 0),
		NewCube(2, -1, -1),
	)

	assert.True(t, expected.Equal(coords.ReflectS()))
}
