package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func plotExample() {
	fig := plot.NewFigure()

	center := coord.NewAxial(0, 0)
	grid := center.MovementRange(3)

	waterStyle := style.Color(colornames.Lightblue).FontSize(40).Name("🌊")
	landStyle := style.Color(colornames.Sandybrown).FontSize(40).Name("🏝️")

	fig.AddStyledCoords(
		grid,
		waterStyle,
	)

	fig.AddStyledCoords(
		coord.Axials{
			coord.NewAxial(0, 0),
			coord.NewAxial(1, 0),
			coord.NewAxial(1, -1),
			coord.NewAxial(0, -1),
			coord.NewAxial(-1, 0),
		},
		landStyle,
	)

	fig.AddStyledCoord(
		coord.NewAxial(1, 1),
		landStyle.Name("🏖️"),
	)

	_ = fig.RenderFile("images/plot.svg")
}
