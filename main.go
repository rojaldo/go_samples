package main

import (
	"errors"
)

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(node *Node) {
	if s.Top == nil {
		s.Top = node
		return
	}
	node.Next = s.Top
	s.Top = node
}

func (s *Stack) Pop() (*Node, error) {
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

func main() {

	stack := Stack{}
	stack.Push(&Node{Val: 12})
	stack.Push(&Node{Val: 14})
	stack.Push(&Node{Val: 16})
	stack.PrintStack()
	println("Popping")
	stack.Pop()
	// println(res.Val)
	stack.PrintStack()
	println("Popping")
	_, error := stack.Pop()
	if error != nil {
		println(error.Error())
	}
	stack.PrintStack()
	println("Popping")
	_, error = stack.Pop()
	if error != nil {
		println(error.Error())
	}
	stack.PrintStack()
	println("Last Popping")
	_, error = stack.Pop()
	if error != nil {
		println(error.Error())
	}

}
