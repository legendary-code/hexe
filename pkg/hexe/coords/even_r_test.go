package coords

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvenR_Axial(t *testing.T) {
	assert.Equal(t, Axial(0, 0), EvenR(0, 0).Axial())
	assert.Equal(t, Axial(0, 2), EvenR(1, 2).Axial())
}

func TestEvenR_Cube(t *testing.T) {
	assert.Equal(t, Cube(0, 0, 0), EvenR(0, 0).Cube())
	assert.Equal(t, Cube(0, 5, -5), EvenR(3, 5).Cube())
}

func TestEvenR_OddR(t *testing.T) {
	assert.Equal(t, OddR(0, 0), EvenR(0, 0).OddR())
	assert.Equal(t, OddR(-2, 3), EvenR(-1, 3).OddR())
}

func TestEvenR_EvenQ(t *testing.T) {
	assert.Equal(t, EvenQ(0, 0), EvenR(0, 0).EvenQ())
	assert.Equal(t, EvenQ(-3, 2), EvenR(-1, 3).EvenQ())
}

func TestEvenR_OddQ(t *testing.T) {
	assert.Equal(t, OddQ(0, 0), EvenR(0, 0).OddQ())
	assert.Equal(t, OddQ(0, 2), EvenR(1, 2).OddQ())
}

func TestEvenR_DoubleWidth(t *testing.T) {
	assert.Equal(t, DoubleWidth(0, 0), EvenR(0, 0).DoubleWidth())
	assert.Equal(t, DoubleWidth(3, 1), EvenR(2, 1).DoubleWidth())
}

func TestEvenR_DoubleHeight(t *testing.T) {
	assert.Equal(t, DoubleHeight(0, 0), EvenR(0, 0).DoubleHeight())
	assert.Equal(t, DoubleHeight(0, 4), EvenR(1, 2).DoubleHeight())
}

func TestEvenR_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		[consts.Sides]CoordQR{
			EvenR(2, 1),
			EvenR(0, 1),
			EvenR(1, 0),
			EvenR(0, 2),
			EvenR(0, 0),
			EvenR(1, 2),
		},
		EvenR(1, 1).Neighbors(),
	)
}
