package coords

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvenQ_Axial(t *testing.T) {
	assert.Equal(t, Axial(0, 0), EvenQ(0, 0).Axial())
	assert.Equal(t, Axial(1, 1), EvenQ(1, 2).Axial())
}

func TestEvenQ_Cube(t *testing.T) {
	assert.Equal(t, Cube(0, 0, 0), EvenQ(0, 0).Cube())
	assert.Equal(t, Cube(3, 3, -6), EvenQ(3, 5).Cube())
}

func TestEvenQ_OddR(t *testing.T) {
	assert.Equal(t, OddR(0, 0), EvenQ(0, 0).OddR())
	assert.Equal(t, OddR(0, 3), EvenQ(-1, 3).OddR())
}

func TestEvenQ_EvenR(t *testing.T) {
	assert.Equal(t, EvenR(0, 0), EvenQ(0, 0).EvenR())
	assert.Equal(t, EvenR(1, 3), EvenQ(-1, 3).EvenR())
}

func TestEvenQ_OddQ(t *testing.T) {
	assert.Equal(t, OddQ(0, 0), EvenQ(0, 0).OddQ())
	assert.Equal(t, OddQ(1, 1), EvenQ(1, 2).OddQ())
}

func TestEvenQ_DoubleWidth(t *testing.T) {
	assert.Equal(t, DoubleWidth(0, 0), EvenQ(0, 0).DoubleWidth())
	assert.Equal(t, DoubleWidth(4, 0), EvenQ(2, 1).DoubleWidth())
}

func TestEvenQ_DoubleHeight(t *testing.T) {
	assert.Equal(t, DoubleHeight(0, 0), EvenQ(0, 0).DoubleHeight())
	assert.Equal(t, DoubleHeight(1, 3), EvenQ(1, 2).DoubleHeight())
}
