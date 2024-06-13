package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type DoubleWidth coords.CoordQR

type DoubleWidthGrid[T any] interface {
	QRGrid[T, DoubleWidth]
}

func newDoubleWidth[T any](grid *grid[T]) DoubleWidthGrid[T] {
	return &qrGrid[T, DoubleWidth]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coords.DoubleWidth(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) (int, int) {
			return coords.Axial(q, r).DoubleWidth().Unpack()
		},
	}
}
