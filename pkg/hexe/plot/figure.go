package plot

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"image/color"
)

type CellStyle struct {
	Color    color.RGBA
	FontSize int
}

type Cell struct {
	Coord coord.Coord
	Style *CellStyle
	Text  string
}

type Figure struct {
	orientation consts.Orientation
	coordType   consts.CoordType
	theme       *Theme
	cells       map[coord.Axial]*Cell
}

type CellStyleFunc func(coord coord.Coord) (string, *CellStyle)

func NewFigure() *Figure {
	return &Figure{
		orientation: consts.PointyTop,
		coordType:   consts.Cube,
		theme:       DefaultTheme(),
		cells:       make(map[coord.Axial]*Cell),
	}
}

func (f *Figure) SetCoordType(coordType consts.CoordType) {
	f.coordType = coordType
}

func (f *Figure) AddCoord(coord coord.Coord) {
	f.cells[coord.Axial()] = &Cell{
		Coord: coord,
		Style: nil,
	}
}

func (f *Figure) AddCoords(coords coord.Coords) {
	for _, c := range coords.ToSlice() {
		f.AddCoord(c)
	}
}

func (f *Figure) AddStyledCoord(coord coord.Coord, styleFunc CellStyleFunc) {
	text, style := styleFunc(coord)
	f.cells[coord.Axial()] = &Cell{
		Coord: coord,
		Style: style,
		Text:  text,
	}
}

func (f *Figure) AddStyledCoords(coords coord.Coords, styleFunc CellStyleFunc) {
	for _, c := range coords.ToSlice() {
		f.AddStyledCoord(c, styleFunc)
	}
}
