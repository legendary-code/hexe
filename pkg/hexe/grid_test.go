package hexe

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGridGetSet(t *testing.T) {
	g := NewCubeGrid[string]()
	cc := coord.NewCube(1, 0, -1)

	s, ok := g.Index(cc)
	assert.False(t, ok)
	assert.Equal(t, "", s)

	g.Set(cc, "foo")
	s, ok = g.Index(cc)
	assert.True(t, ok)
	assert.Equal(t, "foo", s)

	s = g.Get(cc)
	assert.Equal(t, "foo", s)

	a := g.Axial()
	ac := coord.NewAxial(1, 0)
	s = a.Get(ac)
	assert.Equal(t, "foo", s)

	a.Delete(ac)
	s, ok = a.Index(ac)
	assert.False(t, ok)
	assert.Equal(t, "", s)
}
