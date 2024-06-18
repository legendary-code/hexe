package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type GridIterator[T any, C coord.CCoord, CS coord.CCoords] interface {
	Next() bool
	Item() T
	Index() C
}

type gridIterator[T any, C coord.CCoord, CS coord.CCoords] struct {
	grid        *grid[T, C, CS]
	convertFunc func(axial coord.Axial) C
	indices     []coord.Axial
	index       int
}

func (o *gridIterator[T, C, CS]) Next() bool {
	if o.index >= len(o.indices)-1 {
		return false
	}

	o.index++
	return true
}

func (o *gridIterator[T, C, CS]) Item() T {
	return o.grid.get(o.Index())
}

func (o *gridIterator[T, C, CS]) Index() C {
	return o.convertFunc(o.indices[o.index])
}
