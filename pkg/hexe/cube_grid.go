package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type CubeGrid[T any] interface {
	QRSGrid[T, coord.Cube]
}

func newCubeGrid[T any](grid *grid[T]) CubeGrid[T] {
	return &qrsGrid[T, coord.Cube]{
		grid: grid,
		toAxial: func(q int, r int, s int) (int, int) {
			return coord.NewCube(q, r, s).Axial().Unpack()
		},
		fromAxial: func(q int, r int) coord.Cube {
			return coord.NewAxial(q, r).Cube()
		},
	}
}
