package plot

import (
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
)

// Cell represents a hexagon to be drawn
type Cell struct {
	Coord coord.Coord
	Style *CellStyle
	Text  string
}

// Figure represents a collection of hex coordinates to be plotted
type Figure struct {
	orientation consts.Orientation
	coordType   consts.CoordType
	theme       *Theme
	cells       map[coord.Axial]*Cell
}

// CellStyleFunc is a function for styling a cell in the figure
type CellStyleFunc func(coord coord.Coord) (string, *CellStyle)

// NewFigure returns a new Figure instance
func NewFigure() *Figure {
	return &Figure{
		orientation: consts.PointyTop,
		coordType:   consts.Cube,
		theme:       DefaultTheme(),
		cells:       make(map[coord.Axial]*Cell),
	}
}

// SetCoordType sets the coordinate system type to use for rendering coordinate values
func (f *Figure) SetCoordType(coordType consts.CoordType) {
	f.coordType = coordType
}

// SetOrientation sets the orientation of the hexagons rendered by the figure
func (f *Figure) SetOrientation(orientation consts.Orientation) {
	f.orientation = orientation
}

// AddCoord adds an un-styled hex cell at the given coordinate
func (f *Figure) AddCoord(coord coord.Coord) {
	f.cells[coord.Axial()] = &Cell{
		Coord: coord,
		Style: nil,
	}
}

// AddCoords adds un-styled hex cells at the given coordinates
func (f *Figure) AddCoords(coords coord.Coords) {
	for _, c := range coords.Axials().ToSlice() {
		f.AddCoord(c.Convert(coords.Type()))
	}
}

// AddStyledCoord adds a styled hex cell at the given coordinate, using the given style
func (f *Figure) AddStyledCoord(coord coord.Coord, style Style) {
	text, cellStyle := style(coord)
	f.cells[coord.Axial()] = &Cell{
		Coord: coord,
		Style: cellStyle,
		Text:  text,
	}
}

// AddStyledCoords adds a styled hex cells at the given coordinates, using the given style
func (f *Figure) AddStyledCoords(coords coord.Coords, style Style) {
	for _, c := range coords.Axials().ToSlice() {
		f.AddStyledCoord(c.Convert(coords.Type()), style)
	}
}
