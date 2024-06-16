package main

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
)

func instantiationExample() {
	// new axial coordinate (0, 1)
	a := coord.NewAxial(0, 1)

	// convert to cube coordinates (0, 1, -1)
	c := a.Cube()
	fmt.Println(c.Q(), c.R(), c.S())

	// zero value
	c = coord.ZeroCube()

	// accessing components
	fmt.Println(c.Q(), c.R(), c.S())
}
