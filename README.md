# Hexe

![](images/hexe.png)

An easy-to-use golang library for working with hexagonal grids.

## Visualization
To help visualize hex grids generated in code, simple plotting functionality
is provided for drawing hex grid coordinates and styling the cells.

#### Example code:
```go
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

	f.AddStyledCoords(
		grid.Coords(),
		func(coord coord.Coord) (string, *plot.CellStyle) { return "🌊", waterStyle },
	)

	f.AddStyledCoords(
		[]coord.Coord{
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

	_ = f.RenderFile("example.svg")
}
```

#### Output:
![Example](images/example.svg)
