package main

import (
	_ "embed"
	"fmt"
	"github.com/legendary-code/hexe/internal/hexe/check"
	"github.com/legendary-code/hexe/internal/hexe/unsafe"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	exampleFunctions := findExampleFunctions()
	src, fset := generateCode(exampleFunctions)
	writeGeneratedCode(src, fset)
}

func findExampleFunctions() []string {
	exampleFunctions := make([]string, 0)

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, 0)
	check.Error(err)

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				funcDecl, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}

				if strings.HasSuffix(funcDecl.Name.Name, "Example") {
					exampleFunctions = append(exampleFunctions, funcDecl.Name.Name)
				}
			}
		}
	}

	return exampleFunctions
}

func generateCode(exampleFunctions []string) (ast.Node, *token.FileSet) {
	wd, err := os.Getwd()
	check.Error(err)

	templateFile := filepath.Join(wd, "..", "internal/hexe/gen/examples/main.go.tmpl")
	fset := token.NewFileSet()

	src, err := parser.ParseFile(fset, templateFile, nil, parser.ParseComments)
	check.Error(err)

	ast.Inspect(src, func(node ast.Node) bool {
		decl, ok := node.(*ast.ValueSpec)
		if !ok {
			return true
		}

		if len(decl.Names) != 1 || decl.Names[0].Name != "examples" {
			return true
		}

		decl.Values[0], err = parser.ParseExpr(fmt.Sprintf("[]func(){%s}", strings.Join(exampleFunctions, ",")))
		check.Error(err)

		return true
	})

	return src, fset
}

func writeGeneratedCode(src ast.Node, fset *token.FileSet) {
	file, err := os.Create("main.go")
	check.Error(err)
	defer unsafe.CloseIgnoreError(file)

	err = format.Node(file, fset, src)
	check.Error(err)
}
