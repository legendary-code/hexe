package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type OddQGrid[T any] interface {
	QRGrid[T, coord.OddQ]
}

func newOddQ[T any](grid *grid[T]) OddQGrid[T] {
	return &qrGrid[T, coord.OddQ]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coord.NewOddQ(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) coord.OddQ {
			return coord.NewAxial(q, r).OddQ()
		},
	}
}
