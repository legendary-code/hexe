package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAxial_Cube(t *testing.T) {
	assert.Equal(t, NewCube(0, 0, 0), NewAxial(0, 0).Cube())
	assert.Equal(t, NewCube(3, 5, -8), NewAxial(3, 5).Cube())
}

func TestAxial_OddR(t *testing.T) {
	assert.Equal(t, NewOddR(0, 0), NewAxial(0, 0).OddR())
	assert.Equal(t, NewOddR(0, 3), NewAxial(-1, 3).OddR())
}

func TestAxial_EvenR(t *testing.T) {
	assert.Equal(t, NewEvenR(0, 0), NewAxial(0, 0).EvenR())
	assert.Equal(t, NewEvenR(1, 3), NewAxial(-1, 3).EvenR())
}

func TestAxial_OddQ(t *testing.T) {
	assert.Equal(t, NewOddQ(0, 0), NewAxial(0, 0).OddQ())
	assert.Equal(t, NewOddQ(1, 2), NewAxial(1, 2).OddQ())
}

func TestAxial_EvenQ(t *testing.T) {
	assert.Equal(t, NewEvenQ(0, 0), NewAxial(0, 0).EvenQ())
	assert.Equal(t, NewEvenQ(1, 3), NewAxial(1, 2).EvenQ())
}

func TestAxial_DoubleWidth(t *testing.T) {
	assert.Equal(t, NewDoubleWidth(0, 0), NewAxial(0, 0).DoubleWidth())
	assert.Equal(t, NewDoubleWidth(5, 1), NewAxial(2, 1).DoubleWidth())
}

func TestAxial_DoubleHeight(t *testing.T) {
	assert.Equal(t, NewDoubleHeight(0, 0), NewAxial(0, 0).DoubleHeight())
	assert.Equal(t, NewDoubleHeight(1, 5), NewAxial(1, 2).DoubleHeight())
}

func TestAxial_Neighbors(t *testing.T) {
	assert.Equal(
		t,
		Axials{
			NewAxial(2, 1),
			NewAxial(1, 2),
			NewAxial(0, 2),
			NewAxial(0, 1),
			NewAxial(1, 0),
			NewAxial(2, 0),
		},
		NewAxial(1, 1).Neighbors(),
	)
}

func TestAxial_DistanceTo(t *testing.T) {
	assert.Equal(t, 0, NewAxial(0, 0).DistanceTo(NewAxial(0, 0)))
	assert.Equal(t, 3, NewAxial(1, -2).DistanceTo(NewAxial(2, 0)))
}

func TestAxial_FloodFill(t *testing.T) {
	actual := NewAxial(1, 0).FloodFill(2, func(coord Axial) bool {
		return coord.Q()+coord.R() == 0
	}).Sorted()
	expected := Axials{
		NewAxial(-1, 2),
		NewAxial(0, 1),
		NewAxial(0, 2),
		NewAxial(1, 0),
		NewAxial(1, 1),
		NewAxial(1, 2),
		NewAxial(2, -1),
		NewAxial(2, 0),
		NewAxial(2, 1),
		NewAxial(3, -2),
		NewAxial(3, -1),
		NewAxial(3, 0),
	}
	assert.Equal(t, expected, actual)
}
