package hexe

import (
	"github.com/legendary-code/hexe/pkg/hexe/coords"
)

type Cube coords.CoordQRS

type CubeGrid[T any] interface {
	QRSGrid[T, Cube]
}

func newCubeGrid[T any](grid *grid[T]) CubeGrid[T] {
	return &qrsGrid[T, Cube]{
		grid: grid,
		toAxial: func(q int, r int, s int) (int, int) {
			return coords.Cube(q, r, s).Axial().Unpack()
		},
		fromAxial: func(q int, r int) (int, int, int) {
			return coords.Axial(q, r).Cube().Unpack()
		},
	}
}
