package hexe

type Encoder[T any] interface {
	Encode(value T) ([]byte, error)
}

type Decoder[T any] interface {
	Decode([]byte) (T, error)
}
