package main

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func init() {
	addExample("find_path_bfs", func() *plot.Figure {
		fig := plot.NewFigure()
		grid, walls := createArena()

		person := coord.NewAxial(-6, 1)
		target := coord.NewAxial(6, -3)

		blocked := func(coord coord.Axial) bool {
			for _, wall := range walls {
				if wall == coord {
					return true
				}
			}

			return false
		}

		path := person.FindPathBFS(target, 20, blocked)

		wallStyle := style.Color(colornames.Bisque)
		pathStyle := style.Color(colornames.Lightblue).FontSize(40)
		personStyle := pathStyle.Name("🧍")
		targetStyle := pathStyle.Name("❌")

		fig.AddCoords(grid)
		fig.AddStyledCoords(walls, wallStyle)
		fig.AddStyledCoords(path, pathStyle)
		fig.AddStyledCoord(person, personStyle)
		fig.AddStyledCoord(target, targetStyle)

		return fig
	})
}
