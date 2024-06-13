package coords

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleWidth_Axial(t *testing.T) {
	assert.Equal(t, Axial(0, 0), DoubleWidth(0, 0).Axial())
	assert.Equal(t, Axial(-1, 4), DoubleWidth(2, 4).Axial())
}

func TestDoubleWidth_Cube(t *testing.T) {
	assert.Equal(t, Cube(0, 0, 0), DoubleWidth(0, 0).Cube())
	assert.Equal(t, Cube(-1, 4, -3), DoubleWidth(2, 4).Cube())
}

func TestDoubleWidth_OddR(t *testing.T) {
	assert.Equal(t, OddR(0, 0), DoubleWidth(0, 0).OddR())
	assert.Equal(t, OddR(1, 4), DoubleWidth(2, 4).OddR())
}

func TestDoubleWidth_EvenR(t *testing.T) {
	assert.Equal(t, EvenR(0, 0), DoubleWidth(0, 0).EvenR())
	assert.Equal(t, EvenR(1, 4), DoubleWidth(2, 4).EvenR())
}

func TestDoubleWidth_OddQ(t *testing.T) {
	assert.Equal(t, OddQ(0, 0), DoubleWidth(0, 0).OddQ())
	assert.Equal(t, OddQ(-1, 3), DoubleWidth(2, 4).OddQ())
}

func TestDoubleWidth_EvenQ(t *testing.T) {
	assert.Equal(t, EvenQ(0, 0), DoubleWidth(0, 0).EvenQ())
	assert.Equal(t, EvenQ(-1, 4), DoubleWidth(2, 4).EvenQ())
}

func TestDoubleWidth_DoubleHeight(t *testing.T) {
	assert.Equal(t, DoubleHeight(0, 0), DoubleWidth(0, 0).DoubleHeight())
	assert.Equal(t, DoubleHeight(-1, 7), DoubleWidth(2, 4).DoubleHeight())
}
