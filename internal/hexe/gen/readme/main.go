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
		"example":         example,
		"logo":            logo,
		"topAnchor":       topAnchor,
		"backToTop":       backToTop,
		"tableOfContents": tableOfContents,
	})

	tmpl, err := tmpl.Parse(readMeTemplate)
	check.Error(err)

	file, err := os.Create("../../../../README.md")
	check.Error(err)
	defer unsafe.CloseIgnoreError(file)

	check.Error(tmpl.Execute(file, nil))
}

func getExampleCodeMarkdown(name string) string {
	wd, err := os.Getwd()
	check.Error(err)

	fileName := filepath.Join(wd, "../../../..", "examples", fmt.Sprintf("%s.go", name))
	file, err := os.Open(fileName)
	check.Error(err)
	defer unsafe.CloseIgnoreError(file)

	bytes, err := io.ReadAll(file)
	check.Error(err)

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("https://github.com/legendary-code/hexe/blob/main/examples/%s.go\n", name))
	sb.WriteString("```go\n")
	sb.WriteString(string(bytes))
	sb.WriteString("```\n")

	return sb.String()
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
	embeddedCode := getExampleCodeMarkdown(name)

	sb := strings.Builder{}
	sb.WriteString(embeddedCode)

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
	return `<p align="right">(<a href="#readme-top">back to top</a>)</p>`
}

func makeAnchorLink(text string) string {
	return fmt.Sprintf("#%s", strings.ToLower(strings.ReplaceAll(text, " ", "-")))
}

func tableOfContents() string {
	sb := strings.Builder{}
	sb.WriteString("<details>\n")
	sb.WriteString("\t<summary>Table of Contents</summary>\n")
	sb.WriteString("\t<ol>\n")

	level := 0

	readMeLines := strings.Split(readMeTemplate, "\n")

	for _, line := range readMeLines {
		if strings.HasPrefix(line, "## ") {
			headerText := strings.TrimPrefix(line, "## ")
			anchorLink := makeAnchorLink(headerText)

			if level > 2 {
				sb.WriteString("\t\t\t</ul>\n")
				sb.WriteString("\t\t</li>\n")
			} else if level == 2 {
				sb.WriteString("\t\t</li>\n")
			}

			sb.WriteString("\t\t<li>\n")
			sb.WriteString(fmt.Sprintf("\t\t\t<a href=\"%s\">%s</a>\n", anchorLink, headerText))

			level = 2
		} else if strings.HasPrefix(line, "### ") {
			headerText := strings.TrimPrefix(line, "### ")
			anchorLink := makeAnchorLink(headerText)

			if level < 3 {
				sb.WriteString("\t\t\t<ul>\n")
			}

			sb.WriteString(fmt.Sprintf("\t\t\t\t<li><a href=\"%s\">%s</a></li>\n", anchorLink, headerText))

			level = 3
		}
	}

	if level > 2 {
		sb.WriteString("\t\t\t</ul>\n")
		sb.WriteString("\t\t</li>\n")
	}

	if level > 1 {
		sb.WriteString("\t\t</li>\n")
	}

	sb.WriteString("\t</ol>\n")
	sb.WriteString("</details>\n")
	return sb.String()
}
