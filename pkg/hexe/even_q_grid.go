package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type EvenQ coords.CoordQR

type EvenQGrid[T any] interface {
	QRGrid[T, EvenQ]
}

func newEvenQ[T any](grid *grid[T]) EvenQGrid[T] {
	return &qrGrid[T, EvenQ]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coords.EvenQ(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) (int, int) {
			return coords.Axial(q, r).EvenQ().Unpack()
		},
	}
}
