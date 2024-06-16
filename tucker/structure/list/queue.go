package main

import (
	"container/list"
	"fmt"
)

// list로  queue 만들기

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val) // list 뒤에 요소 추가 ( 0 > 0 1 > 0 1 2 )
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()

	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
