package linkedlists

import (
	"fmt"
	"slices"
)

var directions = []string{"forward", "backward"}

type DLLNode struct {
	Data any
	Prev *DLLNode
	Next *DLLNode
}

// NewDLLNode returns a new DLLNode with the values provided.
func NewDLLNode(data any, prev, next *DLLNode) *DLLNode {
	return &DLLNode{
		Prev: prev,
		Data: data,
		Next: next,
	}
}

func (dlln DLLNode) String() string {
	return fmt.Sprintf("DLLNode: %v", dlln.Data)
}

type DoublyLinkedList struct {
	Head *DLLNode
	Last *DLLNode
}

func NewDLL(data []any) *DoublyLinkedList {
	dll := &DoublyLinkedList{}
	var prevNode *DLLNode

	for index := range len(data) {
		node := NewDLLNode(data[index], nil, nil)
		if index == 0 {
			dll.Head = node
			prevNode = node
			continue
		}

		prevNode.Next = node
		node.Prev = prevNode
		prevNode = node
	}

	dll.Last = prevNode
	return dll
}

func (dll DoublyLinkedList) String() string {
	data := []string{}
	node := dll.Head
	for node != nil {
		data = append(data, fmt.Sprintf("%v", node.Data))
		node = node.Next
	}

	return fmt.Sprintf("DLL: %v", data)
}

func (dll *DoublyLinkedList) Display(direction string) (string, error) {
	if !slices.Contains(directions, direction) {
		return "", fmt.Errorf("invalid traversal direction. should be one of %v", directions)
	}

	if dll.Head == nil {
		return "DLL: []", nil
	}

	var data []string
	var node *DLLNode
	isForward := direction == directions[0]
	if isForward {
		node = dll.Head
	} else {
		node = dll.Last
	}

	for node != nil {
		data = append(data, fmt.Sprintf("%v", node.Data))
		if isForward {
			node = node.Next
		} else {
			node = node.Prev
		}
	}

	return fmt.Sprintf("DLL: %v", data), nil
}

func (dll *DoublyLinkedList) Insert(data any) {
	node := NewDLLNode(data, nil, nil)

	if dll.Head == nil {
		dll.Head = node
		return
	}

	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
}

func (dll *DoublyLinkedList) InsertLast(data any) {
	node := NewDLLNode(data, dll.Last, nil)
	dll.Last.Next = node
	dll.Last = node
}

func (dll *DoublyLinkedList) Delete() {
	node := dll.Head
	dll.Head = node.Next
	dll.Head.Prev = nil
	node.Next = nil
}

func (dll *DoublyLinkedList) DeleteLast() {
	node := dll.Last
	dll.Last = node.Prev
	dll.Last.Next = nil
	node.Prev = nil
}
