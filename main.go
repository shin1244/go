package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func Append[T any](node *Node[T], next *Node[T]) *Node[T] {
	node.next = next
	return next
}

func Push[T any](root *Node[T], idx int, node *Node[T]) *Node[T] {
	if idx == 0 {
		node.next = root
		root = node
		return root
	}
	now := root
	for i := 1; i < idx; i++ {
		now = now.next
	}
	node.next = now.next
	now.next = node

	return root
}

func main() {
	root := &Node[int]{nil, 10}
	tail := root

	for i := 1; i < 4; i++ {
		tail = Append(tail, &Node[int]{nil, 10 + i*10})
	}

	for n := root; n != nil; n = n.next {
		fmt.Println(n.val)
	}

	fmt.Println()
	fmt.Println()
	root = Push(root, 4, &Node[int]{nil, 100})
	for n := root; n != nil; n = n.next {
		fmt.Println(n.val)
	}
}
