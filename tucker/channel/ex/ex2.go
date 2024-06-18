package ex

import (
	"fmt"
	"sync"
)

/* 다음코드에서 발생하는 데드락 문제를 올바르게 수정한 보기를 고르시오 */

var wg sync.WaitGroup

func printDeadLockSquare(ch chan int) {
	for n := range ch {
		fmt.Println(n * n)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	ch := make(chan int)
	go printDeadLockSquare(ch)
	for i := 0; i < 5; i++ {
		ch <- i * 2
	}
	close(ch) // 수정(추가)
	wg.Wait()
}
