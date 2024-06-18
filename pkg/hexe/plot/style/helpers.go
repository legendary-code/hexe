package style

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/plot"
	"image/color"
)

// Color creates a new cell style with the given color
func Color(color color.RGBA) plot.Style {
	return func(coord coord.Coord) (string, *plot.CellStyle) {
		return "", &plot.CellStyle{Color: color}
	}
}

// Name creates a new cell style with the given name
func Name(name string) plot.Style {
	return func(coord coord.Coord) (string, *plot.CellStyle) {
		return name, nil
	}
}

// FontSize creates a new cell style with the given font size
func FontSize(size int) plot.Style {
	return func(coord coord.Coord) (string, *plot.CellStyle) {
		return "", &plot.CellStyle{FontSize: size}
	}
}
