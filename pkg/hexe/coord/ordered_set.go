package coord

type OrderedSetIterator[T comparable] interface {
	Next() bool
	Item() T
}

// OrderedSet represents an ordered set of coordinates
type OrderedSet[T comparable] interface {
	Add(T) bool
	Remove(T) bool
	Contains(T) bool
	Clear() bool
	Copy() OrderedSet[T]
	IsEmpty() bool
	Size() int
	Iterator() OrderedSetIterator[T]
	ForEach(func(T) bool)
	ToSlice() []T
	Union(OrderedSet[T]) OrderedSet[T]
	Intersect(OrderedSet[T]) OrderedSet[T]
	Difference(OrderedSet[T]) OrderedSet[T]
	Equal(OrderedSet[T]) bool
	SetEqual(OrderedSet[T]) bool
}

type item[T comparable] struct {
	value T
	next  *item[T]
	prev  *item[T]
}

type orderedSet[T comparable] struct {
	head  *item[T]
	tail  *item[T]
	items map[T]*item[T]
	size  int
}

type orderedSetIterator[T comparable] struct {
	item *item[T]
	next *item[T]
}

func newOrderedSet[T comparable](values ...T) *orderedSet[T] {
	s := &orderedSet[T]{
		head:  nil,
		tail:  nil,
		items: make(map[T]*item[T]),
		size:  0,
	}
	for _, value := range values {
		s.Add(value)
	}
	return s
}

func (o *orderedSet[T]) Remove(value T) bool {
	i, ok := o.items[value]
	if !ok {
		return false
	}

	if i.prev != nil {
		i.prev.next = i.next
	} else {
		o.head = i.next
	}

	if i.next != nil {
		i.next.prev = i.prev
	} else {
		o.tail = i.prev
	}

	delete(o.items, value)
	o.size--

	return true
}

func (o *orderedSet[T]) Add(value T) bool {
	if o.Contains(value) {
		return false
	}
	i := &item[T]{value: value}

	if o.tail == nil {
		o.head = i
		o.tail = i
	} else {
		i.prev = o.tail
		o.tail.next = i
		o.tail = i
	}

	o.items[value] = i
	o.size++

	return true
}

func (o *orderedSet[T]) ToSlice() []T {
	items := make([]T, o.size)
	n := 0
	for i := o.head; i != nil; i = i.next {
		items[n] = i.value
		n++
	}
	return items
}

func (o *orderedSet[T]) Contains(value T) bool {
	_, ok := o.items[value]
	return ok
}

func (o *orderedSet[T]) Clear() bool {
	if o.IsEmpty() {
		return false
	}

	o.items = make(map[T]*item[T])
	o.head = nil
	o.tail = nil
	o.size = 0

	return true
}

func (o *orderedSet[T]) IsEmpty() bool {
	return o.size == 0
}

func (o *orderedSet[T]) Iterator() OrderedSetIterator[T] {
	return &orderedSetIterator[T]{next: o.head}
}

func (o *orderedSet[T]) ForEach(iterFunc func(T) bool) {
	for i := o.Iterator(); i.Next() && iterFunc(i.Item()); {
	}
}

func (o *orderedSet[T]) Size() int {
	return o.size
}

func (o *orderedSet[T]) Copy() OrderedSet[T] {
	s := newOrderedSet[T]()
	for i := o.Iterator(); i.Next(); {
		s.Add(i.Item())
	}
	return s
}

func (o *orderedSet[T]) Union(set OrderedSet[T]) OrderedSet[T] {
	s := o.Copy()
	for i := set.Iterator(); i.Next(); {
		s.Add(i.Item())
	}
	return s
}

func (o *orderedSet[T]) Intersect(set OrderedSet[T]) OrderedSet[T] {
	s := newOrderedSet[T]()
	for i := o.Iterator(); i.Next(); {
		if set.Contains(i.Item()) {
			s.Add(i.Item())
		}
	}
	return s
}

func (o *orderedSet[T]) Difference(set OrderedSet[T]) OrderedSet[T] {
	s := o.Copy()
	for i := set.Iterator(); i.Next(); {
		s.Remove(i.Item())
	}
	return s
}

func (o *orderedSet[T]) Equal(set OrderedSet[T]) bool {
	lhs := o.head
	rhs := set.(*orderedSet[T]).head

	for {
		if lhs == nil && rhs == nil {
			return true
		}

		if lhs == nil || rhs == nil {
			return false
		}

		if lhs.value != rhs.value {
			return false
		}

		lhs = lhs.next
		rhs = rhs.next
	}
}

func (o *orderedSet[T]) SetEqual(set OrderedSet[T]) bool {
	for i := o.Iterator(); i.Next(); {
		if !set.Contains(i.Item()) {
			return false
		}
	}

	for i := set.Iterator(); i.Next(); {
		if !o.Contains(i.Item()) {
			return false
		}
	}

	return true
}

func (o *orderedSetIterator[T]) Next() bool {
	if o.next == nil {
		return false
	}

	o.item = o.next
	o.next = o.next.next

	return true
}

func (o *orderedSetIterator[T]) Item() T {
	return o.item.value
}
