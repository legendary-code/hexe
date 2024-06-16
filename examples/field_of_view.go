package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
	"image/color"
)

func init() {
	addExample("field_of_view", func() *plot.Figure {
		fig := plot.NewFigure()
		grid, walls := createArena()

		person1 := coord.NewAxial(-1, -3)
		person2 := coord.NewAxial(-1, 3)

		blocked := func(coord coord.Axial) bool {
			for _, wall := range walls {
				if wall == coord {
					return true
				}
			}

			return false
		}

		fov1 := person1.FieldOfView(4, blocked)
		fov2 := person2.FieldOfView(4, blocked)
		fov3 := fov1.IntersectWith(fov2)
		fov1 = fov1.DifferenceWith(fov3)
		fov2 = fov2.DifferenceWith(fov3)

		lightGreen := color.RGBA{R: 0xdd, G: 0xff, B: 0xdd, A: 0xff}
		lightBlue := color.RGBA{R: 0xdd, G: 0xdd, B: 0xff, A: 0xff}
		lightPurple := color.RGBA{R: 0xdd, G: 0xee, B: 0xee, A: 0xff}

		wallStyle := style.Color(colornames.Bisque)
		person1Style := style.Color(lightGreen).FontSize(40).Name("🧍")
		person2Style := style.Color(lightBlue).FontSize(40).Name("🧍")
		fov1Style := style.Color(lightGreen)
		fov2Style := style.Color(lightBlue)
		fov3Style := style.Color(lightPurple)

		fig.AddCoords(grid)
		fig.AddStyledCoords(walls, wallStyle)
		fig.AddStyledCoords(fov1, fov1Style)
		fig.AddStyledCoords(fov2, fov2Style)
		fig.AddStyledCoords(fov3, fov3Style)
		fig.AddStyledCoord(person1, person1Style)
		fig.AddStyledCoord(person2, person2Style)

		return fig
	})
}
