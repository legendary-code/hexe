package coord

type OrderedSetIterator[T comparable] interface {
	Next() bool
	Item() T
	Index() int
}

// OrderedSet represents an ordered set of coordinates
type OrderedSet[T comparable, TS any] interface {
	Add(T) bool
	Remove(T) bool
	Contains(T) bool
	Clear() bool
	Copy() TS
	IsEmpty() bool
	Size() int
	Iterator() OrderedSetIterator[T]
	Head() T
	Tail() T
	ForEach(func(T) bool)
	ToSlice() []T
	Union(TS) TS
	Intersect(TS) TS
	Difference(TS) TS
	Equal(TS) bool
	SetEqual(TS) bool
	getHead() *item[T]
}

type item[T comparable] struct {
	value T
	next  *item[T]
	prev  *item[T]
}

type orderedSet[T comparable, TS OrderedSet[T, TS]] struct {
	head  *item[T]
	tail  *item[T]
	items map[T]*item[T]
	size  int
	ctor  func() TS
}

type orderedSetIterator[T comparable] struct {
	item  *item[T]
	next  *item[T]
	index int
}

func newOrderedSet[T comparable, TS OrderedSet[T, TS]](ctor func() TS, values ...T) *orderedSet[T, TS] {
	s := &orderedSet[T, TS]{
		head:  nil,
		tail:  nil,
		items: make(map[T]*item[T]),
		size:  0,
		ctor:  ctor,
	}
	for _, value := range values {
		s.Add(value)
	}
	return s
}

func (o *orderedSet[T, TS]) Remove(value T) bool {
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

func (o *orderedSet[T, TS]) Add(value T) bool {
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

func (o *orderedSet[T, TS]) ToSlice() []T {
	items := make([]T, o.size)
	n := 0
	for i := o.head; i != nil; i = i.next {
		items[n] = i.value
		n++
	}
	return items
}

func (o *orderedSet[T, TS]) Contains(value T) bool {
	_, ok := o.items[value]
	return ok
}

func (o *orderedSet[T, TS]) Clear() bool {
	if o.IsEmpty() {
		return false
	}

	o.items = make(map[T]*item[T])
	o.head = nil
	o.tail = nil
	o.size = 0

	return true
}

func (o *orderedSet[T, TS]) Head() T {
	return o.head.value
}

func (o *orderedSet[T, TS]) Tail() T {
	return o.tail.value
}

func (o *orderedSet[T, TS]) IsEmpty() bool {
	return o.size == 0
}

func (o *orderedSet[T, TS]) Iterator() OrderedSetIterator[T] {
	return &orderedSetIterator[T]{next: o.head, index: -1}
}

func (o *orderedSet[T, TS]) ForEach(iterFunc func(T) bool) {
	for i := o.Iterator(); i.Next() && iterFunc(i.Item()); {
	}
}

func (o *orderedSet[T, TS]) Size() int {
	return o.size
}

func (o *orderedSet[T, TS]) Copy() TS {
	s := o.ctor()
	for i := o.Iterator(); i.Next(); {
		s.Add(i.Item())
	}
	return s
}

func (o *orderedSet[T, TS]) Union(set TS) TS {
	s := o.Copy()
	for i := set.Iterator(); i.Next(); {
		s.Add(i.Item())
	}
	return s
}

func (o *orderedSet[T, TS]) Intersect(set TS) TS {
	s := o.ctor()
	for i := o.Iterator(); i.Next(); {
		if set.Contains(i.Item()) {
			s.Add(i.Item())
		}
	}
	return s
}

func (o *orderedSet[T, TS]) getHead() *item[T] {
	return o.head
}

func (o *orderedSet[T, TS]) Difference(set TS) TS {
	s := o.Copy()
	for i := set.Iterator(); i.Next(); {
		s.Remove(i.Item())
	}
	return s
}

func (o *orderedSet[T, TS]) Equal(set TS) bool {
	lhs := o.head
	rhs := set.getHead()

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

func (o *orderedSet[T, TS]) SetEqual(set TS) bool {
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
	o.index++

	return true
}

func (o *orderedSetIterator[T]) Item() T {
	return o.item.value
}

func (o *orderedSetIterator[T]) Index() int {
	return o.index
}
