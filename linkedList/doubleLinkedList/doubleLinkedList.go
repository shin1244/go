package doublelinkedlist

type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

func (l *LinkedList[T]) PushBack(val T) {
	newNode := &Node[T]{
		Value: val,
	}
	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		l.count = 1
		return
	}
	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
	l.count++
}

func (l *LinkedList[T]) PushFront(val T) {
	newNode := &Node[T]{
		Value: val,
	}
	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		l.count = 1
		return
	}
	l.root.prev = newNode
	newNode.next = l.root
	l.root = newNode
	l.count++
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if l.count <= idx {
		return nil
	}
	nowNode := l.root
	for i := 0; i < idx; i++ {
		nowNode = nowNode.next
	}
	return nowNode
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	nowNode := l.root
	for ; nowNode != nil; nowNode = nowNode.next {
		if nowNode == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{
		Value: val,
	}
	newNode.next = node.next
	node.next = newNode
	newNode.prev = node
	if newNode.next != nil {
		newNode.next.prev = newNode
	}
	if l.tail == node {
		l.tail = newNode
	}
	l.count++
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{
		Value: val,
	}
	newNode.prev = node.prev
	newNode.next = node
	node.prev = newNode
	if newNode.prev != nil {
		newNode.prev.next = newNode
	}
	if node == l.root {
		l.root = newNode
	}
	l.count++
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}
	targetNode := l.root
	if targetNode.next != nil {
		targetNode.next.prev = nil
		l.root = targetNode.next
		l.count--
	} else {
		l.root = nil
		l.tail = nil
		l.count = 0
	}
	return targetNode
}

func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}
	targetNode := l.tail
	if targetNode.prev != nil {
		targetNode.prev.next = nil
		l.tail = targetNode.prev
		l.count--
	} else {
		l.root = nil
		l.tail = nil
		l.count = 0
	}
	return targetNode
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if !l.isIncluded(node) {
		return
	}
	if node == l.root {
		l.PopFront()
		return
	} else if node == l.tail {
		l.PopBack()
		return
	}
	prevNode := node.prev
	nextNode := node.next

	prevNode.next, nextNode.prev = nextNode, prevNode
	l.count--
}

func (l *LinkedList[T]) Reverse() {
	if l.root == nil {
		return
	}
	nowNode := l.root
	var nextNode *Node[T]

	for nowNode != nil {
		nextNode = nowNode.next
		nowNode.prev, nowNode.next = nextNode, nowNode.prev
		nowNode = nextNode
	}
	l.root, l.tail = l.tail, l.root
}

func (l *LinkedList[T]) ListToArray() []any {
	arr := []any{}
	node := l.root
	for ; node != nil; node = node.next {
		arr = append(arr, node.Value)
	}

	return arr
}
