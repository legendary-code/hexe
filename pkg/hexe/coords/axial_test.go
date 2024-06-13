package coords

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAxial_Cube(t *testing.T) {
	assert.Equal(t, Cube(0, 0, 0), Axial(0, 0).Cube())
	assert.Equal(t, Cube(3, 5, -8), Axial(3, 5).Cube())
}

func TestAxial_OddR(t *testing.T) {
	assert.Equal(t, OddR(0, 0), Axial(0, 0).OddR())
	assert.Equal(t, OddR(0, 3), Axial(-1, 3).OddR())
}

func TestAxial_EvenR(t *testing.T) {
	assert.Equal(t, EvenR(0, 0), Axial(0, 0).EvenR())
	assert.Equal(t, EvenR(1, 3), Axial(-1, 3).EvenR())
}

func TestAxial_OddQ(t *testing.T) {
	assert.Equal(t, OddQ(0, 0), Axial(0, 0).OddQ())
	assert.Equal(t, OddQ(1, 2), Axial(1, 2).OddQ())
}

func TestAxial_EvenQ(t *testing.T) {
	assert.Equal(t, EvenQ(0, 0), Axial(0, 0).EvenQ())
	assert.Equal(t, EvenQ(1, 3), Axial(1, 2).EvenQ())
}

func TestAxial_DoubleWidth(t *testing.T) {
	assert.Equal(t, DoubleWidth(0, 0), Axial(0, 0).DoubleWidth())
	assert.Equal(t, DoubleWidth(5, 1), Axial(2, 1).DoubleWidth())
}

func TestAxial_DoubleHeight(t *testing.T) {
	assert.Equal(t, DoubleHeight(0, 0), Axial(0, 0).DoubleHeight())
	assert.Equal(t, DoubleHeight(1, 5), Axial(1, 2).DoubleHeight())
}

func TestAxial_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		[consts.Sides]CoordQR{
			Axial(2, 1),
			Axial(0, 1),
			Axial(2, 0),
			Axial(0, 2),
			Axial(1, 0),
			Axial(1, 2),
		},
		Axial(1, 1).Neighbors(),
	)
}
