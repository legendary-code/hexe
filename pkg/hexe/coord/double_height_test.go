package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleHeight_Axial(t *testing.T) {
	assert.Equal(t, NewAxial(0, 0), NewDoubleHeight(0, 0).Axial())
	assert.Equal(t, NewAxial(2, 1), NewDoubleHeight(2, 4).Axial())
}

func TestDoubleHeight_Cube(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewDoubleHeight(0, 0).Cube())
	assert.Equal(t, NewCube(2, 1, -3), NewDoubleHeight(2, 4).Cube())
}

func TestDoubleHeight_OddR(t *testing.T) {
	assert.Equal(t, NewOddR(0, 0), NewDoubleHeight(0, 0).OddR())
	assert.Equal(t, NewOddR(2, 1), NewDoubleHeight(2, 4).OddR())
}

func TestDoubleHeight_EvenR(t *testing.T) {
	assert.Equal(t, NewEvenR(0, 0), NewDoubleHeight(0, 0).EvenR())
	assert.Equal(t, NewEvenR(3, 1), NewDoubleHeight(2, 4).EvenR())
}

func TestDoubleHeight_OddQ(t *testing.T) {
	assert.Equal(t, NewOddQ(0, 0), NewDoubleHeight(0, 0).OddQ())
	assert.Equal(t, NewOddQ(2, 2), NewDoubleHeight(2, 4).OddQ())
}

func TestDoubleHeight_EvenQ(t *testing.T) {
	assert.Equal(t, NewEvenQ(0, 0), NewDoubleHeight(0, 0).EvenQ())
	assert.Equal(t, NewEvenQ(2, 2), NewDoubleHeight(2, 4).EvenQ())
}

func TestDoubleHeight_DoubleWidth(t *testing.T) {
	assert.Equal(t, NewDoubleWidth(0, 0), NewDoubleHeight(0, 0).DoubleWidth())
	assert.Equal(t, NewDoubleWidth(5, 1), NewDoubleHeight(2, 4).DoubleWidth())
}

func TestDoubleHeight_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		DoubleHeights{
			NewDoubleHeight(2, 2),
			NewDoubleHeight(1, 3),
			NewDoubleHeight(0, 2),
			NewDoubleHeight(0, 0),
			NewDoubleHeight(1, -1),
			NewDoubleHeight(2, 0),
		},
		NewDoubleHeight(1, 1).Neighbors(),
	)
}

func TestDoubleHeight_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewDoubleHeight(0, 0).DistanceTo(NewDoubleHeight(0, 0)))
	assert.Equal(t, 3, NewDoubleHeight(2, 2).DistanceTo(NewDoubleHeight(4, 6)))
}
