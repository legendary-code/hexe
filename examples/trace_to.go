package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func traceToExample() {
	fig := plot.NewFigure()

	grid, walls := createArena()
	from := coord.NewAxial(-1, -1)
	to := coord.NewAxial(0, 2)
	trace := from.TraceTo(to, walls.Contains)

	fig.AddCoords(grid)
	fig.AddStyledCoords(walls, style.Color(colornames.Bisque))
	fig.AddStyledCoords(trace, style.Color(colornames.Lightgreen))

	_ = fig.RenderFile("images/trace_to.svg")
}
