package hexe

type Encoder[T any] interface {
	Encode(value T) []byte
}

type Decoder[T any] interface {
	Decode([]byte) T
}
