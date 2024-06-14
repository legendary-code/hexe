package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type EvenQGrid[T any] interface {
	QRGrid[T, coord.EvenQ]
}

func newEvenQ[T any](grid *grid[T]) EvenQGrid[T] {
	return &qrGrid[T, coord.EvenQ]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return coord.NewEvenQ(q, r).Axial().Unpack()
		},
		fromAxial: func(q int, r int) coord.EvenQ {
			return coord.NewAxial(q, r).EvenQ()
		},
	}
}
