package hexe

type testEncoder struct {
}

func (e *testEncoder) Encode(value string) []byte {
	return []byte(value)
}

type testDecoder struct {
}

func (e *testDecoder) Decode(value []byte) string {
	return string(value)
}
