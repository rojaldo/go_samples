package structures

import (
	"errors"
)

type StackNode[T any] struct {
	Val  T
	Next *StackNode[T]
}

type Stack[T any] struct {
	Top *StackNode[T]
}

func (s *Stack[T]) NewStack() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(node *StackNode[T]) {
	if s.Top == nil {
		s.Top = node
		return
	}
	node.Next = s.Top
	s.Top = node
}

func (s *Stack[T]) Pop() (*StackNode[T], error) {
	if s.Top == nil {
		return nil, errors.New("Stack is empty")
	}
	res := s.Top
	s.Top = s.Top.Next
	return res, nil
}

func (s *Stack[T]) PrintStack() {
	myElem := s.Top
	for myElem != nil {
		println(myElem.Val)
		myElem = myElem.Next
	}
}
