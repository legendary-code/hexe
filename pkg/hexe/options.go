package hexe

// Option represents an option for configuring a grid
type Option[T any] interface {
	apply(*configurableGrid[T])
}

type option[T any] struct {
	applyFunc func(*configurableGrid[T])
}

func (o *option[T]) apply(grid *configurableGrid[T]) {
	o.applyFunc(grid)
}

// WithEncoder returns an Option[T] for configuring a value encoder for grid persistence
func WithEncoder[T any](encoder Encoder[T]) Option[T] {
	return &option[T]{
		applyFunc: func(grid *configurableGrid[T]) {
			grid.encoder = encoder
		},
	}
}

// WithDecoder returns an Option[T] for configuring a value decoder for grid persistence
func WithDecoder[T any](decoder Decoder[T]) Option[T] {
	return &option[T]{
		applyFunc: func(grid *configurableGrid[T]) {
			grid.decoder = decoder
		},
	}
}

// WithEncoderDecoder returns an Option[T] for configuring a value encoder and decoder for grid persistence
func WithEncoderDecoder[T any](codec EncoderDecoder[T]) Option[T] {
	return &option[T]{
		applyFunc: func(grid *configurableGrid[T]) {
			grid.encoder = codec
			grid.decoder = codec
		},
	}
}
