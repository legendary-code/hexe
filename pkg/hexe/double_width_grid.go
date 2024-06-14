package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type DoubleWidthGrid[T any] interface {
	QRGrid[T, coord.DoubleWidth]
}

func newDoubleWidth[T any](grid *grid[T]) DoubleWidthGrid[T] {
	return &qrGrid[T, coord.DoubleWidth]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coord.NewDoubleWidth(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) coord.DoubleWidth {
			return coord.NewAxial(q, r).DoubleWidth()
		},
	}
}
