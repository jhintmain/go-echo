package main

import "fmt"

/**
- 포함된 interface
- 빈 interface
- interface 기본값 : nil
- interface > 원래 구조체 type 변경 : .()
	: 런타임 에러가 발생 가능성이 높ㅇ므로 타입변환 가능 여부 체크한후 메소드 사용하자
*/

// 1. interface를 포함하는 interface
type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

// Read(), Write(), Close()를 모두 포함하는 interface
type ReadWriter interface {
	Reader
	Writer
}

// 2. 빈 interface
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("int : %d\n", int(t))
	case float64:
		fmt.Printf("float64 : %f\n", float64(t))
	case string:
		fmt.Printf("string : %s\n", string(t))
	default:
		fmt.Printf("unknown type %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})
}
