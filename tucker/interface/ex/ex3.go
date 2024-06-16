package main

/* 어떤 경우에도 런타입 에러가 발생하지 않도록 CheckAndRun() 함수 수정 */

type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

func CheckAndRun(stringer Stringer) {
	//r := stringer.(Reader)
	//r.Read()

	if r, ok := stringer.(Reader); ok {
		r.Read()
	}
}
