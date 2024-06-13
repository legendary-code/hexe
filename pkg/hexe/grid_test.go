package hexe

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid_GetSet(t *testing.T) {
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
}

func TestGrid_Convert(t *testing.T) {
	cg := NewCubeGrid[string]()
	cc := coord.NewCube(1, 0, -1)

	cg.Set(cc, "foo")

	dwg := cg.DoubleWidth()
	assert.Equal(t, "foo", dwg.Get(cc.DoubleWidth()))
}

func TestGrid_Delete(t *testing.T) {
	g := NewCubeGrid[string]()
	cc := coord.NewCube(1, 0, -1)

	g.Set(cc, "foo")
	s := g.Get(cc)
	assert.Equal(t, "foo", s)

	g.Delete(cc)
	s = g.Get(cc)
	assert.Equal(t, "", s)
}

func TestGrid_GetAll(t *testing.T) {
	g := NewCubeGrid[string]()

	cc1 := coord.NewCube(1, 0, -1)
	cc2 := coord.NewCube(1, 1, -2)
	cc3 := coord.NewCube(1, 2, -3)

	ccs := coord.NewCubes(
		cc1,
		cc2,
		cc3,
	)

	g.Set(cc1, "foo")
	g.Set(cc2, "bar")

	ss := g.GetAll(ccs)

	v1 := ss[cc1]
	v2 := ss[cc2]
	v3 := ss[cc3]

	assert.Equal(t, "foo", v1)
	assert.Equal(t, "bar", v2)
	assert.Equal(t, "", v3)
}
