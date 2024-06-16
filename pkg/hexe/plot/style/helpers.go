package style

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"image/color"
)

func Color(color color.RGBA) plot.Style {
	return func(coord coord.Coord) (string, *plot.CellStyle) {
		return "", &plot.CellStyle{Color: color}
	}
}

func Name(name string) plot.Style {
	return func(coord coord.Coord) (string, *plot.CellStyle) {
		return "", nil
	}
}

func FontSize(size int) plot.Style {
	return func(coord coord.Coord) (string, *plot.CellStyle) {
		return "", &plot.CellStyle{FontSize: size}
	}
}
