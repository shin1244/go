package main

import "fmt"

type Node[T any] struct {
	next *Node
	val  T
}

func main() {
	fmt.Println("하이요")
}
