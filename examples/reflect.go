package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func reflectExample() {
	fig := plot.NewFigure()

	grid, walls := createArena()
	walls = walls.ReflectR()

	fig.AddCoords(grid)
	fig.AddStyledCoords(walls, style.Color(colornames.Bisque))

	_ = fig.RenderFile("images/reflect.svg")
}
