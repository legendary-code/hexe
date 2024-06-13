package plot

import (
	_ "embed"
	"fmt"
	"github.com/legendary-code/hexe/internal/hexe/unsafe"
	"html/template"
	"image/color"
	"io"
	"os"
	"path/filepath"
)

//go:embed figure.svg.tmpl
var svgTemplate string

// RenderFile renders the figure to the specified SVG file
func (f *Figure) RenderFile(name string) error {
	if !filepath.IsAbs(name) {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		name = filepath.Join(wd, name)
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer unsafe.CloseIgnoreError(file)
	return f.Render(file)
}

func (f *Figure) buildRenderVars() *renderVars {
	b := newRenderVarsBuilder(f.theme, f.orientation, f.coordType)
	for _, cell := range f.cells {
		b.addCell(cell)
	}
	return b.build()
}

// Render renders the figure as an SVG to the specified writer
func (f *Figure) Render(writer io.Writer) error {
	tmpl := template.New("").Funcs(
		map[string]any{
			"hex": func(color color.RGBA) string {
				return fmt.Sprintf("#%02x%02x%02x%02x", color.R, color.G, color.B, color.A)
			},
			"renderClass": func(style *svgStyle) string {
				if style == nil {
					return ""
				}

				return fmt.Sprintf(`class="%s"`, style.Class)
			},
		},
	)

	tmpl, err := tmpl.Parse(svgTemplate)
	if err != nil {
		return err
	}

	v := f.buildRenderVars()
	return tmpl.Execute(writer, v)
}
