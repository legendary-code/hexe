package main

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
)

func setsExample() {
	// Create a set of axial coordinates
	a := coord.Axials{
		coord.NewAxial(0, 0),
		coord.NewAxial(0, 1),
		coord.NewAxial(1, 0),
		coord.NewAxial(1, 1),
	}

	// Convert them to cube coordinates
	c := a.Cubes()

	// You can iterate over them
	for _, v := range c {
		fmt.Println(v)
	}
}
