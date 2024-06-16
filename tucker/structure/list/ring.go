package main

import (
	"container/ring"
	"fmt"
)

/**
	링 (ring)
    - 맨뒤와 맨앞 요소가 서로 연결괸 자료구조 ( 환형 리스트 )
    - 시작도 없고 끝도 없다. 현재 위치만 존재 할 뿐
	- 사용하는 case : 저장갯수가 고정되고, 오래된 요소는 지워도 되는 경우 (ex) 실행취소, replay)
*/

func main() {
	r := ring.New(5)

	n := r.Len()

	// ring의 각 요소에 값 대입
	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	// 순회
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	// 역순회
	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}

}
