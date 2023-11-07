package main

import "fmt"
// 기본적인 generic 사용법
type Node[T any] struct { // 모든 타입이 사용이 가능한 Node
	val  T
	next *Node[T]
}

func NewNode[T any](v T) *Node[T] { //generic function
	return &Node[T]{val: v}
}

func (n *Node[T]) Push(v T) *Node[T] {
	node := NewNode(v)
	n.next = node
	return node
}
func main() {
	node1 := NewNode(1) //*Node[int]
	node1.Push(2).Push(3).Push(4)

	for node1 != nil {
		fmt.Print(node1.val, " - ")
		node1 = node1.next
	}

	node2 := NewNode("Hi") // *Node[string]
	node2.Push("Hello").Push("How are ypu?")
	for node2 != nil {
		fmt.Print(node2.val, " - ")
		node2 = node2.next
	}
}
