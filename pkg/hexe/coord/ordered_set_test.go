package coord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedSet_Add(t *testing.T) {
	s := newOrderedSet[int]()

	assert.True(t, s.Add(0))
	assert.True(t, s.Add(1))
	assert.True(t, s.Add(2))
	assert.False(t, s.Add(0))

	assert.Equal(t, []int{0, 1, 2}, s.ToSlice())
}

func TestOrderedSet_Remove(t *testing.T) {
	s := newOrderedSet[int](0, 1, 2)

	assert.False(t, s.Remove(3))
	assert.Equal(t, []int{0, 1, 2}, s.ToSlice())

	assert.True(t, s.Remove(0))
	assert.Equal(t, []int{1, 2}, s.ToSlice())

	assert.True(t, s.Remove(2))
	assert.Equal(t, []int{1}, s.ToSlice())

	assert.True(t, s.Remove(1))
	assert.Equal(t, []int{}, s.ToSlice())
}

func TestOrderedSet_Contains(t *testing.T) {
	s := newOrderedSet[int](0, 1, 2)

	assert.True(t, s.Contains(0))
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(2))
	assert.False(t, s.Contains(3))
}

func TestOrderedSet_IsEmpty(t *testing.T) {
	s := newOrderedSet[int]()

	assert.True(t, s.IsEmpty())

	s.Add(0)
	assert.False(t, s.IsEmpty())

	s.Remove(0)
	assert.True(t, s.IsEmpty())
}

func TestOrderedSet_Clear(t *testing.T) {
	s := newOrderedSet[int]()

	assert.False(t, s.Clear())
	assert.True(t, s.IsEmpty())

	s.Add(0)
	assert.True(t, s.Clear())
	assert.True(t, s.IsEmpty())
	assert.Equal(t, []int{}, s.ToSlice())
}

func TestOrderedSet_Iterator(t *testing.T) {
	s := newOrderedSet[int](0, 1, 2)

	v := make([]int, 0)
	for i := s.Iterator(); i.Next(); {
		v = append(v, i.Item())
	}

	assert.Equal(t, []int{0, 1, 2}, v)
}

func TestOrderedSet_ForEach(t *testing.T) {
	s := newOrderedSet[int](0, 2, 4, 5, 6)

	v := make([]int, 0)
	s.ForEach(func(i int) bool {
		v = append(v, i)
		return true
	})
	assert.Equal(t, []int{0, 2, 4, 5, 6}, v)

	v = make([]int, 0)
	s.ForEach(func(i int) bool {
		if i%2 != 0 {
			return false
		}
		v = append(v, i)
		return true
	})
	assert.Equal(t, []int{0, 2, 4}, v)
}

func TestOrderedSet_Union(t *testing.T) {
	s1 := newOrderedSet[int](0, 1, 2, 3)
	s2 := newOrderedSet[int](2, 3, 4, 5)
	s3 := s1.Union(s2)

	assert.Equal(t, []int{0, 1, 2, 3}, s1.ToSlice())
	assert.Equal(t, []int{2, 3, 4, 5}, s2.ToSlice())
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, s3.ToSlice())
}

func TestOrderedSet_Intersect(t *testing.T) {
	s1 := newOrderedSet[int](0, 1, 2, 3)
	s2 := newOrderedSet[int](5, 4, 3, 2)
	s3 := s1.Intersect(s2)

	assert.Equal(t, []int{0, 1, 2, 3}, s1.ToSlice())
	assert.Equal(t, []int{5, 4, 3, 2}, s2.ToSlice())
	assert.Equal(t, []int{2, 3}, s3.ToSlice())
}

func TestOrderedSet_Difference(t *testing.T) {
	s1 := newOrderedSet[int](0, 1, 2, 3)
	s2 := newOrderedSet[int](5, 4, 3, 2)
	s3 := s1.Difference(s2)

	assert.Equal(t, []int{0, 1, 2, 3}, s1.ToSlice())
	assert.Equal(t, []int{5, 4, 3, 2}, s2.ToSlice())
	assert.Equal(t, []int{0, 1}, s3.ToSlice())
}

func TestOrderedSet_Equal(t *testing.T) {
	s1 := newOrderedSet[int](0, 1, 2)
	s2 := newOrderedSet[int](0, 1, 3)

	assert.Equal(t, s1, s1)
	assert.NotEqual(t, s1, s2)

	s2.Remove(3)
	s2.Add(2)
	assert.Equal(t, s1, s2)

	s1 = newOrderedSet[int](0, 1, 2)
	s2 = newOrderedSet[int](2, 1, 0)
	assert.NotEqual(t, s1, s2)
}

func TestOrderedSet_SetEqual(t *testing.T) {
	s1 := newOrderedSet[int](0, 1, 2)
	s2 := newOrderedSet[int](0, 1, 3)

	assert.True(t, s1.SetEqual(s1))
	assert.False(t, s1.SetEqual(s2))

	s2.Remove(3)
	s2.Add(2)
	assert.True(t, s1.SetEqual(s2))

	s1 = newOrderedSet[int](0, 1, 2)
	s2 = newOrderedSet[int](2, 1, 0)
	assert.True(t, s1.SetEqual(s2))
}
