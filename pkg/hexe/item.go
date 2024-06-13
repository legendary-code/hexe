package hexe

type Item[T any, C any] interface {
	Index() C
	Value() T
}

type item[T any, C any] struct {
	index C
	value T
}

func newItem[T any, C any](index C, value T) Item[T, C] {
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
