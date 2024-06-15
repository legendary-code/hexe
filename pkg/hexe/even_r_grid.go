package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type EvenRGrid[T any] interface {
	QRGrid[T, coord.EvenR, coord.EvenRs]
}

func newEvenR[T any](grid *grid[T]) EvenRGrid[T] {
	return &qrGrid[T, coord.EvenR, coord.EvenRs]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coord.NewEvenR(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) coord.EvenR {
			return coord.NewAxial(q, r).EvenR()
		},
	}
}
