package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCube_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		Cubes{
			NewCube(2, 1, -3),
			NewCube(1, 2, -3),
			NewCube(0, 2, -2),
			NewCube(0, 1, -1),
			NewCube(1, 0, -1),
			NewCube(2, 0, -2),
		},
		NewCube(1, 1, -2).Neighbors(),
	)
}

func TestCube_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewCube(0, 0, 0).DistanceTo(NewCube(0, 0, 0)))
	assert.Equal(t, 3, NewCube(1, -2, 1).DistanceTo(NewCube(2, 0, -2)))
}

func TestCube_FloodFill(t *testing.T) {
	n := 2
	blockedFunc := func(coord Cube) bool {
		return coord.Q()+coord.R() == 0
	}

	actual := NewCube(1, 0, -1).FloodFill(n, blockedFunc).Sort()
	expected := Cubes{
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
	}
	assert.Equal(t, expected, actual)
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
