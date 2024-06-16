package main

import "fmt"

// 연습문제 379 page

/* 다음 공란을 채우세오 */

func sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}
func main() {
	fmt.Println("1~5까지 합 ", sum(1, 2, 3, 4, 5))
}
