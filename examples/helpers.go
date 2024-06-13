package main

import "github.com/legendary-code/hexe/pkg/hexe/coord"

func createArena() (*coord.Axials, *coord.Axials) {
	grid := coord.ZeroAxial().MovementRange(3)
	walls := coord.ZeroAxial().Ring(3)
	walls.Union(
		coord.NewAxials(
			coord.NewAxial(-2, 1),
			coord.NewAxial(-1, 1),
			coord.NewAxial(0, 0),
		),
	)

	return grid, walls
}
