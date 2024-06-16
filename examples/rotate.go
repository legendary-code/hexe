package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func rotateExample() {
	fig := plot.NewFigure()

	grid, walls := createArena()
	walls = walls.Rotate(coord.ZeroAxial(), 2)

	fig.AddCoords(grid)
	fig.AddStyledCoords(walls, style.Color(colornames.Bisque))

	_ = fig.RenderFile("images/rotate.svg")
}
