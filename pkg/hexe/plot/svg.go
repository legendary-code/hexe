package plot

import (
	_ "embed"
	"fmt"
	"github.com/legendary-code/hexe/internal/hexe/unsafe"
	"html/template"
	"image/color"
	"io"
	"os"
)

//go:embed figure.svg.tmpl
var svgTemplate string

func (f *Figure) RenderFile(name string) error {
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

				return fmt.Sprintf(`class="%s"`, style.Id)
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
