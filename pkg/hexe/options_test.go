package hexe

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGridOptions(t *testing.T) {
	g := NewGrid[string](
		WithEncoder[string](&testEncoder{}),
		WithDecoder[string](&testDecoder{}),
	)

	if gg, ok := g.(*qrGrid[string, coord.Axial]); !ok {
		assert.Fail(t, "failed to cast")
	} else {
		assert.NotNil(t, gg.encoder)
		assert.NotNil(t, gg.decoder)
	}
}
