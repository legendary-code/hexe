// Code generated by internal/hexe/gen/examples; DO NOT EDIT.

package main

//go:generate go run ../internal/hexe/gen/examples

var examples = []func(){fieldOfViewExample, findPathBfsExample, instantiationExample,

	mathFunctionsExample,

	plotExample, setsExample}

func main() {
	for _, example := range examples {
		example()
	}
}
