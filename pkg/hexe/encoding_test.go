package hexe

type testCodec struct {
}

func (e *testCodec) Encode(value string) ([]byte, error) {
	return []byte(value), nil
}

func (e *testCodec) Decode(value []byte) (string, error) {
	return string(value), nil
}
