package coord

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleWidth_Axial(t *testing.T) {
	assert.Equal(t, NewAxial(0, 0), NewDoubleWidth(0, 0).Axial())
	assert.Equal(t, NewAxial(-1, 4), NewDoubleWidth(2, 4).Axial())
}

func TestDoubleWidth_Cube(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewDoubleWidth(0, 0).Cube())
	assert.Equal(t, NewCube(-1, 4, -3), NewDoubleWidth(2, 4).Cube())
}

func TestDoubleWidth_OddR(t *testing.T) {
	assert.Equal(t, NewOddR(0, 0), NewDoubleWidth(0, 0).OddR())
	assert.Equal(t, NewOddR(1, 4), NewDoubleWidth(2, 4).OddR())
}

func TestDoubleWidth_EvenR(t *testing.T) {
	assert.Equal(t, NewEvenR(0, 0), NewDoubleWidth(0, 0).EvenR())
	assert.Equal(t, NewEvenR(1, 4), NewDoubleWidth(2, 4).EvenR())
}

func TestDoubleWidth_OddQ(t *testing.T) {
	assert.Equal(t, NewOddQ(0, 0), NewDoubleWidth(0, 0).OddQ())
	assert.Equal(t, NewOddQ(-1, 3), NewDoubleWidth(2, 4).OddQ())
}

func TestDoubleWidth_EvenQ(t *testing.T) {
	assert.Equal(t, NewEvenQ(0, 0), NewDoubleWidth(0, 0).EvenQ())
	assert.Equal(t, NewEvenQ(-1, 4), NewDoubleWidth(2, 4).EvenQ())
}

func TestDoubleWidth_DoubleHeight(t *testing.T) {
	assert.Equal(t, NewDoubleHeight(0, 0), NewDoubleWidth(0, 0).DoubleHeight())
	assert.Equal(t, NewDoubleHeight(-1, 7), NewDoubleWidth(2, 4).DoubleHeight())
}

func TestDoubleWidth_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		[consts.Sides]DoubleWidth{
			NewDoubleWidth(3, 1),
			NewDoubleWidth(2, 2),
			NewDoubleWidth(0, 2),
			NewDoubleWidth(-1, 1),
			NewDoubleWidth(0, 0),
			NewDoubleWidth(2, 0),
		},
		NewDoubleWidth(1, 1).Neighbors(),
	)
}

func TestDoubleWidth_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewDoubleWidth(0, 0).DistanceTo(NewDoubleWidth(0, 0)))
	assert.Equal(t, 3, NewDoubleWidth(6, 2).DistanceTo(NewDoubleWidth(11, 3)))
}
