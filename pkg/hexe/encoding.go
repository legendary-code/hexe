package hexe

// Encoder represents a value encoder for encoding grid values
type Encoder[T any] interface {
	Encode(value T) ([]byte, error)
}

// Decoder represents a value decoder for decoding grid values
type Decoder[T any] interface {
	Decode([]byte) (T, error)
}

// EncoderDecoder represents a value codec for encoding and decoding grid values
type EncoderDecoder[T any] interface {
	Encoder[T]
	Decoder[T]
}
