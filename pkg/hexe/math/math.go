package math

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"golang.org/x/exp/constraints"
	"math"
)

const root3 = 1.7320508075688772935274463415059
const twoPi = 2.0 * math.Pi

func CalculateWidth[T constraints.Float](orientation consts.Orientation, size T) T {
	if orientation == consts.FlatTop {
		return 2.0 * size
	} else {
		return root3 * size
	}
}

func CalculateHeight[T constraints.Float](orientation consts.Orientation, size T) T {
	if orientation == consts.PointyTop {
		return 2.0 * size
	} else {
		return root3 * size
	}
}

func CalculateHorizontalSpacing[T constraints.Float](orientation consts.Orientation, size T) T {
	if orientation == consts.FlatTop {
		return 1.5 * size
	} else {
		return root3 * size
	}
}

func CalculateVerticalSpacing[T constraints.Float](orientation consts.Orientation, size T) T {
	if orientation == consts.PointyTop {
		return 1.5 * size
	} else {
		return root3 * size
	}
}

func wrapAroundRadians[T constraints.Float](radians T) T {
	radians += T(math.Ceil(float64(-radians)/twoPi) * twoPi)
	radians -= T(math.Floor(float64(radians)/twoPi) * twoPi)
	return radians
}

func CalculateCorners[T constraints.Float](center [2]T, size T, startRadians T, clockwise bool) [consts.Sides][2]T {
	corners := [consts.Sides][2]T{{}, {}, {}, {}, {}, {}}
	r := startRadians
	r = wrapAroundRadians(r)
	inc := T(consts.InnerAngleRadians)

	if clockwise {
		inc = -inc
	}

	for i := 0; i < consts.Sides; i++ {
		corners[i] = [2]T{
			center[0] + size*T(math.Cos(float64(r))),
			center[1] + size*T(math.Sin(float64(r))),
		}

		r += inc
		r = wrapAroundRadians(r)
	}

	return corners
}

func max(values ...int) int {
	m := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > m {
			m = values[i]
		}
	}

	return m
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func CubeDistance(aq int, ar int, as int, bq int, br int, bs int) int {
	return max(abs(aq-bq), abs(ar-br), abs(as-bs))
}
