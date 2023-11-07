package main

// generic type의 제한사항
// method, functions 확인
// go 1.19에서는 method도 type parameters를 쓸 수 있게 하고 있다는데
// 1.20에서 안되는거 보니 지원이 안되고있다.

// Data Structure에서 쓰기 좋음
// ex) Linked List, Heap, Tree, Map, Tree, Sparse Map - interface{}
// 다양한 타입들이 많이 오기에..
// generic을 많이 쓰면 코드 가독성이 떨어지기에 나~~중에 코드 리펙토링하면서 변환해도 괜찮다.

type Node struct { // type 선언할 떄는 type parameter를 넣을 수 있음
	val  int
	next *Node
}

func Print[T any](a T) { // 함수는 type parameters를 넣을수 있음

}

// func (n *Node) Push[T any](a T) { // method must have no type parameters
//
//										method에서는 type parameters를 받지 못한다
//	}
func main() {

}
