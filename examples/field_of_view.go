package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
	"image/color"
)

func fieldOfViewExample() {
	fig := plot.NewFigure()
	grid, walls := createArena()

	person := coord.NewAxial(-1, 2)

	blocked := func(coord coord.Axial) bool {
		for _, wall := range walls {
			if wall == coord {
				return true
			}
		}

		return false
	}

	fov := person.FieldOfView(3, blocked)

	wallStyle := style.Color(colornames.Bisque)
	fovStyle := style.Color(color.RGBA{R: 0xdd, G: 0xff, B: 0xdd, A: 0xff})
	personStyle := fovStyle.FontSize(40).Name("🧍")

	fig.AddCoords(grid)
	fig.AddStyledCoords(walls, wallStyle)
	fig.AddStyledCoords(fov, fovStyle)
	fig.AddStyledCoord(person, personStyle)

	_ = fig.RenderFile("../images/field_of_view.svg")
}
