package math

import (
	"golang.org/x/exp/constraints"
	"math"
)

const TwoPi = 2.0 * math.Pi

func Maxi(values ...int) int {
	m := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > m {
			m = values[i]
		}
	}

	return m
}

func Mini(values ...int) int {
	m := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] < m {
			m = values[i]
		}
	}

	return m
}

func Absi(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func Lerp(a float64, b float64, t float64) float64 {
	return a + (b-a)*t
}

func WrapAroundRadians[T constraints.Float](radians T) T {
	radians += T(math.Ceil(float64(-radians)/TwoPi) * TwoPi)
	radians -= T(math.Floor(float64(radians)/TwoPi) * TwoPi)
	return radians
}
