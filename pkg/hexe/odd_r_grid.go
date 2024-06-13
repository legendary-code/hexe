package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type OddR coords.CoordQR

type OddRGrid[T any] interface {
	QRGrid[T, OddR]
}

func newOddR[T any](grid *grid[T]) OddRGrid[T] {
	return &qrGrid[T, OddR]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coords.OddR(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) (int, int) {
			return coords.Axial(q, r).OddR().Unpack()
		},
	}
}
