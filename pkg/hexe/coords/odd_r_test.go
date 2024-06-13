package coords

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddR_Axial(t *testing.T) {
	assert.Equal(t, Axial(0, 0), OddR(0, 0).Axial())
	assert.Equal(t, Axial(0, 2), OddR(1, 2).Axial())
}

func TestOddR_Cube(t *testing.T) {
	assert.Equal(t, Cube(0, 0, 0), OddR(0, 0).Cube())
	assert.Equal(t, Cube(1, 5, -6), OddR(3, 5).Cube())
}

func TestOddR_OddQ(t *testing.T) {
	assert.Equal(t, OddQ(0, 0), OddR(0, 0).OddQ())
	assert.Equal(t, OddQ(-2, 2), OddR(-1, 3).OddQ())
}

func TestOddR_EvenQ(t *testing.T) {
	assert.Equal(t, EvenQ(0, 0), OddR(0, 0).EvenQ())
	assert.Equal(t, EvenQ(-2, 2), OddR(-1, 3).EvenQ())
}

func TestOddR_EvenR(t *testing.T) {
	assert.Equal(t, EvenR(0, 0), OddR(0, 0).EvenR())
	assert.Equal(t, EvenR(1, 2), OddR(1, 2).EvenR())
}

func TestOddR_DoubleWidth(t *testing.T) {
	assert.Equal(t, DoubleWidth(0, 0), OddR(0, 0).DoubleWidth())
	assert.Equal(t, DoubleWidth(5, 1), OddR(2, 1).DoubleWidth())
}

func TestOddR_DoubleHeight(t *testing.T) {
	assert.Equal(t, DoubleHeight(0, 0), OddR(0, 0).DoubleHeight())
	assert.Equal(t, DoubleHeight(0, 4), OddR(1, 2).DoubleHeight())
}

func TestOddR_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		[consts.Sides]CoordQR{
			OddR(2, 1),
			OddR(0, 1),
			OddR(2, 0),
			OddR(1, 2),
			OddR(1, 0),
			OddR(2, 2),
		},
		OddR(1, 1).Neighbors(),
	)
}
