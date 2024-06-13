package hexe

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGrid_EncodeDecode(t *testing.T) {
	codec := &testCodec{}
	g := NewAxialGrid[string](
		WithEncoderDecoder[string](codec),
	)

	g.Set(coord.NewAxial(0, 1), "foo")
	g.Set(coord.NewAxial(1, 0), "bar")

	sb := strings.Builder{}

	err := g.Encode(&sb)
	if err != nil {
		panic(err)
	}

	g.Clear()

	r := strings.NewReader(sb.String())
	err = g.Decode(r)
	if err != nil {
		panic(err)
	}

	v1 := g.Get(coord.NewAxial(0, 1))
	v2 := g.Get(coord.NewAxial(1, 0))

	assert.Equal(t, "foo", v1)
	assert.Equal(t, "bar", v2)
}
