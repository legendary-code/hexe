package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coords"

type Axial coords.CoordQR

type AxialGrid[T any] interface {
	QRGrid[T, Axial]
}

func newAxialGrid[T any](grid *grid[T]) AxialGrid[T] {
	return &qrGrid[T, Axial]{
		grid: grid,
		toAxial: func(q int, r int) (int, int) {
			return q, r
		},
		fromAxial: func(q int, r int) (int, int) {
			return q, r
		},
	}
}
