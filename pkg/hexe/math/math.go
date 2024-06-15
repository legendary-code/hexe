package math

import (
	hm "github.com/legendary-code/hexe/internal/hexe/math"
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

func CubeDistance(aq int, ar int, as int, bq int, br int, bs int) int {
	return hm.Maxi(hm.Absi(aq-bq), hm.Absi(ar-br), hm.Absi(as-bs))
}

func CubeRound(q float64, r float64, s float64) (int, int, int) {
	qr := math.Round(q)
	rr := math.Round(r)
	sr := math.Round(s)

	qDiff := math.Abs(qr - q)
	rDiff := math.Abs(rr - r)
	sDiff := math.Abs(sr - s)

	if qDiff > rDiff && qDiff > sDiff {
		qr = -rr - sr
	} else if rDiff > sDiff {
		rr = -qr - sr
	} else {
		sr = -qr - rr
	}

	return int(qr), int(rr), int(sr)
}

func lerp(a float64, b float64, t float64) float64 {
	return a + (b-a)*t
}

func CubeLerp(aq float64, ar float64, as float64, bq float64, br float64, bs float64, t float64) (float64, float64, float64) {
	q := lerp(aq, bq, t)
	r := lerp(ar, br, t)
	s := lerp(as, bs, t)
	return q, r, s
}

func CubeLineDraw(aq int, ar int, as int, bq int, br int, bs int) [][3]int {
	n := CubeDistance(aq, ar, as, bq, br, bs)
	coords := make([][3]int, 0)

	for i := 0; i < n; i++ {
		tn := 1.0 / float64(n) * float64(i)
		q, r, s := CubeLerp(float64(aq), float64(ar), float64(as), float64(bq), float64(br), float64(bs), tn)
		qr, rr, sr := CubeRound(q, r, s)
		coords = append(coords, [3]int{qr, rr, sr})
	}

	return coords
}
