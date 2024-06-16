package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func ringExample() {
	fig := plot.NewFigure()

	grid, walls := createArena()
	walls = walls.ReflectR()
	ring := coord.ZeroAxial().Ring(1)

	fig.AddCoords(grid)
	fig.AddStyledCoords(walls, style.Color(colornames.Bisque))
	fig.AddStyledCoords(ring, style.Color(colornames.Lightcoral))

	_ = fig.RenderFile("images/ring.svg")
}
