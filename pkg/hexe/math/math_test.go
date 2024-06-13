package math

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCalculateWidth(t *testing.T) {
	assert.Equal(t, 2.0, CalculateWidth(consts.FlatTop, 1.0))
	assert.Equal(t, root3, CalculateWidth(consts.PointyTop, 1.0))
}

func TestCalculateHeight(t *testing.T) {
	assert.Equal(t, 2.0, CalculateHeight(consts.PointyTop, 1.0))
	assert.Equal(t, root3, CalculateHeight(consts.FlatTop, 1.0))
}

func TestCalculateHorizontalSpacing(t *testing.T) {
	assert.Equal(t, 1.5, CalculateHorizontalSpacing(consts.FlatTop, 1.0))
	assert.Equal(t, root3, CalculateHorizontalSpacing(consts.PointyTop, 1.0))
}

func TestCalculateVerticalSpacing(t *testing.T) {
	assert.Equal(t, 1.5, CalculateVerticalSpacing(consts.PointyTop, 1.0))
	assert.Equal(t, root3, CalculateVerticalSpacing(consts.FlatTop, 1.0))
}

func TestCalculateCorners(t *testing.T) {
	expected := [6][2]float64{
		{2.0, 1.0},
		{1.5, 1.0 - math.Sqrt(3)/2.0},
		{0.5, 1.0 - math.Sqrt(3)/2.0},
		{0.0, 1.0},
		{0.5, 1 + math.Sqrt(3)/2.0},
		{1.5, 1 + math.Sqrt(3)/2.0},
	}
	actual := CalculateCorners([2]float64{1.0, 1.0}, 1.0, 0.0, true)

	for s := 0; s < consts.Sides; s++ {
		assert.InDelta(t, expected[s][0], actual[s][0], 0.001)
		assert.InDelta(t, expected[s][1], actual[s][1], 0.001)
	}
}
