package structures

import (
	"errors"
)

type StackNode struct {
	Val  int
	Next *StackNode
}

type Stack struct {
	Top *StackNode
}

func (s *Stack) Push(node *StackNode) {
	if s.Top == nil {
		s.Top = node
		return
	}
	node.Next = s.Top
	s.Top = node
}

func (s *Stack) Pop() (*StackNode, error) {
	if s.Top == nil {
		return nil, errors.New("Stack is empty")
	}
	res := s.Top
	s.Top = s.Top.Next
	return res, nil
}

func (s *Stack) PrintStack() {
	myElem := s.Top
	for myElem != nil {
		println(myElem.Val)
		myElem = myElem.Next
	}
}
