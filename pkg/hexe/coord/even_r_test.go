package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvenR_Axial(t *testing.T) {
	assert.Equal(t, NewAxial(0, 0), NewEvenR(0, 0).Axial())
	assert.Equal(t, NewAxial(0, 2), NewEvenR(1, 2).Axial())
}

func TestEvenR_Cube(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewEvenR(0, 0).Cube())
	assert.Equal(t, NewCube(0, 5, -5), NewEvenR(3, 5).Cube())
}

func TestEvenR_OddR(t *testing.T) {
	assert.Equal(t, NewOddR(0, 0), NewEvenR(0, 0).OddR())
	assert.Equal(t, NewOddR(-2, 3), NewEvenR(-1, 3).OddR())
}

func TestEvenR_EvenQ(t *testing.T) {
	assert.Equal(t, NewEvenQ(0, 0), NewEvenR(0, 0).EvenQ())
	assert.Equal(t, NewEvenQ(-3, 2), NewEvenR(-1, 3).EvenQ())
}

func TestEvenR_OddQ(t *testing.T) {
	assert.Equal(t, NewOddQ(0, 0), NewEvenR(0, 0).OddQ())
	assert.Equal(t, NewOddQ(0, 2), NewEvenR(1, 2).OddQ())
}

func TestEvenR_DoubleWidth(t *testing.T) {
	assert.Equal(t, NewDoubleWidth(0, 0), NewEvenR(0, 0).DoubleWidth())
	assert.Equal(t, NewDoubleWidth(3, 1), NewEvenR(2, 1).DoubleWidth())
}

func TestEvenR_DoubleHeight(t *testing.T) {
	assert.Equal(t, NewDoubleHeight(0, 0), NewEvenR(0, 0).DoubleHeight())
	assert.Equal(t, NewDoubleHeight(0, 4), NewEvenR(1, 2).DoubleHeight())
}

func TestEvenR_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		EvenRs{
			NewEvenR(2, 1),
			NewEvenR(1, 2),
			NewEvenR(0, 2),
			NewEvenR(0, 1),
			NewEvenR(0, 0),
			NewEvenR(1, 0),
		},
		NewEvenR(1, 1).Neighbors(),
	)
}

func TestEvenR_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewEvenR(0, 0).DistanceTo(NewEvenR(0, 0)))
	assert.Equal(t, 3, NewEvenR(4, 1).DistanceTo(NewEvenR(1, 2)))
}
