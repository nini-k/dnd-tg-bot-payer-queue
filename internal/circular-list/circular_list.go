package circular_list

type node[T any] struct {
	value T
	next  *node[T]
}

type CircularLinkedList[T any] struct {
	head    *node[T]
	tail    *node[T]
	current *node[T]
	size    int
}

func New[T any](values ...T) CircularLinkedList[T] {
	list := CircularLinkedList[T]{}
	for _, v := range values {
		list.Push(v)
	}

	return list
}

func (l *CircularLinkedList[T]) Push(value T) {
	if l.head == nil {
		l.head = &node[T]{value: value}
		l.head.next = l.head
		l.current = l.head
		l.tail = l.head
	} else {
		tmp := l.tail.next
		l.tail.next = &node[T]{value: value, next: tmp}
		l.tail = l.tail.next
	}
	l.size++
}

func (l *CircularLinkedList[T]) Pop() (value T) {
	if l.head == nil {
		return
	}

	if l.head == l.tail {
		value = l.head.value
		l.head.next = nil
		l.head = nil
		l.tail.next = nil
		l.tail = nil
	} else {
		value = l.head.value
		tmp := l.head.next
		l.head.next = nil
		l.head = tmp
		l.tail.next = l.head
	}
	l.size--

	return value
}

func (l *CircularLinkedList[T]) Current() (value T) {
	return l.current.value
}

func (l *CircularLinkedList[T]) Next() {
	l.current = l.current.next
}

func (l *CircularLinkedList[T]) RemoveByCond(cond func(node T) bool) {
	prev := l.head
	for prev != nil {
		cur := prev.next
		if cond(cur.value) {
			if cur == l.tail && cur == l.head {
				l.tail.next = nil
				l.head.next = nil
				l.head = nil
				l.tail = nil

				l.size--
				break
			}

			if cur == l.tail {
				l.tail = prev
			}

			if cur == l.head {
				l.head = cur.next
				l.tail.next = l.head
			}

			prev.next = cur.next
			l.size--
			break
		}

		prev = prev.next

		if prev == l.head {
			break
		}
	}
}

func (l *CircularLinkedList[T]) ForEach(fn func(value T)) {
	cur, size := l.head, l.size
	for cur != nil && size > 0 {
		fn(cur.value)
		cur = cur.next
		size--
	}
}

func (l *CircularLinkedList[T]) Size() int {
	return l.size
}

func (l *CircularLinkedList[T]) ConvertToSlice() []T {
	out, cur, size := make([]T, 0, l.size), l.head, l.size
	for cur != nil && size > 0 {
		out = append(out, cur.value)
		cur = cur.next
		size--
	}

	return out
}
