package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/legendary-code/hexe/internal/hexe/check"
	"github.com/legendary-code/hexe/internal/hexe/unsafe"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:generate go run .
//go:embed README.md.tmpl
var readMeTemplate string

func main() {
	fmt.Println("Processing README.md.tmpl...")

	tmpl := template.New("")
	tmpl.Funcs(map[string]any{
		"example":   example,
		"logo":      logo,
		"topAnchor": topAnchor,
		"backToTop": backToTop,
	})

	tmpl, err := tmpl.Parse(readMeTemplate)
	check.Error(err)

	file, err := os.Create("../../../../README.md")
	check.Error(err)
	defer unsafe.CloseIgnoreError(file)

	check.Error(tmpl.Execute(file, nil))
}

func countExampleCodeLines(name string) int {
	wd, err := os.Getwd()
	check.Error(err)

	fileName := filepath.Join(wd, "../../../..", "examples", fmt.Sprintf("%s.go", name))
	file, err := os.Open(fileName)
	check.Error(err)
	defer unsafe.CloseIgnoreError(file)

	bytes, err := io.ReadAll(file)
	check.Error(err)

	lineCount := 1
	for _, b := range bytes {
		if b == '\n' {
			lineCount++
		}
	}

	return lineCount
}

func hasExampleImage(name string) bool {
	wd, err := os.Getwd()
	check.Error(err)

	fileName := filepath.Join(wd, "../../../..", "images", fmt.Sprintf("%s.svg", name))
	fi, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return !fi.IsDir()
}

func example(name string) string {
	sb := strings.Builder{}
	numLines := countExampleCodeLines(name)

	sb.WriteString(
		fmt.Sprintf(
			"https://github.com/legendary-code/hexe/blob/master/examples/%s.go#L1-L%d\n\n",
			name,
			numLines,
		),
	)

	if hasExampleImage(name) {
		sb.WriteString(
			fmt.Sprintf("#### Output:\n![Example](images/%s.svg)\n\n", name),
		)
	}

	return sb.String()
}

func logo() string {
	return `<img src="images/hexe.png" alt="Logo">`
}

func topAnchor() string {
	return `<a name="readme-top"></a>`
}

func backToTop() string {
	return `<p style="text-align: right">(<a href="#readme-top">back to top</a>)</p>`
}
