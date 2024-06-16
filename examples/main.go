// Code generated by internal/hexe/gen/examples; DO NOT EDIT.

package main

//go:generate go run ../internal/hexe/gen/examples

var examples = []func(){neighborsExample, fieldOfViewExample, findPathBfsExample,

	mathFunctionsExample,

	traceToExample, diagonalNeighborsExample, plotExample,

	setsExample, rotateExample,

	instantiationExample,
	lineToExample, reflectExample,
	ringExample, floodFillExample, movementRangeExample}

func main() {
	for _, example := range examples {
		example()
	}
}
