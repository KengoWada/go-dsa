package stackandqueue

import "fmt"

type Stack struct {
	Top     int
	MaxSize int
	Data    []any
}

func NewStack(maxSize int) *Stack {
	return &Stack{
		Top:     -1,
		MaxSize: maxSize,
		Data:    make([]any, maxSize),
	}
}

func (s *Stack) IsFull() bool {
	return s.Top == s.MaxSize-1
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) Push(data any) error {
	if s.IsFull() {
		return fmt.Errorf("could not insert %v, stack is full", data)
	}

	s.Top += 1
	s.Data[s.Top] = data

	return nil
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("can not pop from empty stack")
	}

	data := s.Data[s.Top]
	s.Top -= 1
	s.Data = s.Data[:s.Top]
	return data, nil
}

func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return nil
	}

	return s.Data[s.Top]
}
