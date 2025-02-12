package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("should create a new DLL", func(t *testing.T) {
		data := [][]any{
			{1, 2, 3, 4},
			{false, "word", 3.14},
		}
		output := []string{
			"DLL: [1 2 3 4]",
			"DLL: [false word 3.14]",
		}

		for index := range len(data) {
			dll := NewDLL(data[index])
			assert.Equal(t, output[index], dll.String())
		}
	})

	t.Run("should insert new node at the begining of the DLL", func(t *testing.T) {
		data := []*DoublyLinkedList{
			NewDLL([]any{1, 2, 3, 4, 5}),
			NewDLL([]any{false, 2.3, "4", 5}),
			NewDLL([]any{}),
		}
		input := []any{1738, "yeah", false}
		output := []string{
			"DLL: [1738 1 2 3 4 5]",
			"DLL: [yeah false 2.3 4 5]",
			"DLL: [false]",
		}

		for index := range len(data) {
			dll := data[index]
			dll.Insert(input[index])

			assert.Equal(t, output[index], dll.String())
		}
	})

	t.Run("should insert a new node at the end of the DLL", func(t *testing.T) {
		data := []*DoublyLinkedList{
			NewDLL([]any{1, 2, 3, 4, 5}),
			NewDLL([]any{false, 2.3, "4", 5}),
		}
		input := []any{1738, "yeah"}
		output := []string{
			"DLL: [1 2 3 4 5 1738]",
			"DLL: [false 2.3 4 5 yeah]",
		}

		for index := range len(data) {
			dll := data[index]
			dll.InsertLast(input[index])

			assert.Equal(t, output[index], dll.String())
		}
	})

	t.Run("should delete the node at the begining of the DLL", func(t *testing.T) {
		data := []*DoublyLinkedList{
			NewDLL([]any{1, 2, 3, 4, 5}),
			NewDLL([]any{false, 2.3, "4", 5}),
		}
		output := []string{
			"DLL: [2 3 4 5]",
			"DLL: [2.3 4 5]",
		}

		for index := range len(data) {
			dll := data[index]
			dll.Delete()

			assert.Equal(t, output[index], dll.String())
		}
	})

	t.Run("should delete the node at the end of the DLL", func(t *testing.T) {
		data := []*DoublyLinkedList{
			NewDLL([]any{1, 2, 3, 4, 5}),
			NewDLL([]any{false, 2.3, "4", 5}),
		}
		output := []string{
			"DLL: [1 2 3 4]",
			"DLL: [false 2.3 4]",
		}

		for index := range len(data) {
			dll := data[index]
			dll.DeleteLast()

			assert.Equal(t, output[index], dll.String())
		}
	})

	t.Run("should return a string representation of the DLL", func(t *testing.T) {
		data := []*DoublyLinkedList{
			NewDLL([]any{1, 2, 3, 4, 5}),
			NewDLL([]any{6, 7, 8, 9, 10}),
			NewDLL([]any{11, 12, 13, 14, 15}),
			NewDLL([]any{}),
		}
		input := []string{"forward", "backward", "backward", "forward"}
		output := []string{
			"DLL: [1 2 3 4 5]",
			"DLL: [10 9 8 7 6]",
			"DLL: [15 14 13 12 11]",
			"DLL: []",
		}

		for index := range len(data) {
			dll := data[index]
			dllString, err := dll.Display(input[index])

			assert.Nil(t, err)
			assert.Equal(t, output[index], dllString)
		}
	})

	t.Run("should return an error for string representation of the DLL", func(t *testing.T) {
		data := []*DoublyLinkedList{
			NewDLL([]any{1, 2, 3, 4, 5}),
			NewDLL([]any{6, 7, 8, 9, 10}),
		}
		input := []string{"front", "reverse"}

		for index := range len(data) {
			dll := data[index]
			dllString, err := dll.Display(input[index])

			assert.Equal(t, dllString, "")
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), "invalid traversal direction. should be one of [forward backward]")
		}
	})
}
