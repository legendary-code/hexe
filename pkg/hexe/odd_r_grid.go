package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type OddRGrid[T any] interface {
	QRGrid[T, coord.OddR]
}

func newOddR[T any](grid *grid[T]) OddRGrid[T] {
	return &qrGrid[T, coord.OddR]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coord.NewOddR(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) coord.OddR {
			return coord.NewAxial(q, r).OddR()
		},
	}
}
