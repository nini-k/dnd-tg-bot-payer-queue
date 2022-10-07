package circular_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularLinkedList(t *testing.T) {
	t.Run("pop element from empty list", func(t *testing.T) {
		list := New[int]()

		actual := list.Pop()

		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)

		assert.Equal(t, 0, actual)
		assert.Equal(t, 0, list.Size())
	})

	t.Run("pop element from list with one element", func(t *testing.T) {
		const el = 1

		list := New[int](el)

		actual := list.Pop()

		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)

		assert.Equal(t, el, actual)
		assert.Equal(t, 0, list.Size())
	})

	t.Run("pop element from list with two elements", func(t *testing.T) {
		const (
			el0 = 1
			el1 = 2
		)

		list := New[int](el0, el1)

		actual := list.Pop()

		assert.NotNil(t, list.tail)
		assert.NotNil(t, list.head)

		assert.Equal(t, el0, actual)
		assert.Equal(t, 1, list.Size())

		assert.NotNil(t, list.tail)
		assert.NotNil(t, list.head)

		assert.NotNil(t, list.tail.next)
		assert.NotNil(t, list.head.next)

		assert.Equal(t, list.head.value, el1)
		assert.Equal(t, list.tail.value, el1)

		assert.Equal(t, list.head, list.tail)
		assert.Equal(t, list.head.next, list.head)
		assert.Equal(t, list.tail.next, list.tail)
	})

	t.Run("pop n elements to list", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		list := New[int](elements...)

		for _, v := range elements {
			actual := list.Pop()
			assert.Equal(t, v, actual)
		}

		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)

		assert.Equal(t, 0, list.Size())
	})

	t.Run("push one element to list", func(t *testing.T) {
		const el = 0

		list := New[int]()

		list.Push(el)

		assert.NotNil(t, list.tail)
		assert.NotNil(t, list.head)

		assert.NotNil(t, list.tail.next)
		assert.NotNil(t, list.head.next)

		assert.Equal(t, list.head.value, el)
		assert.Equal(t, list.tail.value, el)

		assert.Equal(t, list.head, list.tail)
		assert.Equal(t, list.head.next, list.head)
		assert.Equal(t, list.tail.next, list.tail)

		assert.Equal(t, 1, list.Size())
	})

	t.Run("push two elements to list", func(t *testing.T) {
		const (
			el0 = iota
			el1
		)

		list := New[int]()

		list.Push(el0)

		assert.NotNil(t, list.tail)
		assert.NotNil(t, list.head)

		assert.Equal(t, list.head.value, el0)
		assert.Equal(t, list.tail.value, el0)

		list.Push(el1)

		assert.Equal(t, list.head.value, el0)
		assert.Equal(t, list.tail.value, el1)

		assert.NotNil(t, list.tail.next)
		assert.NotNil(t, list.head.next)

		assert.Equal(t, list.head.next, list.tail)
		assert.Equal(t, list.tail.next, list.head)

		assert.Equal(t, 2, list.Size())
	})

	t.Run("push n elements to list", func(t *testing.T) {
		elements := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		list := New[int]()

		for _, v := range elements {
			list.Push(v)

			assert.Equal(t, list.tail.value, v)
			assert.Equal(t, list.head.value, elements[0])
		}

		assert.Equal(t, len(elements), list.Size())
	})

	t.Run("remove element by condition from list with one element", func(t *testing.T) {
		const el = 1

		list := New[int](el)

		list.RemoveByCond(func(val int) bool {
			return val == el
		})

		assert.Nil(t, list.tail)
		assert.Nil(t, list.head)
		assert.Equal(t, 0, list.Size())
	})

	t.Run("remove head by condition from list with two elements", func(t *testing.T) {
		const (
			el0 = iota
			el1
		)

		list := New[int](el0, el1)

		list.RemoveByCond(func(val int) bool {
			return val == el0
		})

		assert.NotNil(t, list.tail)
		assert.NotNil(t, list.head)

		assert.NotNil(t, list.tail.next)
		assert.NotNil(t, list.head.next)

		assert.Equal(t, list.head.value, el1)
		assert.Equal(t, list.tail.value, el1)

		assert.Equal(t, list.head, list.tail)
		assert.Equal(t, list.head.next, list.head)
		assert.Equal(t, list.tail.next, list.tail)

		assert.Equal(t, 1, list.Size())
	})

	t.Run("remove tail by condition from list with two elements", func(t *testing.T) {
		const (
			el0 = iota
			el1
		)

		list := New[int](el0, el1)

		list.RemoveByCond(func(val int) bool {
			return val == el1
		})

		assert.NotNil(t, list.tail)
		assert.NotNil(t, list.head)
		assert.Equal(t, list.head.value, el0)
		assert.Equal(t, list.tail.value, el0)
		assert.Equal(t, list.head, list.tail)
		assert.Equal(t, 1, list.Size())
	})

	t.Run("remove head by condition from list with n elements", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		list := New[int](elements...)

		list.RemoveByCond(func(val int) bool {
			return val == elements[0]
		})

		assert.NotNil(t, list.tail)
		assert.NotNil(t, list.head)

		assert.Equal(t, list.head.value, elements[1])
		assert.Equal(t, list.tail.value, elements[len(elements)-1])
		assert.Equal(t, list.tail.next, list.head)

		assert.Equal(t, len(elements)-1, list.Size())
	})

	t.Run("remove tail by condition from list with n elements", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		list := New[int](elements...)

		list.RemoveByCond(func(val int) bool {
			return val == elements[len(elements)-1]
		})

		assert.NotNil(t, list.tail)
		assert.NotNil(t, list.head)

		assert.Equal(t, list.head.value, elements[0])
		assert.Equal(t, list.tail.value, elements[len(elements)-2])
		assert.Equal(t, list.tail.next, list.head)

		assert.Equal(t, len(elements)-1, list.Size())
	})

	t.Run("list convert to slice", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		list := New[int](elements...)

		head, tail := list.head, list.tail

		actual := list.ConvertToSlice()
		assert.Equal(t, cap(elements), cap(actual))
		assert.Equal(t, elements, actual)

		assert.Equal(t, len(elements), list.Size())
		assert.Equal(t, head, list.head)
		assert.Equal(t, tail, list.tail)
		assert.Equal(t, list.tail.next, list.head)
	})

	t.Run("for each iterate", func(t *testing.T) {
		elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		list := New[int](elements...)

		head, tail := list.head, list.tail

		actual := make([]int, 0, len(elements))
		list.ForEach(func(val int) {
			actual = append(actual, val)
		})

		assert.Equal(t, cap(elements), cap(actual))
		assert.Equal(t, elements, actual)

		assert.Equal(t, len(elements), list.Size())
		assert.Equal(t, head, list.head)
		assert.Equal(t, tail, list.tail)
		assert.Equal(t, list.tail.next, list.head)
	})
}
