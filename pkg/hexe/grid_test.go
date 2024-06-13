package hexe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGridGetSet(t *testing.T) {
	g := NewGrid[string]()
	c := g.Cube()

	s, ok := c.Index(1, 0, -1)
	assert.False(t, ok)
	assert.Equal(t, "", s)

	c.Set(1, 0, -1, "foo")
	s, ok = c.Index(1, 0, -1)
	assert.True(t, ok)
	assert.Equal(t, "foo", s)

	s = c.Get(1, 0, -1)
	assert.Equal(t, "foo", s)

	a := c.Axial()
	s = a.Get(1, 0)
	assert.Equal(t, "foo", s)

	a.Delete(1, 0)
	s, ok = a.Index(1, 0)
	assert.False(t, ok)
	assert.Equal(t, "", s)
}
