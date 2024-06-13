package main

import (
	"fmt"
	"github.com/legendary-code/hexe/internal/hexe/check"
	"github.com/legendary-code/hexe/pkg/hexe/consts"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const GCoord = "GCoord"
const GReceiver = "g"
const GCoordLower = "gCoord"

func main() {
	wd, err := os.Getwd()
	check.Error(err)

	templates, err := filepath.Glob("*.go.tmpl")
	check.Error(err)

	for _, name := range templates {
		templateFile := filepath.Join(wd, name)
		processTemplate(templateFile)
	}
}

func getOutputFile(templateFile string) string {
	base := filepath.Dir(templateFile)
	fileName := strings.TrimSuffix(filepath.Base(templateFile), ".go.tmpl")
	return filepath.Join(base, fmt.Sprintf("%s_gen.go", fileName))
}

func parseAst(templateFile string) *ast.File {
	file, err := parser.ParseFile(token.NewFileSet(), templateFile, nil, parser.ParseComments)
	check.Error(err)
	return file
}

func processTemplate(templateFile string) {
	fmt.Printf("Processing %s...\n", filepath.Base(templateFile))

	sb := strings.Builder{}

	fset := token.NewFileSet()
	file := parseAst(templateFile)
	sb.WriteString("// Code generated by internal/hexe/gen/coords; DO NOT EDIT.\n")
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("package %s\n\n", file.Name.Name))

	ast.Inspect(file, func(node ast.Node) bool {
		importDecl, ok := node.(*ast.GenDecl)
		if ok && importDecl.Tok == token.IMPORT {
			err := format.Node(&sb, fset, importDecl)
			check.Error(err)
			sb.WriteString("\n")
			return true
		}

		return true
	})

	for _, coordType := range consts.CoordinateTypes() {
		if coordType == consts.Cube {
			// Cube is our exemplar implementation that all generated code is based off of
			continue
		}

		file = parseAst(templateFile)
		coordTypeName := coordType.Name()
		coordLowerCaseTypeName := strings.ToLower(coordTypeName[:1]) + coordTypeName[1:]
		coordsTypeReceiver := strings.ToLower(string(coordType.Name()[0]))

		// Replace placeholder types/receivers
		ast.Inspect(file, func(node ast.Node) bool {
			cmt, ok := node.(*ast.Comment)
			if !ok {
				return true
			}

			cmt.Text = strings.ReplaceAll(cmt.Text, GCoord, coordTypeName)
			cmt.Text = strings.ReplaceAll(cmt.Text, GCoordLower, coordLowerCaseTypeName)
			return true
		})

		ast.Inspect(file, func(node ast.Node) bool {
			ident, ok := node.(*ast.Ident)
			if !ok {
				return true
			}

			if ident.Name == GReceiver {
				ident.Name = coordsTypeReceiver
			} else if strings.Contains(ident.Name, GCoord) {
				ident.Name = strings.ReplaceAll(ident.Name, GCoord, coordTypeName)
			} else if strings.Contains(ident.Name, GCoordLower) {
				ident.Name = strings.ReplaceAll(ident.Name, GCoordLower, coordLowerCaseTypeName)
			}

			return true
		})

		ast.Inspect(file, func(node ast.Node) bool {
			lit, ok := node.(*ast.BasicLit)
			if !ok {
				return true
			}

			if lit.Kind == token.STRING {
				lit.Value = strings.ReplaceAll(lit.Value, GCoord, coordTypeName)
			}
			return true
		})

		// Write to file
		ast.Inspect(file, func(node ast.Node) bool {
			funcDecl, ok := node.(*ast.FuncDecl)
			if ok {
				err := format.Node(&sb, fset, funcDecl)
				check.Error(err)
				sb.WriteString("\n")
				return true
			}

			typeDecl, ok := node.(*ast.GenDecl)
			if ok && typeDecl.Tok == token.TYPE {
				err := format.Node(&sb, fset, typeDecl)
				check.Error(err)
				sb.WriteString("\n")
				return true
			}

			return true
		})
	}

	err := os.WriteFile(getOutputFile(templateFile), []byte(sb.String()), 0644)
	check.Error(err)
}
