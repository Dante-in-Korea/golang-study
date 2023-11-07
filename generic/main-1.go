package main

import (
	"fmt"
	"strconv"
)

// interface and 타입제한자 차이

type Stringer interface { //인터페이스
	~int8 | ~int16 | ~int32 | ~int64 | ~int // 이렇게도 사용이 가능하지만 조건이 AND임
	String() string
}

func PrintMin[T Stringer](a, b T) {
	if a < b {
		fmt.Println(a.String())
	} else {
		fmt.Println(b.String())
	}
}

type MyInt int //MyInt는 int의 별칠 타입

func (m MyInt) String() string {
	return strconv.Itoa(int(m))
}

//type Integer interface { //타입제한자(여기중에 포함된)
//	~int8 | ~int16 | ~int32 | ~int64 | ~int
//}

//func (m MyString) String() string {
//	return m.name
//}

//func Print1(a Stringer) { //interface
//	fmt.Println(a.String())
//}
//func Print2[T Stringer](a T) { //generic
//	fmt.Println(a.String())
//}

//type MyString struct {
//	name string
//}

func main() {
	var m1 MyInt = 10
	var m2 MyInt = 20
	PrintMin(m1, m2)
}
