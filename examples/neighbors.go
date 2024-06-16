package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func neighborsExample() {
	fig := plot.NewFigure()

	center := coord.ZeroAxial()
	neighbors := center.Neighbors()

	fig.AddStyledCoords(neighbors, style.Color(colornames.Lightblue))
	fig.AddCoord(center)

	_ = fig.RenderFile("images/neighbors.svg")
}
