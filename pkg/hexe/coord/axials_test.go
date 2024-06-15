package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAxials_Copy(t *testing.T) {
	expected := Axials{
		NewAxial(0, 0),
		NewAxial(0, 1),
		NewAxial(1, 0),
		NewAxial(1, 1),
	}

	actual := expected.Copy()
	actual[0] = NewAxial(2, 2)

	assert.NotEqual(t, expected, actual)
}

func TestAxials_Sorted(t *testing.T) {
	expected := Axials{
		NewAxial(0, 0),
		NewAxial(0, 1),
		NewAxial(1, 0),
		NewAxial(1, 1),
	}

	actual := Axials{
		NewAxial(1, 1),
		NewAxial(0, 0),
		NewAxial(1, 0),
		NewAxial(0, 1),
	}

	actual = actual.Sorted()
	assert.Equal(t, expected, actual)
}

func TestAxials_UnionWith(t *testing.T) {
	a := Axials{
		NewAxial(0, 0),
		NewAxial(1, 0),
	}

	b := Axials{
		NewAxial(0, 1),
		NewAxial(1, 1),
	}

	c := a.UnionWith(b)
	expected := Axials{
		NewAxial(0, 0),
		NewAxial(1, 0),
		NewAxial(0, 1),
		NewAxial(1, 1),
	}

	assert.Equal(t, expected, c)
}

func TestAxials_IntersectWith(t *testing.T) {
	a := Axials{
		NewAxial(0, 0),
		NewAxial(1, 0),
		NewAxial(0, 1),
		NewAxial(1, 1),
	}

	b := Axials{
		NewAxial(0, 1),
		NewAxial(1, 0),
	}

	c := a.IntersectWith(b).Sorted()
	expected := Axials{
		NewAxial(0, 1),
		NewAxial(1, 0),
	}

	assert.Equal(t, expected, c)
}

func TestAxials_DifferenceWith(t *testing.T) {
	a := Axials{
		NewAxial(0, 0),
		NewAxial(1, 0),
		NewAxial(0, 1),
		NewAxial(1, 1),
	}

	b := Axials{
		NewAxial(0, 1),
		NewAxial(1, 0),
	}

	c := a.DifferenceWith(b).Sorted()
	expected := Axials{
		NewAxial(0, 0),
		NewAxial(1, 1),
	}

	assert.Equal(t, expected, c)
}
