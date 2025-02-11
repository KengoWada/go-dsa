package linkedlists

import "fmt"

// A node for a singly linked list.
type SLLNode struct {
	Data any
	Next *SLLNode
}

func (n SLLNode) String() string {
	return fmt.Sprintf("%v", n.Data)
}

type SinglyLinkedList struct {
	Head *SLLNode
}

// NewSLL returns a singly linked list from
// the data provided.
func NewSLL(data []any) SinglyLinkedList {
	sll := SinglyLinkedList{}

	for i, value := range data {
		node := &SLLNode{Data: value}
		if i == 0 {
			sll.Head = node
			continue
		}

		lastNode := sll.FindLastNode()
		lastNode.Next = node
	}

	return sll
}

func (sll SinglyLinkedList) String() string {
	data := []string{}
	SLLNode := sll.Head
	for SLLNode != nil {
		data = append(data, SLLNode.String())
		SLLNode = SLLNode.Next
	}

	return fmt.Sprintf("SLL: %v", data)
}

// Insert creates a new node as the head of
// the singly linked list.
func (sll *SinglyLinkedList) Insert(data any) {
	SLLNode := &SLLNode{Data: data, Next: sll.Head}
	sll.Head = SLLNode
}

// InsertEnd creates a new node as the last node
// of the SLL.
func (sll *SinglyLinkedList) InsertEnd(data any) {
	lastSLLNode := sll.FindLastNode()
	lastSLLNode.Next = &SLLNode{Data: data}
}

// InsertAtPos creates a new node in the position
// of, n, node provided.
func (sll *SinglyLinkedList) InsertAtPos(n *SLLNode, data any) {
	if n == nil {
		fmt.Println("invalid SLLNode")
		return
	}

	newNode := &SLLNode{Data: data}

	if n == sll.Head {
		sll.Head = newNode
		newNode.Next = n
		return
	}

	var prevNode *SLLNode
	node := sll.Head
	for node != nil {
		if node.Next == n {
			prevNode = node
			break
		}

		node = node.Next
	}

	if prevNode != nil {
		prevNode.Next = newNode
		newNode.Next = n
	}

}

// FindNode returns the first node in a SLL that
// has a Data value equal to the key.
func (sll *SinglyLinkedList) FindNode(key any) *SLLNode {
	node := sll.Head
	for node != nil {
		if node.Data == key {
			return node
		}
		node = node.Next
	}

	return nil
}

// FindLastNode returns the last node of a SLL.
func (sll *SinglyLinkedList) FindLastNode() *SLLNode {
	node := sll.Head
	for node.Next != nil {
		node = node.Next
	}
	return node
}

// Delete removes the first node from the SLL.
func (sll *SinglyLinkedList) Delete() {
	node := sll.Head
	sll.Head = node.Next
	node.Next = nil
}

// DeleteEnd removes the last node from a SLL.
func (sll *SinglyLinkedList) DeleteEnd() {
	node := sll.Head
	for node.Next.Next != nil {
		node = node.Next
	}

	node.Next = nil
}

// DeleteAtPos removes a node from the SLL that
// has a Data value that matches the provided key.
func (sll *SinglyLinkedList) DeleteAtPos(key any) {
	node := sll.FindNode(key)
	if node == nil {
		return
	}

	if node == sll.Head {
		head := sll.Head
		sll.Head = node.Next
		head.Next = nil
		return
	}

	var prevNode *SLLNode
	n := sll.Head
	for n != nil {
		if n.Next == node {
			prevNode = n
		}

		n = n.Next
	}

	prevNode.Next = node.Next
	node.Next = nil
}

// Reverse reverses the order of the SLL.
func (sll *SinglyLinkedList) Reverse() {
	var prev *SLLNode
	node := sll.Head
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	sll.Head = prev
}
