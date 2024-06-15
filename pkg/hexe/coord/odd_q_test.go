package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddQ_Axial(t *testing.T) {
	assert.Equal(t, NewAxial(0, 0), NewOddQ(0, 0).Axial())
	assert.Equal(t, NewAxial(1, 2), NewOddQ(1, 2).Axial())
}

func TestOddQ_Cube(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewOddQ(0, 0).Cube())
	assert.Equal(t, NewCube(3, 4, -7), NewOddQ(3, 5).Cube())
}

func TestOddQ_OddR(t *testing.T) {
	assert.Equal(t, NewOddR(0, 0), NewOddQ(0, 0).OddR())
	assert.Equal(t, NewOddR(1, 4), NewOddQ(-1, 3).OddR())
}

func TestOddQ_EvenQ(t *testing.T) {
	assert.Equal(t, NewEvenQ(0, 0), NewOddQ(0, 0).EvenQ())
	assert.Equal(t, NewEvenQ(-1, 4), NewOddQ(-1, 3).EvenQ())
}

func TestOddQ_EvenR(t *testing.T) {
	assert.Equal(t, NewEvenR(0, 0), NewOddQ(0, 0).EvenR())
	assert.Equal(t, NewEvenR(2, 2), NewOddQ(1, 2).EvenR())
}

func TestOddQ_DoubleWidth(t *testing.T) {
	assert.Equal(t, NewDoubleWidth(0, 0), NewOddQ(0, 0).DoubleWidth())
	assert.Equal(t, NewDoubleWidth(4, 0), NewOddQ(2, 1).DoubleWidth())
}

func TestOddQ_DoubleHeight(t *testing.T) {
	assert.Equal(t, NewDoubleHeight(0, 0), NewOddQ(0, 0).DoubleHeight())
	assert.Equal(t, NewDoubleHeight(1, 5), NewOddQ(1, 2).DoubleHeight())
}

func TestOddQ_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		OddQs{
			NewOddQ(2, 2),
			NewOddQ(1, 2),
			NewOddQ(0, 2),
			NewOddQ(0, 1),
			NewOddQ(1, 0),
			NewOddQ(2, 1),
		},
		NewOddQ(1, 1).Neighbors(),
	)
}

func TestOddQ_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewOddQ(0, 0).DistanceTo(NewOddQ(0, 0)))
	assert.Equal(t, 3, NewOddQ(3, 1).DistanceTo(NewOddQ(2, 4)))
}
