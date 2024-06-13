package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type OddQ coords.CoordQR

type OddQGrid[T any] interface {
	QRGrid[T, OddQ]
}

func newOddQ[T any](grid *grid[T]) OddQGrid[T] {
	return &qrGrid[T, OddQ]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coords.OddQ(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) (int, int) {
			return coords.Axial(q, r).OddQ().Unpack()
		},
	}
}
