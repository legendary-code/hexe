package hexe

type Option[T any] interface {
	apply(*grid[T])
}

type option[T any] struct {
	applyFunc func(*grid[T])
}

func (o *option[T]) apply(grid *grid[T]) {
	o.applyFunc(grid)
}

func WithEncoder[T any](encoder Encoder[T]) Option[T] {
	return &option[T]{
		applyFunc: func(grid *grid[T]) {
			grid.encoder = encoder
		},
	}
}

func WithDecoder[T any](decoder Decoder[T]) Option[T] {
	return &option[T]{
		applyFunc: func(grid *grid[T]) {
			grid.decoder = decoder
		},
	}
}
