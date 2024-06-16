package main

import (
	"fmt"
	"github.com/legendary-code/hexe/internal/hexe/check"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"os"
	"path/filepath"
)

type ExampleFunc func() *plot.Figure

var examples = make(map[string]ExampleFunc)

func main() {
	wd, err := os.Getwd()
	check.Error(err)

	for name, exampleFunc := range examples {
		outSvg := filepath.Join(wd, "images", fmt.Sprintf("%s.svg", name))
		fig := exampleFunc()
		err = fig.RenderFile(outSvg)
		check.Error(err)
	}
}

func addExample(name string, example ExampleFunc) {
	examples[name] = example
}

func createArena() (coord.Axials, coord.Axials) {
	return coord.ZeroAxial().MovementRange(8),
		coord.ZeroAxial().Ring(8).UnionWith(
			coord.Axials{
				coord.NewAxial(-5, -2),
				coord.NewAxial(-4, -2),
				coord.NewAxial(-4, -1),
				coord.NewAxial(-4, 0),
				coord.NewAxial(-4, 1),
				coord.NewAxial(-3, -1),
				coord.NewAxial(-6, 7),
				coord.NewAxial(-5, 6),
				coord.NewAxial(-4, 5),
				coord.NewAxial(-3, 5),
				coord.NewAxial(-3, 6),
				coord.NewAxial(-3, 7),
				coord.NewAxial(-4, 4),
				coord.NewAxial(-3, 3),
				coord.NewAxial(0, -2),
				coord.NewAxial(1, -2),
				coord.NewAxial(1, -1),
				coord.NewAxial(2, -1),
				coord.NewAxial(3, -1),
				coord.NewAxial(3, 0),
				coord.NewAxial(2, 1),
				coord.NewAxial(1, 2),
				coord.NewAxial(4, -2),
				coord.NewAxial(5, -3),
			},
		)
}
