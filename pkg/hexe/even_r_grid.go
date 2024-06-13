package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type EvenR coords.CoordQR

type EvenRGrid[T any] interface {
	QRGrid[T, EvenR]
}

func newEvenR[T any](grid *grid[T]) EvenRGrid[T] {
	return &qrGrid[T, EvenR]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coords.EvenR(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) (int, int) {
			return coords.Axial(q, r).EvenR().Unpack()
		},
	}
}
