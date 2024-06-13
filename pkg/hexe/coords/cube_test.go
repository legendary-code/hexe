package coords

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCube_Axial(t *testing.T) {
	assert.Equal(t, Axial(0, 0), Cube(0, 0, 0).Axial())
	assert.Equal(t, Axial(3, 5), Cube(3, 5, -8).Axial())
}

func TestCube_OddR(t *testing.T) {
	assert.Equal(t, OddR(0, 0), Cube(0, 0, 0).OddR())
	assert.Equal(t, OddR(0, 3), Cube(-1, 3, -2).OddR())
}

func TestCube_EvenR(t *testing.T) {
	assert.Equal(t, EvenR(0, 0), Cube(0, 0, 0).EvenR())
	assert.Equal(t, EvenR(1, 3), Cube(-1, 3, -2).EvenR())
}

func TestCube_OddQ(t *testing.T) {
	assert.Equal(t, OddQ(0, 0), Cube(0, 0, 0).OddQ())
	assert.Equal(t, OddQ(1, 2), Cube(1, 2, -3).OddQ())
}

func TestCube_EvenQ(t *testing.T) {
	assert.Equal(t, EvenQ(0, 0), Cube(0, 0, 0).EvenQ())
	assert.Equal(t, EvenQ(1, 3), Cube(1, 2, -3).EvenQ())
}

func TestCube_DoubleWidth(t *testing.T) {
	assert.Equal(t, DoubleWidth(0, 0), Cube(0, 0, 0).DoubleWidth())
	assert.Equal(t, DoubleWidth(5, 1), Cube(2, 1, -3).DoubleWidth())
}

func TestCube_DoubleHeight(t *testing.T) {
	assert.Equal(t, DoubleHeight(0, 0), Cube(0, 0, 0).DoubleHeight())
	assert.Equal(t, DoubleHeight(1, 5), Cube(1, 2, -3).DoubleHeight())
}
