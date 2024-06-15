package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddR_Axial(t *testing.T) {
	assert.Equal(t, NewAxial(0, 0), NewOddR(0, 0).Axial())
	assert.Equal(t, NewAxial(0, 2), NewOddR(1, 2).Axial())
}

func TestOddR_Cube(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewOddR(0, 0).Cube())
	assert.Equal(t, NewCube(1, 5, -6), NewOddR(3, 5).Cube())
}

func TestOddR_OddQ(t *testing.T) {
	assert.Equal(t, NewOddQ(0, 0), NewOddR(0, 0).OddQ())
	assert.Equal(t, NewOddQ(-2, 2), NewOddR(-1, 3).OddQ())
}

func TestOddR_EvenQ(t *testing.T) {
	assert.Equal(t, NewEvenQ(0, 0), NewOddR(0, 0).EvenQ())
	assert.Equal(t, NewEvenQ(-2, 2), NewOddR(-1, 3).EvenQ())
}

func TestOddR_EvenR(t *testing.T) {
	assert.Equal(t, NewEvenR(0, 0), NewOddR(0, 0).EvenR())
	assert.Equal(t, NewEvenR(1, 2), NewOddR(1, 2).EvenR())
}

func TestOddR_DoubleWidth(t *testing.T) {
	assert.Equal(t, NewDoubleWidth(0, 0), NewOddR(0, 0).DoubleWidth())
	assert.Equal(t, NewDoubleWidth(5, 1), NewOddR(2, 1).DoubleWidth())
}

func TestOddR_DoubleHeight(t *testing.T) {
	assert.Equal(t, NewDoubleHeight(0, 0), NewOddR(0, 0).DoubleHeight())
	assert.Equal(t, NewDoubleHeight(0, 4), NewOddR(1, 2).DoubleHeight())
}

func TestOddR_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		OddRs{
			NewOddR(2, 1),
			NewOddR(2, 2),
			NewOddR(1, 2),
			NewOddR(0, 1),
			NewOddR(1, 0),
			NewOddR(2, 0),
		},
		NewOddR(1, 1).Neighbors(),
	)
}

func TestOddR_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewOddR(0, 0).DistanceTo(NewOddR(0, 0)))
	assert.Equal(t, 3, NewOddR(3, 1).DistanceTo(NewOddR(1, 2)))
}
