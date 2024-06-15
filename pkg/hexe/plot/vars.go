package plot

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/legendary-code/hexe/pkg/hexe/math"
	"golang.org/x/exp/maps"
	"image/color"
	"reflect"
)

const viewBoxPadding = 100

type svgViewBox struct {
	MinX   int
	MinY   int
	Width  int
	Height int
}

type svgCell struct {
	X         float64
	Y         float64
	R         string
	Q         string
	S         string
	Style     *svgStyle
	PointyTop bool
	Triplet   bool
	Text      string
}

type svgStyle struct {
	Id       string
	Color    color.RGBA
	FontSize int
}

type renderVars struct {
	Theme   *Theme
	Cells   []*svgCell
	Styles  []*svgStyle
	ViewBox *svgViewBox
}

type renderVarsBuilder struct {
	theme       *Theme
	orientation consts.Orientation
	coordType   consts.CoordType
	cells       map[coord.Axial]*Cell
	styles      map[CellStyle]int
}

func newRenderVarsBuilder(theme *Theme, orientation consts.Orientation, coordType consts.CoordType) *renderVarsBuilder {
	return &renderVarsBuilder{
		theme:       theme,
		orientation: orientation,
		coordType:   coordType,
		cells:       make(map[coord.Axial]*Cell),
		styles:      make(map[CellStyle]int),
	}
}

func (r *renderVarsBuilder) addStyleCount(style CellStyle, amount int) {
	count := r.styles[style]
	r.styles[style] = count + amount
}

func (r *renderVarsBuilder) addCell(cell *Cell) {
	index := cell.Coord.Axial()
	existing, ok := r.cells[index]

	if cell.Style != nil {
		r.addStyleCount(*cell.Style, 1)
	}

	if ok && existing.Style != nil {
		r.addStyleCount(*existing.Style, -1)
	}

	r.cells[index] = cell
}

func fmtInt(i int) string {
	if i > 0 {
		return fmt.Sprintf("+%d", i)
	}
	return fmt.Sprintf("%d", i)
}

func (r *renderVarsBuilder) build() *renderVars {
	styles := make(map[CellStyle]*svgStyle)
	cells := make([]*svgCell, 0)
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for style, count := range r.styles {
		if count == 0 {
			continue
		}

		existing, ok := styles[style]
		if !ok {
			existing = &svgStyle{
				Id:       fmt.Sprintf("style-%d", len(styles)+1),
				Color:    style.Color,
				FontSize: style.FontSize,
			}
			styles[style] = existing
		}
	}

	for index, cell := range r.cells {
		var cellStyle *svgStyle
		if cell.Style != nil {
			cellStyle = styles[*cell.Style]
		}

		x, y := math.AxialToPixel(index.Q(), index.R(), 100.0, r.orientation)

		if int(x) < minX {
			minX = int(x)
		}

		if int(x) > maxX {
			maxX = int(x)
		}

		if int(y) < minY {
			minY = int(y)
		}

		if int(y) > maxY {
			maxY = int(y)
		}

		c := &svgCell{
			X:         x,
			Y:         y,
			Style:     cellStyle,
			Triplet:   r.coordType == consts.Cube,
			PointyTop: r.orientation == consts.PointyTop,
			Text:      cell.Text,
		}

		cv := cell.Coord.Convert(r.coordType)

		switch ct := cv.(type) {
		case coord.QRS:
			if ct.Q() == 0 && ct.R() == 0 && ct.S() == 0 {
				c.Q = "q"
				c.R = "r"
				c.S = "s"
			} else {
				c.Q = fmtInt(ct.Q())
				c.R = fmtInt(ct.R())
				c.S = fmtInt(ct.S())
			}
		case coord.QR:
			if ct.Q() == 0 && ct.R() == 0 {
				c.Q = "q"
				c.R = "r"
			} else {
				c.Q = fmtInt(ct.Q())
				c.R = fmtInt(ct.R())
			}
		default:
			panic(fmt.Sprintf("invalid coord type: %s", reflect.TypeOf(ct).Name()))
		}

		cells = append(cells, c)
	}

	minX = minX - int(math.CalculateHorizontalSpacing[float64](r.orientation, viewBoxPadding))
	maxX = maxX + int(math.CalculateHorizontalSpacing[float64](r.orientation, viewBoxPadding))
	minY = minY - int(math.CalculateVerticalSpacing[float64](r.orientation, viewBoxPadding))
	maxY = maxY + int(math.CalculateVerticalSpacing[float64](r.orientation, viewBoxPadding))

	return &renderVars{
		Theme:  r.theme,
		Cells:  cells,
		Styles: maps.Values(styles),
		ViewBox: &svgViewBox{
			MinX:   minX,
			MinY:   minY,
			Width:  maxX - minX,
			Height: maxY - minY,
		},
	}
}
