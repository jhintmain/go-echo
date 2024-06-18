package main

import (
	"fmt"
	"sync"
	"time"
)

/**
채널 (channel) > 메세지 queue
- 고 루틴간 메세지 효율적 전달
- mutex없이 동시성 프로그래밍 가능
- context를 이용하여 작업 요청시 작업시간,취소, 추가 데이터등 지정 가능

select문  > 여러 channel 동시에 대기
- case로 채널 구분하며, 한번 처리되면 break되므로 for문 안에 써주는것이 일반
*/

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // 채널 생성 (크기 0)
	quit := make(chan bool)

	wg.Add(3)
	go square(&wg, ch)
	ch <- 9

	go squareSelect(&wg, ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- true

	go squareTick(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- 9999
	}

	close(ch)   // 채널 썻으면 닫아줘라
	close(quit) // 채널 썻으면 닫아줘라
	wg.Wait()

}

func square(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	n := <-ch // 채널에서 데이터 추출
	time.Sleep(time.Second)
	fmt.Printf("square : %d\n", n*n)
}

func squareSelect(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case <-quit:
			wg.Done()
			return
		case n := <-ch:
			time.Sleep(time.Second)
			fmt.Printf("squareSelect : %d\n", n*n)
		}
	}
}

func squareTick(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(2 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Printf("Tick\n")
		case <-terminate:
			fmt.Printf("Terminate\n")
			wg.Done()
			return
		case n := <-ch:
			time.Sleep(time.Second)
			fmt.Printf("squareTick : %d ", n*n)
		}
	}
}
