package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type Item[T any, C coord.Coord] interface {
	Index() C
	Value() T
}

type item[T any, C coord.Coord] struct {
	index C
	value T
}

func newItem[T any, C coord.Coord](index C, value T) Item[T, C] {
	return &item[T, C]{
		index: index,
		value: value,
	}
}

func (i *item[T, C]) Index() C {
	return i.index
}

func (i *item[T, C]) Value() T {
	return i.value
}
