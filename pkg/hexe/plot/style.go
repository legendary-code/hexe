package plot

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"image/color"
)

type Style func(coord.Coord) (string, *CellStyle)

type CellStyle struct {
	Color    color.RGBA
	FontSize int
}

func (s Style) Name(name string) Style {
	return func(c coord.Coord) (string, *CellStyle) {
		_, style := s(c)
		return name, style
	}
}

func (s Style) Color(color color.RGBA) Style {
	return func(c coord.Coord) (string, *CellStyle) {
		name, style := s(c)
		if style == nil {
			style = &CellStyle{}
		}
		style.Color = color
		return name, style

	}
}

func (s Style) FontSize(size int) Style {
	return func(c coord.Coord) (string, *CellStyle) {
		name, style := s(c)
		if style == nil {
			style = &CellStyle{}
		}
		style.FontSize = size
		return name, style
	}
}
