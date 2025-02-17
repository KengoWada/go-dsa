package stackandqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	t.Run("should insert element at the back of the queue", func(t *testing.T) {
		data := []Queue{
			NewQueue([]int{1, 2, 3, 4}),
			NewQueue([]int{7, 8, 9, 0}),
			NewQueue([]int{90, 91, 92, 93}),
			NewQueue([]int{72, 73, 74, 75}),
		}
		input := []int{12, 23, 34, 45}

		for index, queue := range data {
			queueLength := len(queue.Data)
			queueFront := queue.Front

			queue.Enqueue(input[index])

			assert.Equal(t, input[index], queue.Data[0])
			assert.Equal(t, len(queue.Data), queueLength+1)
			assert.Equal(t, queueFront+1, queue.Front)
		}

	})

	t.Run("should remove element from the front of the queue", func(t *testing.T) {
		data := []Queue{
			NewQueue([]int{1, 2, 3, 4}),
			NewQueue([]int{7, 8, 9, 0}),
			NewQueue([]int{90, 91, 92, 93}),
			NewQueue([]int{72, 73, 74, 75}),
		}
		output := []int{4, 0, 93, 75}

		for index, queue := range data {
			queueLength := len(queue.Data)
			queueFront := queue.Front

			data := queue.Dequeue()

			assert.Equal(t, output[index], data)
			assert.Equal(t, len(queue.Data), queueLength-1)
			assert.Equal(t, queueFront-1, queue.Front)
		}
	})

	t.Run("should return element at the front of queue", func(t *testing.T) {
		data := []Queue{
			NewQueue([]int{1, 2, 3, 4}),
			NewQueue([]int{7, 8, 9, 0}),
			NewQueue([]int{90, 91, 92, 93}),
			NewQueue([]int{72, 73, 74, 75}),
		}
		output := []int{4, 0, 93, 75}

		for index, queue := range data {
			queueLength := len(queue.Data)
			queueFront := queue.Front

			data := queue.Peek()
			assert.Equal(t, output[index], data)
			assert.Equal(t, len(queue.Data), queueLength)
			assert.Equal(t, queueFront, queue.Front)
		}
	})

	t.Run("should panic when dequeueing or peeking empty queue", func(t *testing.T) {
		queue := NewQueue([]int{})
		assert.Panics(t, func() { queue.Dequeue() })
		assert.Panics(t, func() { queue.Peek() })
	})
}

func TestStaticQueue(t *testing.T) {
	t.Run("should panic when data size is greater than max size", func(t *testing.T) {
		input := []map[string]any{
			{"maxSize": 4, "data": []int{1, 2, 3, 4, 5}},
			{"maxSize": 5, "data": []int{1, 2, 3, 4, 5, 7, 8}},
			{"maxSize": 2, "data": []int{1, 2, 3, 4, 5}},
			{"maxSize": 6, "data": []int{1, 2, 3, 4, 5, 6, 0, 8}},
		}

		for _, value := range input {
			assert.Panics(t, func() { NewStaticQueue(value["maxSize"].(int), value["data"].([]int)) })
		}
	})

	t.Run("should panic trying to insert in a full queue", func(t *testing.T) {
		data := []StaticQueue{
			NewStaticQueue(3, []int{1, 2, 4}),
			NewStaticQueue(5, []int{1, 2, 4, 7, 3}),
			NewStaticQueue(4, []int{1, 2, 4, 0}),
		}
		input := []int{23, 97, 44}

		for index, queue := range data {
			assert.Panics(t, func() { queue.Enqueue(input[index]) })
		}
	})

	t.Run("should insert at the back of the queue", func(t *testing.T) {
		data := []StaticQueue{
			NewStaticQueue(9, []int{1, 2, 4}),
			NewStaticQueue(9, []int{1, 2, 4, 7, 3}),
			NewStaticQueue(9, []int{1, 2, 4, 0}),
		}
		input := []int{23, 97, 44}

		for index, queue := range data {
			queue.Enqueue(input[index])
			assert.Equal(t, input[index], queue.Data[queue.Rear])
		}
	})

	t.Run("should panic when peeking or dequeueing empty queue", func(t *testing.T) {
		queue := NewStaticQueue(9, []int{})
		assert.Panics(t, func() { queue.Peek() })
		assert.Panics(t, func() { queue.Dequeue() })
	})

	t.Run("should remove and return front element of the queue", func(t *testing.T) {
		data := []StaticQueue{
			NewStaticQueue(9, []int{1, 2, 4}),
			NewStaticQueue(9, []int{1, 2, 4, 7, 3}),
			NewStaticQueue(9, []int{1, 2, 4, 0}),
		}
		output := []int{1, 1, 1}

		for index, queue := range data {
			queueLength := len(queue.Data)
			queueRear := queue.Rear

			item := queue.Dequeue()

			assert.Equal(t, output[index], item)
			assert.Equal(t, queue.Rear, queueRear-1)
			assert.Equal(t, len(queue.Data), queueLength-1)
		}
	})

	t.Run("should return and not remove front element of the queue", func(t *testing.T) {
		data := []StaticQueue{
			NewStaticQueue(9, []int{1, 2, 4}),
			NewStaticQueue(9, []int{1, 2, 4, 7, 3}),
			NewStaticQueue(9, []int{1, 2, 4, 0}),
		}
		output := []int{1, 1, 1}

		for index, queue := range data {
			queueLength := len(queue.Data)
			queueRear := queue.Rear

			item := queue.Peek()

			assert.Equal(t, output[index], item)
			assert.Equal(t, queue.Rear, queueRear)
			assert.Equal(t, len(queue.Data), queueLength)
		}
	})
}
