package plot

import (
	"golang.org/x/image/colornames"
	"image/color"
)

// DefaultTheme returns the default theme for styling figures
func DefaultTheme() *Theme {
	return &Theme{
		BackgroundColor: colornames.White,
		BorderColor:     color.RGBA{R: 0xb3, G: 0xb3, B: 0xb3, A: 0xff},
		CellColor:       color.RGBA{R: 0xe4, G: 0xe4, B: 0xe1, A: 0xff},
		CoordQColor:     color.RGBA{R: 0x59, G: 0xb2, B: 0x01, A: 0xff},
		CoordRColor:     color.RGBA{R: 0x00, G: 0x99, B: 0xe5, A: 0xff},
		CoordSColor:     color.RGBA{R: 0xe6, G: 0x19, B: 0xe5, A: 0xff},
		TextColor:       colornames.Black,
	}
}

// Theme represents a theme for coloring a figure and its cells
type Theme struct {
	BackgroundColor color.RGBA
	BorderColor     color.RGBA
	CellColor       color.RGBA
	CoordQColor     color.RGBA
	CoordRColor     color.RGBA
	CoordSColor     color.RGBA
	TextColor       color.RGBA
}
