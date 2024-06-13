package coords

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddQ_Axial(t *testing.T) {
	assert.Equal(t, Axial(0, 0), OddQ(0, 0).Axial())
	assert.Equal(t, Axial(1, 2), OddQ(1, 2).Axial())
}

func TestOddQ_Cube(t *testing.T) {
	assert.Equal(t, Cube(0, 0, 0), OddQ(0, 0).Cube())
	assert.Equal(t, Cube(3, 4, -7), OddQ(3, 5).Cube())
}

func TestOddQ_OddR(t *testing.T) {
	assert.Equal(t, OddR(0, 0), OddQ(0, 0).OddR())
	assert.Equal(t, OddR(1, 4), OddQ(-1, 3).OddR())
}

func TestOddQ_EvenQ(t *testing.T) {
	assert.Equal(t, EvenQ(0, 0), OddQ(0, 0).EvenQ())
	assert.Equal(t, EvenQ(-1, 4), OddQ(-1, 3).EvenQ())
}

func TestOddQ_EvenR(t *testing.T) {
	assert.Equal(t, EvenR(0, 0), OddQ(0, 0).EvenR())
	assert.Equal(t, EvenR(2, 2), OddQ(1, 2).EvenR())
}

func TestOddQ_DoubleWidth(t *testing.T) {
	assert.Equal(t, DoubleWidth(0, 0), OddQ(0, 0).DoubleWidth())
	assert.Equal(t, DoubleWidth(4, 0), OddQ(2, 1).DoubleWidth())
}

func TestOddQ_DoubleHeight(t *testing.T) {
	assert.Equal(t, DoubleHeight(0, 0), OddQ(0, 0).DoubleHeight())
	assert.Equal(t, DoubleHeight(1, 5), OddQ(1, 2).DoubleHeight())
}
