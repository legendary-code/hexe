package hexe

import (
	"github.com/legendary-code/hexe/pkg/hexe/coord"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItems_Values(t *testing.T) {
	is := Items[string, coord.Axial, coord.Axials]{
		newItem(coord.NewAxial(0, 0), "foo"),
		newItem(coord.NewAxial(0, 1), "bar"),
		newItem(coord.NewAxial(1, 0), "baz"),
	}

	assert.Equal(t, []string{"foo", "bar", "baz"}, is.Values())
}

func TestItems_Indices(t *testing.T) {
	is := Items[string, coord.Axial, coord.Axials]{
		newItem(coord.NewAxial(0, 0), "foo"),
		newItem(coord.NewAxial(0, 1), "bar"),
		newItem(coord.NewAxial(1, 0), "baz"),
	}

	assert.Equal(t, coord.Axials{coord.NewAxial(0, 0), coord.NewAxial(0, 1), coord.NewAxial(1, 0)}, is.Indices())
}
