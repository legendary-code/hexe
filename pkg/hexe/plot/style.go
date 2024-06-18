package plot

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"image/color"
)

// Style represents text and cell style
type Style func(coord.Coord) (string, *CellStyle)

// CellStyle represents style for a cell
type CellStyle struct {
	Color    color.RGBA
	FontSize int
}

// Name sets the name of the cell style
func (s Style) Name(name string) Style {
	return func(c coord.Coord) (string, *CellStyle) {
		_, style := s(c)
		return name, style
	}
}

// Color sets the color of the cell style
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

// FontSize sets the font size of the cell style
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
