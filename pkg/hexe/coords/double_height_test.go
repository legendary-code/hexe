package coords

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleHeight_Axial(t *testing.T) {
	assert.Equal(t, Axial(0, 0), DoubleHeight(0, 0).Axial())
	assert.Equal(t, Axial(2, 1), DoubleHeight(2, 4).Axial())
}

func TestDoubleHeight_Cube(t *testing.T) {
	assert.Equal(t, Cube(0, 0, 0), DoubleHeight(0, 0).Cube())
	assert.Equal(t, Cube(2, 1, -3), DoubleHeight(2, 4).Cube())
}

func TestDoubleHeight_OddR(t *testing.T) {
	assert.Equal(t, OddR(0, 0), DoubleHeight(0, 0).OddR())
	assert.Equal(t, OddR(2, 1), DoubleHeight(2, 4).OddR())
}

func TestDoubleHeight_EvenR(t *testing.T) {
	assert.Equal(t, EvenR(0, 0), DoubleHeight(0, 0).EvenR())
	assert.Equal(t, EvenR(3, 1), DoubleHeight(2, 4).EvenR())
}

func TestDoubleHeight_OddQ(t *testing.T) {
	assert.Equal(t, OddQ(0, 0), DoubleHeight(0, 0).OddQ())
	assert.Equal(t, OddQ(2, 2), DoubleHeight(2, 4).OddQ())
}

func TestDoubleHeight_EvenQ(t *testing.T) {
	assert.Equal(t, EvenQ(0, 0), DoubleHeight(0, 0).EvenQ())
	assert.Equal(t, EvenQ(2, 2), DoubleHeight(2, 4).EvenQ())
}

func TestDoubleHeight_DoubleWidth(t *testing.T) {
	assert.Equal(t, DoubleWidth(0, 0), DoubleHeight(0, 0).DoubleWidth())
	assert.Equal(t, DoubleWidth(5, 1), DoubleHeight(2, 4).DoubleWidth())
}

func TestDoubleHeight_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		[consts.Sides]CoordQR{
			DoubleHeight(2, 2),
			DoubleHeight(0, 0),
			DoubleHeight(2, 0),
			DoubleHeight(0, 2),
			DoubleHeight(1, -1),
			DoubleHeight(1, 3),
		},
		DoubleHeight(1, 1).Neighbors(),
	)
}
