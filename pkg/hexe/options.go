package hexe

type Option[T any] interface {
	apply(*configurableGrid[T])
}

type option[T any] struct {
	applyFunc func(*configurableGrid[T])
}

func (o *option[T]) apply(grid *configurableGrid[T]) {
	o.applyFunc(grid)
}

func WithEncoder[T any](encoder Encoder[T]) Option[T] {
	return &option[T]{
		applyFunc: func(grid *configurableGrid[T]) {
			grid.encoder = encoder
		},
	}
}

func WithDecoder[T any](decoder Decoder[T]) Option[T] {
	return &option[T]{
		applyFunc: func(grid *configurableGrid[T]) {
			grid.decoder = decoder
		},
	}
}
