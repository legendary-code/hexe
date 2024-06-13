package main

import (
	"fmt"
	"github.com/legendary-code/hexe/pkg/hexe"
	"github.com/legendary-code/hexe/pkg/hexe/coord"
)

func gridExample() {
	grid := hexe.NewAxialGrid[string]()

	// set some values
	coords := coord.ZeroAxial().MovementRange(2)
	for i := coords.Iterator(); i.Next(); {
		c := i.Item()
		grid.Set(c, fmt.Sprintf("%d", c.Q()+c.R()))
	}

	// remove the center value
	grid.Delete(coord.ZeroAxial())

	// get values for a line
	line := coord.NewAxial(1, 1).LineTo(coord.NewAxial(-1, -1))
	values := grid.GetAll(line)

	// print it out
	for c, value := range values {
		fmt.Printf("%v => %s\n", c, value)
	}
}
