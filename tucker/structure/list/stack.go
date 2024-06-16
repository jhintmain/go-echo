package main

/**
list로 stack 구현하기
- Queue는 list로
- Stack은 array로 구현함 : 요소 추가 삭제가 항상 맨뒤에서 발생하기때문에. 성능 손해없음
*/
import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func (s *Stack) Push(v interface{}) {
	s.v.PushBack(v) // 0 > 0  1 > 0 1 2
}
func (s *Stack) Pop() interface{} {
	last := s.v.Back() // 가장 마지막 요소

	if last != nil {
		return s.v.Remove(last)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func main() {
	stack := NewStack()

	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	v := stack.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = stack.Pop()
	}
}
