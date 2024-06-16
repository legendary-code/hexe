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
	"github.com/legendary-code/hexe/pkg/hexe/plot/style"
	"golang.org/x/image/colornames"
)

func main() {
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

	_ = fig.RenderFile("example.svg")
}
```

#### Output:
![Example](images/plot.svg)
