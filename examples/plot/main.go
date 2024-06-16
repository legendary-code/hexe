package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"golang.org/x/image/colornames"
)

func main() {
	f := plot.NewFigure()

	center := coord.NewAxial(0, 0)
	grid := center.MovementRange(3)

	waterStyle := &plot.CellStyle{Color: colornames.Lightblue, FontSize: 40}
	landStyle := &plot.CellStyle{Color: colornames.Sandybrown, FontSize: 40}
	traceStyle := &plot.CellStyle{Color: colornames.Red}

	f.AddStyledCoords(
		grid,
		func(coord coord.Coord) (string, *plot.CellStyle) { return "🌊", waterStyle },
	)

	f.AddStyledCoords(
		coord.Axials{
			coord.NewAxial(0, 0),
			coord.NewAxial(1, 0),
			coord.NewAxial(1, -1),
			coord.NewAxial(0, -1),
			coord.NewAxial(-1, 0),
		},
		func(coord coord.Coord) (string, *plot.CellStyle) { return "🏝️", landStyle },
	)

	f.AddStyledCoord(
		coord.NewAxial(1, 1),
		func(coord coord.Coord) (string, *plot.CellStyle) { return "🏖️", landStyle },
	)

	f.AddStyledCoords(
		coord.NewAxial(4, 4).TraceTo(coord.NewAxial(-4, -4), func(coord coord.Axial) bool {
			return coord.Q() < 0
		}),
		func(coord coord.Coord) (string, *plot.CellStyle) {
			return "", traceStyle
		},
	)

	_ = f.RenderFile("example.svg")
}
