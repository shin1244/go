package singlelinkedlist

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	Root *Node[T]
	Tail *Node[T]
}

func (l *LinkedList[T]) PushBack(val T) {
	node := &Node[T]{
		Value: val,
	}
	if l.Root == nil {
		l.Root = node
		l.Tail = node
		return
	}
	l.Tail.Next = node
	l.Tail = node
}

func (l *LinkedList[T]) PushFront(val T) {
	node := &Node[T]{
		Value: val,
	}
	if l.Root == nil {
		l.Root = node
		l.Tail = node
		return
	}
	node.Next = l.Root
	l.Root = node
}

func (l *LinkedList[T]) Front(val T) *Node[T] {
	return l.Root
}

func (l *LinkedList[T]) Back(val T) *Node[T] {
	return l.Tail
}

func (l *LinkedList[T]) Count() int {
	node := l.Root
	cnt := 0

	for ; node != nil; node = node.Next {
		cnt++
	}
	return cnt
}

func (l *LinkedList[T]) GetAt(n int) *Node[T] {
	node := l.Root
	if node == nil {
		return nil
	}
	for i := 0; i < n; i++ {
		if node.Next == nil {
			return nil
		}
		node = node.Next
	}
	return node
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{
		Next:  node.Next,
		Value: val,
	}
	node.Next = newNode
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	nowNode := l.Root

	for ; nowNode != nil; nowNode = nowNode.Next {
		if node == nowNode {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if l.Root == node {
		l.PushFront(val)
		return
	}
	newNode := &Node[T]{
		Next:  node,
		Value: val,
	}
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}
	prevNode.Next = newNode
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	nowNode := l.Root

	for ; nowNode != nil; nowNode = nowNode.Next {
		if node == nowNode.Next {
			return nowNode
		}
	}
	return nil
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.Root == nil {
		return nil
	}
	node := l.Root
	l.Root.Next, l.Root = nil, node.Next
	if l.Root == nil {
		l.Tail = nil
	}

	return node
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if l.Root == node {
		l.PopFront()
		return
	}

	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}
	prevNode.Next = node.Next
	node.Next = nil

	if node == l.Tail {
		l.Tail = prevNode
	}
}

func (l *LinkedList[T]) Reverse() {
	newLinkedList := &LinkedList[T]{}
	for l.Root != nil {
		node := l.PopFront()
		newLinkedList.PushFront(node.Value)
	}
	l.Root = newLinkedList.Root
	l.Tail = newLinkedList.Tail
}

func (l *LinkedList[T]) ListToArray() []any {
	arr := []any{}
	node := l.Root
	for ; node != nil; node = node.Next {
		arr = append(arr, node.Value)
	}

	return arr
}
