package main

import (
	"fmt"
	"golang.org/x/exp/constraints" //   go mod download golang.org/x/exp
	// ~ 틸트연산자에서 ~int 까지 있는 별칭까지 포함한다
)

//func print[T any](a T) { // generic type 정의 - a의 type은 T
//	fmt.Println(a)
//}

//func min(a, b interface{}) interface{} { // generic type이 없기전가진 빈 interface를 쓰긴했지만
//	// 왜 generic type이 나왔는지 알 수 있음(강타입이다 보니 여러타입을 제어하기 어려웠음)
//	// ./main.go:11:5: invalid operation: a < b (operator < not defined on interface)
//	if a < b {
//		return a
//	}
//	return b
//}

//type Integer interface {
//	int | int8 | int16 | int32 | int64 // 타입제한자
//}
//
//type Float interface {
//	float32 | float64 // 타입제한자
//}
//
//type Numeric interface {
//	Integer | Float
//}

// 이미 constraints 패키지에 구현되어있음
// func min[T any](a, b T) T {
// func min[T int | int16 | float32 | float64 | int64](a, b T) T { //대소비교가 가능한 예약어을 이렇게 선언 가능
// func min[T Integer | Float](a, b T) T {
func min[T constraints.Ordered](a, b T) T { // 미리 정의된 것을 사용 가능
	//# command-line-arguments
	//./main.go:19:5: invalid operation: a < b (type parameter T is not comparable with <)
	// 모든 타입이 올 수 있지만 struct 같은 것들은 대소비교가 안되기 때문에 generic에서 <, > 연산자를 사용할 수 없음
	if a < b {
		return a
	}
	return b
}

func main() {
	var a int = 10
	var b int = 20
	fmt.Println(min(a, b))

	var c int16 = 10
	var d int16 = 20
	fmt.Println(min(c, d))

	var e float32 = 3.14
	var f float32 = 3.19
	fmt.Println(min(e, f))

	var g string = "Hello"
	var h string = "World"
	fmt.Println(min(g, h))
}
