package ex

import (
	"fmt"
	"time"
)

// 연습 문제 : 467 page

/* 다음코드에서 발생하는 데드락 문제를 올바르게 수정한 보기를 고르시오 */

func printSquare(ch chan int) {
	n := <-ch
	fmt.Println(n * n)
}

func main() {
	ch := make(chan int)
	//printSquare(ch)
	go printSquare(ch) // 수정

	ch <- 5
	time.Sleep(5 * time.Second)

}
