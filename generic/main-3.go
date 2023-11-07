package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map[F, T any](s []F, f func(F) T) []T { //generic type이 생기면서 여러 타입으로 반환을 할 수 있게됨
	rst := make([]T, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}

func main() {
	doubled := Map([]int{1, 2, 3}, func(v int) int { // << 이게 lambda function
		return v * 2
	})
	fmt.Println(doubled)

	uppered := Map([]string{"hello", "world", "abc"}, func(v string) string {
		return strings.ToUpper(v)
	})
	fmt.Println(uppered)

	toString := Map([]int{1, 2, 3}, func(v int) string {
		return "str" + strconv.Itoa(v)
	})
	fmt.Println(toString)
}
