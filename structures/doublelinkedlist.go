package structures

import (
	"errors"
)

type LinkedNode struct {
	Val  int
	Next *LinkedNode
	Prev *LinkedNode
}

type DoubleLinkedList struct {
	Head *LinkedNode
	Tail *LinkedNode
}

func (l *DoubleLinkedList) AddNode(node *LinkedNode) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	temp := l.Tail
	l.Tail.Next = node
	node.Prev = temp
	l.Tail = node
}

func (l *DoubleLinkedList) RemoveNode(node *LinkedNode) error {
	if node == nil {
		return errors.New("Node is nil")
	}
	if node.Val == l.Head.Val {
		l.Head = l.Head.Next
		l.Head.Prev = nil
		return nil
	} else if node.Val == l.Tail.Val {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		return nil
	} else {
		// node is in the middle, I need to find it
		myElem := l.Head
		for myElem != nil {
			if myElem.Val == node.Val {
				// remove node
				myElem.Prev.Next = myElem.Next
				myElem.Next.Prev = myElem.Prev
				return nil
			}
			myElem = myElem.Next
		}
		// node not found
		return errors.New("Node not found")
	}
}

func (l *DoubleLinkedList) PrintList() error {
	myElem := l.Head
	if myElem == nil {
		return errors.New("List is empty")
	}
	print("Head -> ")
	for myElem != nil {
		print(myElem.Val, ", ")
		myElem = myElem.Next
	}
	println("<- Tail")
	return nil
}
