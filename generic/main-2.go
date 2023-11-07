package main

import (
	"fmt"
	"hash/fnv"
)

type ComparableHasher interface {
	comparable // ==, !=
	Hash() uint32
}

type MyString string //comparable 함

func (s MyString) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func Equal[T ComparableHasher](a, b T) bool {
	if a == b {
		return true
	}
	return a.Hash() == b.Hash()
}

func main() {
	var s1 MyString = "Hello"
	var s2 MyString = "World"

	fmt.Println(Equal(s1, s2))
}
