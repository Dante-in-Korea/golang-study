package doublelinkedlist

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]

	Value T
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

func (l *LinkedList[T]) PushBack(val T) {
	node := &Node[T]{
		Value: val,
	}
	if l.tail == nil {
		l.root = node
		l.tail = node

		l.count = 1
		return
	}
	l.tail.next = node // it has already tail
	node.prev = l.tail
	l.tail = node
	l.count++
}

func (l *LinkedList[T]) PushFront(val T) {
	node := &Node[T]{
		Value: val,
	}

	if l.root == nil {
		l.root = node
		l.tail = node
		l.count = 1
		return
	}

	l.root.prev = node // it has already root
	node.next = l.root
	l.root = node
	l.count++

}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count() {
		return nil
	}

	i := 0
	for node := l.root; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}

	n := &Node[T]{
		Value: val,
	}

	nextNode := node.next
	node.next = n     // 기존 노드의 넥스트에 새로운 노드를 추가
	n.next = nextNode // 새로운 노드의 넥스트를 원래 넥스트 노드를 추가

	n.prev = node
	if nextNode != nil { // 추가하고자 할 next node가 nil인지 아닌지 확인
		nextNode.prev = n // 기존 노드의 prev에 새로운 노드를 추가
	}

	if node == l.tail { // 추가 되는 노드가 tail과 같다면 tail에 새로운 노드 추가
		l.tail = n
	}
	l.count++
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}

	n := &Node[T]{
		Value: val,
	}

	prevNode := node.prev // 이전 노드 기록
	node.prev = n         // 이전 노드에 새로운 노드 교체
	n.next = node         // 새로운 노드에 현재 노드 교체
	n.prev = prevNode     // 새로운 노드에 이전 노드 교체

	if prevNode != nil {
		prevNode.next = n // 이전 노드에 넥스트를 새로운 노드로 교체
	}

	if node == l.root {
		l.root = n
	}

	l.count++
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}

	n := l.root     // root 기록
	l.root = n.next // root의 다음 노드가 새로운 노드
	if l.root != nil {
		l.root.prev = nil // 새로운 루트의 이전 노드가 없어짐
	} else {
		l.tail = nil // l.root가 nil 일 상황이니, l.tail도 nil이여야 함
	}
	n.next = nil // 현재 노드의 다음 노드도 없음
	l.count--

	return n
}

func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}
	n := l.tail     // tail 기록
	l.tail = n.prev // tail의 이전 노드가 새로운 노드
	if l.tail != nil {
		l.tail.next = nil // 새로운 tail의 다음 노드가 없어짐
	} else {
		l.root = nil // l.root가 nil 일 상황이니, l.tail도 nil이여야 함
	}
	n.next = nil // 현재 노드의 다음 노드도 없음
	l.count--

	return n
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}
	if node == l.tail {
		l.PopBack()
		return
	}
	if !l.isIncluded(node) {
		return
	}

	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	node.prev = nil
	node.next = nil
	l.count--
}
