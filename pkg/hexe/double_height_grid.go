package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type DoubleHeightGrid[T any] interface {
	QRGrid[T, coord.DoubleHeight]
}

func newDoubleHeight[T any](grid *grid[T]) DoubleHeightGrid[T] {
	return &qrGrid[T, coord.DoubleHeight]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coord.NewDoubleHeight(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) coord.DoubleHeight {
			return coord.NewAxial(q, r).DoubleHeight()
		},
	}
}
