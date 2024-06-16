package main

import "fmt"

/* 다음 예쩨의 출력 결과를 쓰세오 */

type OpFunc func(int, int) int

func Process(a, b int, op OpFunc) {
	fmt.Println("Result : ", op(a, b))
}

func main() {
	op := func(a, b int) int {
		return a * b
	}
	Process(3, 4, op) // Result : 12
}
