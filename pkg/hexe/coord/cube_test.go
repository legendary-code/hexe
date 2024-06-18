package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCube_Add(t *testing.T) {
	assert.Equal(t, NewCube(2, 4, -6), NewCube(1, 2, -3).Add(NewCube(1, 2, -3)))
}

func TestCube_Scale(t *testing.T) {
	assert.Equal(t, NewCube(-3, 6, -3), NewCube(-1, 2, -1).Scale(3))
}

func TestCube_Neighbor(t *testing.T) {
	assert.Equal(t, NewCube(-2, 2, 0), NewCube(-1, 2, -1).Neighbor(3))
}

func TestCube_Neighbors(t *testing.T) {
	expected := NewCubes(
		NewCube(2, 1, -3),
		NewCube(1, 2, -3),
		NewCube(0, 2, -2),
		NewCube(0, 1, -1),
		NewCube(1, 0, -1),
		NewCube(2, 0, -2),
	)
	actual := NewCube(1, 1, -2).Neighbors()
	assert.True(t, expected.Equal(actual))
}

func TestCube_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewCube(0, 0, 0).DistanceTo(NewCube(0, 0, 0)))
	assert.Equal(t, 3, NewCube(1, -2, 1).DistanceTo(NewCube(2, 0, -2)))
}

func TestCube_LineTo(t *testing.T) {
	expected := NewCubes(
		NewCube(-1, -1, 2),
		NewCube(-1, 0, 1),
		NewCube(0, 0, 0),
		NewCube(0, 1, -1),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
	)
	actual := NewCube(-1, -1, 2).LineTo(NewCube(1, 2, -3))
	assert.True(t, expected.Equal(actual))
}

func TestCube_TraceTo(t *testing.T) {
	expected := NewCubes(
		NewCube(-1, -1, 2),
		NewCube(-1, 0, 1),
		NewCube(0, 0, 0),
	)
	actual := NewCube(-1, -1, 2).TraceTo(NewCube(1, 2, -3), func(coord Cube) bool {
		return coord.Q() > 0 || coord.R() > 0
	})
	assert.True(t, expected.Equal(actual))
}

func TestCube_FloodFill(t *testing.T) {
	n := 2
	blockedFunc := func(coord Cube) bool {
		return coord.Q()+coord.R() == 0
	}

	actual := NewCube(1, 0, -1).FloodFill(n, blockedFunc)
	expected := NewCubes(
		NewCube(-1, 2, -1),
		NewCube(0, 1, -1),
		NewCube(0, 2, -2),
		NewCube(1, 0, -1),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
		NewCube(2, -1, -1),
		NewCube(2, 0, -2),
		NewCube(2, 1, -3),
		NewCube(3, -2, -1),
		NewCube(3, -1, -2),
		NewCube(3, 0, -3),
	)
	assert.True(t, expected.SetEqual(actual))
}

func TestCube_ReflectQ(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewCube(0, 0, 0).ReflectQ())
	assert.Equal(t, NewCube(1, 1, -2), NewCube(1, -2, 1).ReflectQ())
}

func TestCube_ReflectR(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewCube(0, 0, 0).ReflectR())
	assert.Equal(t, NewCube(1, -2, 1), NewCube(1, -2, 1).ReflectR())
}

func TestCube_ReflectS(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewCube(0, 0, 0).ReflectS())
	assert.Equal(t, NewCube(-2, 1, 1), NewCube(1, -2, 1).ReflectS())
}

func TestCube_Ring(t *testing.T) {
	expected := NewCubes(
		NewCube(1, -1, 0),
		NewCube(2, -1, -1),
		NewCube(3, -1, -2),
		NewCube(4, -1, -3),
		NewCube(4, 0, -4),
		NewCube(4, 1, -5),
		NewCube(4, 2, -6),
		NewCube(3, 3, -6),
		NewCube(2, 4, -6),
		NewCube(1, 5, -6),
		NewCube(0, 5, -5),
		NewCube(-1, 5, -4),
		NewCube(-2, 5, -3),
		NewCube(-2, 4, -2),
		NewCube(-2, 3, -1),
		NewCube(-2, 2, 0),
		NewCube(-1, 1, 0),
		NewCube(0, 0, 0),
	)
	actual := NewCube(1, 2, -3).Ring(3)
	assert.True(t, expected.Equal(actual))
}

func TestCube_FieldOfView(t *testing.T) {
	expected := NewCubes(
		NewCube(0, 1, -1),
		NewCube(0, 2, -2),
		NewCube(0, 3, -3),
		NewCube(0, 4, -4),
		NewCube(1, 0, -1),
		NewCube(1, 1, -2),
		NewCube(1, 2, -3),
		NewCube(1, 3, -4),
		NewCube(1, 4, -5),
		NewCube(2, 0, -2),
		NewCube(2, 1, -3),
		NewCube(2, 2, -4),
		NewCube(2, 3, -5),
		NewCube(3, 0, -3),
		NewCube(3, 1, -4),
		NewCube(3, 2, -5),
	)
	actual := NewCube(1, 2, -3).FieldOfView(2, func(coord Cube) bool {
		return coord.Q() < 0 || coord.R() < 0
	})
	assert.True(t, expected.SetEqual(actual))
}

func TestCube_FindPathBFS(t *testing.T) {
	expected := NewCubes(
		NewCube(1, 2, -3),
		NewCube(2, 2, -4),
		NewCube(3, 2, -5),
		NewCube(4, 2, -6),
		NewCube(5, 2, -7),
		NewCube(6, 2, -8),
		NewCube(6, 3, -9),
	)
	actual := NewCube(1, 2, -3).FindPathBFS(NewCube(6, 3, -9), 10, func(coord Cube) bool {
		return coord.Q() < 0 || coord.R() < 0
	})
	assert.True(t, expected.Equal(actual))
}
