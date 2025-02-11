package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSLLNode(t *testing.T) {
	t.Run("should return node string", func(t *testing.T) {
		data := []any{"word", 12, false, 3.14}
		output := []string{"word", "12", "false", "3.14"}

		for index := range len(data) {
			node := SLLNode{Data: data[index], Next: nil}
			assert.Equal(t, output[index], node.String())
		}
	})
}

func TestSinglyLinkedList(t *testing.T) {
	t.Run("should create SLL", func(t *testing.T) {
		data := [][]any{
			{1, 2, 3, 4},
			{"we", "they", "here", "there"},
			{false, true, true, false},
		}
		output := []string{
			"SLL: [1 2 3 4]",
			"SLL: [we they here there]",
			"SLL: [false true true false]",
		}

		for index := range len(data) {
			sll := NewSLL(data[index])
			assert.Equal(t, output[index], sll.String())
		}
	})

	t.Run("should insert new node at head of the SLL", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		input := []any{"word", 1, true}
		output := []string{
			"SLL: [word 1 2 3 4]",
			"SLL: [1 name false]",
			"SLL: [true 1.75 2.3 3.14 4.7]",
		}
		for index := range len(data) {
			sll := data[index]
			sll.Insert(input[index])
			assert.Equal(t, output[index], sll.String())
		}
	})

	t.Run("should insert new node after the last node", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		input := []any{"word", 1, true}
		output := []string{
			"SLL: [1 2 3 4 word]",
			"SLL: [name false 1]",
			"SLL: [1.75 2.3 3.14 4.7 true]",
		}

		for index := range len(data) {
			sll := data[index]
			sll.InsertEnd(input[index])
			assert.Equal(t, output[index], sll.String())
		}
	})

	t.Run("should find node in sll", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		input := []any{1, "true", 3.14}
		output := []string{"1", "", "3.14"}

		for index := range len(data) {
			key := input[index]

			sll := data[index]
			n := sll.FindNode(key)
			if key == "true" {
				assert.Nil(t, n)
				continue
			}

			assert.Equal(t, output[index], n.String())
		}
	})

	t.Run("should insert a new node at the position of the node provided", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		input := []any{1, "true", 3.94, 2.3}
		output := []string{
			"SLL: [1 1 2 3 4]",
			"SLL: [true name false]",
			"SLL: [1.75 2.3 3.14 3.94 4.7]",
			"SLL: [1.75 2.3 3.14 4.7]",
		}
		nodes := []*SLLNode{
			data[0].FindNode(2),
			data[1].FindNode("name"),
			data[2].FindNode(4.7),
			nil,
		}

		for index := range len(data) {
			sll := data[index]
			sll.InsertAtPos(nodes[index], input[index])
			assert.Equal(t, output[index], sll.String())
		}
	})

	t.Run("should return the last node of the SLL", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		output := []any{4, false, 4.7}

		for index := range len(data) {
			sll := data[index]
			node := sll.FindLastNode()

			assert.NotNil(t, node)
			assert.Nil(t, node.Next)
			assert.Equal(t, output[index], node.Data)
		}
	})

	t.Run("should remove the first element of the SLL", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		output := []string{
			"SLL: [2 3 4]",
			"SLL: [false]",
			"SLL: [2.3 3.14 4.7]",
		}

		for index := range len(data) {
			sll := data[index]
			sll.Delete()

			assert.Equal(t, output[index], sll.String())
		}
	})

	t.Run("should remove the last node of the SLL", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		output := []string{
			"SLL: [1 2 3]",
			"SLL: [name]",
			"SLL: [1.75 2.3 3.14]",
		}

		for index := range len(data) {
			sll := data[index]
			sll.DeleteEnd()

			assert.Equal(t, output[index], sll.String())
		}
	})

	t.Run("should remove a node that has Data field that matches key", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		input := []any{1, "true", 2.3}
		output := []string{
			"SLL: [2 3 4]",
			"SLL: [name false]",
			"SLL: [1.75 3.14 4.7]",
		}

		for index := range len(data) {
			sll := data[index]
			sll.DeleteAtPos(input[index])

			assert.Equal(t, output[index], sll.String())
		}
	})

	t.Run("should reverse the order of the SLL", func(t *testing.T) {
		data := []SinglyLinkedList{
			NewSLL([]any{1, 2, 3, 4}),
			NewSLL([]any{"name", false}),
			NewSLL([]any{1.75, 2.3, 3.14, 4.7}),
		}
		output := []string{
			"SLL: [4 3 2 1]",
			"SLL: [false name]",
			"SLL: [4.7 3.14 2.3 1.75]",
		}

		for index := range len(data) {
			sll := data[index]
			sll.Reverse()
			assert.Equal(t, output[index], sll.String())
		}
	})
}
