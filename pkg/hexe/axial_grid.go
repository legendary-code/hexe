package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type AxialGrid[T any] interface {
	QRGrid[T, coord.Axial, coord.Axials]
}

func newAxialGrid[T any](grid *grid[T]) AxialGrid[T] {
	return &qrGrid[T, coord.Axial, coord.Axials]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return q, r
		},
		fromAxial: func(q int, r int) coord.Axial {
			return coord.NewAxial(q, r)
		},
	}
}
