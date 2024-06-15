package hexe

import "github.com/legendary-code/hexe/pkg/hexe/coord"

type Item[T any, C coord.Coord] interface {
	Index() C
	Value() T
	DistanceTo(other Item[T, C]) int
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

func (i *item[T, C]) DistanceTo(other Item[T, C]) int {
	return i.index.Cube().DistanceTo(other.Index().Cube())
}
