package main

import "github.com/legendary-code/hexe/pkg/hexe/coord"

func createArena() (coord.Axials, coord.Axials) {
	return coord.ZeroAxial().MovementRange(3),
		coord.ZeroAxial().Ring(3).UnionWith(
			coord.Axials{
				coord.NewAxial(-2, 1),
				coord.NewAxial(-1, 1),
				coord.NewAxial(0, 0),
			},
		)
}
