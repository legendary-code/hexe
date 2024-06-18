// Code generated by internal/hexe/gen/examples; DO NOT EDIT.

package main

//go:generate go run ../internal/hexe/gen/examples

var examples = []func(){diagonalNeighborsExample, fieldOfViewExample, findPathBfsExample,

	floodFillExample,

	gridExample, gridPersistenceExample, instantiationExample,

	lineToExample,

	mathFunctionsExample,
	movementRangeExample, neighborsExample,

	plotExample, reflectExample, ringExample, rotateExample, setsExample, traceToExample}

func main() {
	for _, example := range examples {
		example()
	}
}
