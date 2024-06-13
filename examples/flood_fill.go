package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func floodFillExample() {
	fig := plot.NewFigure()

	grid, walls := createArena()
	center := coord.NewAxial(-1, -1)
	fill := center.FloodFill(3, walls.Contains)

	fig.AddCoords(grid)
	fig.AddStyledCoords(walls, style.Color(colornames.Bisque))
	fig.AddStyledCoords(fill, style.Color(colornames.Lightgreen))

	_ = fig.RenderFile("images/flood_fill.svg")
}
