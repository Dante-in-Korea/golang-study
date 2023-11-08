package doublelinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)
	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(3)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.GetAt(2).Value)
	assert.Nil(t, l.GetAt(3))
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)
	l.PushFront(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(3)
	assert.NotNil(t, l.root)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.Count())

	l.PushFront(4)
	assert.NotNil(t, l.root)
	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 4, l.Count())

}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)     // 2
	l.InsertAfter(node, 4) // 1 -> 4 -> 2 -> 3
	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.GetAt(2).Value) //elemental -> 4
	assert.Equal(t, 3, l.Back().Value)   // Last elemental
	assert.Equal(t, 4, node.next.Value)
	assert.Equal(t, 3, node.next.next.Value)
	assert.Equal(t, 4, node.next.next.prev.Value)

	l.InsertAfter(l.Back(), 10) // panic: runtime error: invalid memory address or nil pointer dereference
	// tail의 next도 추가 할 수 있도록 검증 필요
	assert.Equal(t, 10, l.Back().Value)
	tempNode := &Node[int]{
		Value: 100,
	}
	l.InsertAfter(tempNode, 200) // 원래 있던 노드에 추가하는게 아님
	assert.Equal(t, 5, l.Count())
}

func TestInsertBefore(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)                   // 2
	l.InsertBefore(node, 4)              // 1 -> 4 -> 2 -> 3
	assert.Equal(t, 4, l.Count())        // the whole length
	assert.Equal(t, 4, l.GetAt(1).Value) //elemental -> 4
	assert.Equal(t, 2, l.GetAt(2).Value) // Last elemental
	assert.Equal(t, 3, l.Back().Value)   // Last elemental

	l.InsertBefore(l.Front(), 10) // panic: runtime error: invalid memory address or nil pointer dereference
	assert.Equal(t, 10, l.Front().Value)
}

func TestPopFront(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	n := l.PopFront()
	assert.Equal(t, 1, n.Value)
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.PopFront()
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.PopFront()
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestRemove(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Remove(l.GetAt(1)) // remove 2

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.Remove(l.GetAt(0)) // remove 1

	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.Remove(&Node[int]{
		Value: 100,
	})
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
}

func TestReverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverse()

	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(1).Value)
}