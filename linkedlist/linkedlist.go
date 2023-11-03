package golang

type Node[T any] struct {
	next  *Node[T] //Edges 다음 노드를 가리키는 포인터
	Value T
}

type LinkedList[T any] struct {
	root *Node[T] //Fist node pointer
	tail *Node[T] //Last node pointer

	count int
}

func (l *LinkedList[T]) PushBack(value T) { // 맨 뒤에 노드를 추가 방식
	node := &Node[T]{
		Value: value,
	}
	l.count++ // push 될 때 마다 count
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node

}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++ // push 될 때 마다 count
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	node.next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int { //Count Linked list Length O(n)
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

func (l *LinkedList[T]) Count2() int { // O(1) / push가 되거나 제거할때 길이를 수정해야함
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] { // 특정 노드 반환 O(n)
	if idx >= l.Count2() {
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

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isIncluded(node) { // 추가가 될 때 포함된 노드인지 아닌지 체크
		return
	}

	newNode := &Node[T]{
		Value: value,
	}

	node.next, newNode.next = newNode, node.next
	/* it's the same code
	origNext := node.next
	node.next = newNode
	newNode.next = origNext
	*/
	l.count++
}
func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) { // Prev 노드는 어떻게 찾지?
	if node == l.root {
		l.PushFront(value)
		return
	}

	prevNode := l.findPrevNode(node) // 이전 노드 찾는건 이미 처음부터 node의 존재 여부를 찾기 때문에 검증할 필요 없다.
	if prevNode == nil {
		return
	}
	newNode := &Node[T]{
		Value: value,
	}
	prevNode.next, newNode.next = newNode, node
	l.count++
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool { // 노드가 있는지 없는지 체크
	inner := l.root
	for ; inner != nil; inner = inner.next { // root node부터 끝까지 돌림
		if inner == node {
			return true
		}
	}
	return false
}

func (l LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node { // 노드의 다음이 노드면 inner가 이전 노드가 된다
			return inner //
		}
	}
	return nil
}

func (l *LinkedList[T]) PopFront() {
	if l.root == nil {
		return
	}

	l.root.next, l.root = nil, l.root.next
	//l.root = l.root.next
	//l.root.next = nil
	if l.root == nil { //root가 nil이면 tail도 없으니 tail도 nil 처리
		l.tail = nil
	}
	l.count-- // decrease count after popfront
}

func (l *LinkedList[T]) Remove(node *Node[T]) { // 이전 노드 찾아서 next를 지우고자 할 노드의 next로 지정
	if node == l.root {
		l.PopFront()
		return
	}

	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}
	prev.next = node.next
	node.next = nil

	if node == l.tail { // tail 갱신
		l.tail = prev
	}
	l.count--
}
