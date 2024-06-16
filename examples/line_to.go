package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func lineToExample() {
	fig := plot.NewFigure()

	grid := coord.ZeroAxial().MovementRange(3)
	from := coord.NewAxial(-1, -1)
	to := coord.NewAxial(2, 0)
	line := from.LineTo(to)

	fig.AddCoords(grid)
	fig.AddStyledCoords(line, style.Color(colornames.Lightgreen))

	_ = fig.RenderFile("images/line_to.svg")
}
