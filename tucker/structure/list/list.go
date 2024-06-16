package main

/**
자료구조
1. 리스트 : 불연속된 메모리를 데이터에 저장
- 요소들을 포인터로 연결 ( 링크드 리스트 )
- 사용되면 좋은 case : 인덱스를 사용한 접근이 거의없고, 삽입과 삭제가 빈번한 경우
*/

import (
	"container/list"
	"fmt"
)

// list의 자료구조
type Element struct {
	Value interface{} // 데이터 저장 필드
	Next  *Element    // 다음요소 주소 저장 필드
	Prev  *Element    // 이전요소 주소 저장 필드
}

// Next와 Prev가 존재하여 양방향으로접근가능 (양방향 리스트)
func main() {
	v := list.New()      // list생성
	e4 := v.PushBack(4)  // list 뒤에 요소 추가  : 4
	e1 := v.PushFront(1) // list 앞에 요소 추가	: 1 4

	v.InsertBefore(3, e4) // e4 요소 앞에 요소 삽입 : 1 3 4
	v.InsertAfter(2, e1)  // e1 요소 뒤에 요소 삽입 : 1 2 3 4

	// .Front() : 리스트의 가장 첫번째 요소 return
	// .Next() : 현재요소의 다음 요소 return
	for e := v.Front(); e != nil; e = e.Next() { // 각 요소 순회
		fmt.Print(e.Value, " ")
	}

	fmt.Println()

	// .Back() : 리스트의 가장 마지막 요소 returnb
	// .Prev() : 현재요소의 전 요소 return
	for e := v.Back(); e != nil; e = e.Prev() { // 각 요소 역순회
		fmt.Print(e.Value, " ")
	}

}
