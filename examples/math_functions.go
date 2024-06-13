package main

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/math"
)

func mathFunctionsExample() {
	distance := math.CubeDistance(0, 1, -1, 0, 2, -2)
	fmt.Printf("The distance from (0, 1, -1) to (0, 2, -2) is %d\n", distance)
}
