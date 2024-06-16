package math

import (
	"fmt"
	hm "github.com/legendary-code/hexe/internal/hexe/math"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"golang.org/x/exp/constraints"
	"math"
)

const root3 = 1.7320508075688772935274463415059

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

func CalculateCorners[T constraints.Float](center [2]T, size T, startRadians T, clockwise bool) [consts.Sides][2]T {
	corners := [consts.Sides][2]T{{}, {}, {}, {}, {}, {}}
	r := startRadians
	r = hm.WrapAroundRadians(r)
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
		r = hm.WrapAroundRadians(r)
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

func CubeLerp(aq float64, ar float64, as float64, bq float64, br float64, bs float64, t float64) (float64, float64, float64) {
	q := hm.Lerp(aq, bq, t)
	r := hm.Lerp(ar, br, t)
	s := hm.Lerp(as, bs, t)
	return q, r, s
}

func CubeLineDraw(aq int, ar int, as int, bq int, br int, bs int) [][3]int {
	n := CubeDistance(aq, ar, as, bq, br, bs)
	coords := make([][3]int, 0)

	if n == 0 {
		coords = append(coords, [3]int{aq, ar, as})
		return coords
	}

	for i := 0; i <= n; i++ {
		tn := 1.0 / float64(n) * float64(i)
		q, r, s := CubeLerp(float64(aq), float64(ar), float64(as), float64(bq), float64(br), float64(bs), tn)
		qr, rr, sr := CubeRound(q, r, s)
		coords = append(coords, [3]int{qr, rr, sr})
	}

	return coords
}

func AxialToPixel(q int, r int, size float64, orientation consts.Orientation) (float64, float64) {
	switch orientation {
	case consts.FlatTop:
		return 1.5 * size * float64(q), size * (root3*0.5*float64(q) + root3*float64(r))
	case consts.PointyTop:
		return size * (root3*float64(q) + root3*0.5*float64(r)), 1.5 * size * float64(r)
	default:
		panic(fmt.Sprintf("unknown orientation type: %v", orientation))
	}
}

func ClampHexAngle(angle int) int {
	if angle < 0 {
		angle = angle + (-angle/consts.Sides)*consts.Sides + consts.Sides
	}

	return angle % consts.Sides
}
