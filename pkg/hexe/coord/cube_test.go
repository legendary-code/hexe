package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCube_Axial(t *testing.T) {
	assert.Equal(t, NewAxial(0, 0), NewCube(0, 0, 0).Axial())
	assert.Equal(t, NewAxial(3, 5), NewCube(3, 5, -8).Axial())
}

func TestCube_OddR(t *testing.T) {
	assert.Equal(t, NewOddR(0, 0), NewCube(0, 0, 0).OddR())
	assert.Equal(t, NewOddR(0, 3), NewCube(-1, 3, -2).OddR())
}

func TestCube_EvenR(t *testing.T) {
	assert.Equal(t, NewEvenR(0, 0), NewCube(0, 0, 0).EvenR())
	assert.Equal(t, NewEvenR(1, 3), NewCube(-1, 3, -2).EvenR())
}

func TestCube_OddQ(t *testing.T) {
	assert.Equal(t, NewOddQ(0, 0), NewCube(0, 0, 0).OddQ())
	assert.Equal(t, NewOddQ(1, 2), NewCube(1, 2, -3).OddQ())
}

func TestCube_EvenQ(t *testing.T) {
	assert.Equal(t, NewEvenQ(0, 0), NewCube(0, 0, 0).EvenQ())
	assert.Equal(t, NewEvenQ(1, 3), NewCube(1, 2, -3).EvenQ())
}

func TestCube_DoubleWidth(t *testing.T) {
	assert.Equal(t, NewDoubleWidth(0, 0), NewCube(0, 0, 0).DoubleWidth())
	assert.Equal(t, NewDoubleWidth(5, 1), NewCube(2, 1, -3).DoubleWidth())
}

func TestCube_DoubleHeight(t *testing.T) {
	assert.Equal(t, NewDoubleHeight(0, 0), NewCube(0, 0, 0).DoubleHeight())
	assert.Equal(t, NewDoubleHeight(1, 5), NewCube(1, 2, -3).DoubleHeight())
}

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
