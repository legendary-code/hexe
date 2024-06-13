package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func movementRangeExample() {
	fig := plot.NewFigure()

	center := coord.ZeroAxial()
	movementRange := center.MovementRange(2)

	fig.AddStyledCoords(movementRange, style.Color(colornames.Lightblue))
	fig.AddCoord(center)

	_ = fig.RenderFile("images/movement_range.svg")
}
