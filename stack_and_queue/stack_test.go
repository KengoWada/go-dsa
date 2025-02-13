package stackandqueue

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("should return a new stack", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}

		for _, maxSize := range input {
			stack := NewStack(maxSize)

			assert.Equal(t, maxSize, cap(stack.Data))
			assert.Equal(t, maxSize, stack.MaxSize)
			assert.Equal(t, -1, stack.Top)
			assert.False(t, stack.IsFull())
		}
	})

	t.Run("should add an element to the stack", func(t *testing.T) {
		stack := NewStack(4)

		data := []any{1, "two", 3.14, false}
		for index, value := range data {
			err := stack.Push(value)

			assert.Nil(t, err)
			assert.True(t, slices.Contains(stack.Data, value))
			assert.Equal(t, index, stack.Top)
		}

		err := stack.Push("fail")
		assert.Error(t, err)
	})

	t.Run("should remove the element at the top of the stack", func(t *testing.T) {
		stacks := []*Stack{
			{MaxSize: 4, Top: 3, Data: []any{1, 2, 3, 4}},
			{MaxSize: 5, Top: 2, Data: []any{"a", "b", "c"}},
		}
		output := []any{4, "c"}

		for index, stack := range stacks {
			top := stack.Top

			data, err := stack.Pop()
			assert.Nil(t, err)
			assert.Less(t, stack.Top, top)
			assert.Equal(t, output[index], data)
		}

		s := NewStack(4)
		data, err := s.Pop()
		assert.Error(t, err)
		assert.Nil(t, data)
	})

	t.Run("should return the top most element of the stack", func(t *testing.T) {
		stacks := []*Stack{
			{MaxSize: 4, Top: 3, Data: []any{1, 2, 3, 4}},
			{MaxSize: 5, Top: 2, Data: []any{"a", "b", "c"}},
		}
		output := []any{4, "c"}

		for index, stack := range stacks {
			top := stack.Peek()
			assert.NotNil(t, top)
			assert.Equal(t, output[index], top)
		}

		s := NewStack(4)
		top := s.Peek()
		assert.Nil(t, top)
	})
}
