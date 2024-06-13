package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type DoubleHeight coords.CoordQR

type DoubleHeightGrid[T any] interface {
	QRGrid[T, DoubleHeight]
}

func newDoubleHeight[T any](grid *grid[T]) DoubleHeightGrid[T] {
	return &qrGrid[T, DoubleHeight]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coords.DoubleHeight(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) (int, int) {
			return coords.Axial(q, r).DoubleHeight().Unpack()
		},
	}
}
