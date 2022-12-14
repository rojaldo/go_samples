package main

type Node[T any] struct {
	Val    T
	Father *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

type Tree struct {
	Root *Node[int]
}

func (t *Tree) AddNode(node *Node[int], right bool) error {
	if t.Root == nil {
		t.Root = node
		return nil
	}
	myElem := t.Root
	for myElem != nil {
		if right {
			if myElem.Right == nil {
				myElem.Right = node
				node.Father = myElem
				return nil
			}
			myElem = myElem.Right
		} else {
			if myElem.Left == nil {
				myElem.Left = node
				node.Father = myElem
				return nil
			}
			myElem = myElem.Left
		}
	}
	return nil
}

func (t *Tree) PrintTree() error {
	myElem := t.Root
	println(myElem.Val)
	println(myElem.Left.Val)
	println(myElem.Right.Val)
	return nil
}

func main() {
	tree := Tree{}
	tree.AddNode(&Node[int]{Val: 1}, false)
	tree.AddNode(&Node[int]{Val: 2}, false)
	tree.AddNode(&Node[int]{Val: 3}, true)
	tree.PrintTree()
}
