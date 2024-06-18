package hexe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGridOptions(t *testing.T) {
	ag := NewAxialGrid[string](
		WithEncoder[string](&testCodec{}),
		WithDecoder[string](&testCodec{}),
	)
	a, ok := ag.(*axialGrid[string])
	if !ok {
		assert.Fail(t, "failed to cast")
	}

	assert.NotNil(t, a.encoder)
	assert.NotNil(t, a.decoder)
}
