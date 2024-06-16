package main

import "fmt"

/* 함수 리터럴 : 이름없는 함수 , 변수로만 호출 (익명함수, 람다함수)
- 함수 리터럴은 필요한 변수를 내부 상태로 가질 수 있음 (캡처)
- 캡처 : 외부 변수를 내부로 가져와서 사용하는 것. 값 복사가 아닌 참조형태(포인터같음)
*/

type opFunc func(int, int) int

func getOperatorLiteral(op string) opFunc {
	switch op {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	case "*":
		return func(a, b int) int {
			return a * b
		}
	default:
		return nil
	}
}

func main() {
	// ##### 1
	//operatorLiteral := getOperatorLiteral("+")
	//println(operatorLiteral(3, 4))

	// ##### 2
	//i := 0
	//f := func() {
	//	// 외부 변수 i를 내부로 capture(캡처)했다
	//	i += 10 //  i는 외부에서 선언되 변수지만 함수 리터럴 내부에서 사용되어 내부 상태로 저장됨. (포인터 개념 즉, 변경되면 원래 값도 같이 변경된다는것. 매우 중요)
	//}
	//i++
	//f()
	//println(i)

	// ##### 3 현재 같은 결과가나옴.... 책에선 다르게 나오는데 go 컴파일러가 변경된듯함?????
	CaptureLoop()
	CaptureLoop2()
}

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("valueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i) // 캡처된 i의 값을 출력
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("valueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v) // 캡처된 v의 값을 출력
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}
