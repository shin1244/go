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

func Push[T any](node *Node[T], next *Node[T]) {
	next.next = node.next
	node.next = next
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
}
