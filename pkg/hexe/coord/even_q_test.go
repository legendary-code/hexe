package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvenQ_Axial(t *testing.T) {
	assert.Equal(t, NewAxial(0, 0), NewEvenQ(0, 0).Axial())
	assert.Equal(t, NewAxial(1, 1), NewEvenQ(1, 2).Axial())
}

func TestEvenQ_Cube(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewEvenQ(0, 0).Cube())
	assert.Equal(t, NewCube(3, 3, -6), NewEvenQ(3, 5).Cube())
}

func TestEvenQ_OddR(t *testing.T) {
	assert.Equal(t, NewOddR(0, 0), NewEvenQ(0, 0).OddR())
	assert.Equal(t, NewOddR(0, 3), NewEvenQ(-1, 3).OddR())
}

func TestEvenQ_EvenR(t *testing.T) {
	assert.Equal(t, NewEvenR(0, 0), NewEvenQ(0, 0).EvenR())
	assert.Equal(t, NewEvenR(1, 3), NewEvenQ(-1, 3).EvenR())
}

func TestEvenQ_OddQ(t *testing.T) {
	assert.Equal(t, NewOddQ(0, 0), NewEvenQ(0, 0).OddQ())
	assert.Equal(t, NewOddQ(1, 1), NewEvenQ(1, 2).OddQ())
}

func TestEvenQ_DoubleWidth(t *testing.T) {
	assert.Equal(t, NewDoubleWidth(0, 0), NewEvenQ(0, 0).DoubleWidth())
	assert.Equal(t, NewDoubleWidth(4, 0), NewEvenQ(2, 1).DoubleWidth())
}

func TestEvenQ_DoubleHeight(t *testing.T) {
	assert.Equal(t, NewDoubleHeight(0, 0), NewEvenQ(0, 0).DoubleHeight())
	assert.Equal(t, NewDoubleHeight(1, 3), NewEvenQ(1, 2).DoubleHeight())
}

func TestEvenQ_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		EvenQs{
			NewEvenQ(2, 1),
			NewEvenQ(1, 2),
			NewEvenQ(0, 1),
			NewEvenQ(0, 0),
			NewEvenQ(1, 0),
			NewEvenQ(2, 0),
		},
		NewEvenQ(1, 1).Neighbors(),
	)
}

func TestEvenQ_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewEvenQ(0, 0).DistanceTo(NewEvenQ(0, 0)))
	assert.Equal(t, 3, NewEvenQ(4, 1).DistanceTo(NewEvenQ(5, 4)))
}
