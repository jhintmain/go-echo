package main

import "fmt"

/**
interface는 메서드의 집합이다.
덕타이핑 지원 : 타입 선언시 인터페이스 구현 여부를 명시적으로 선언할 필요 x, interface에서 정의한 메소드 포험 여부 만으로 결정하는 방식
 - 장점 :  서비스 사용자 중심의 코딩을 할 수 있다.
~er 을 붙여서 interface 이름을 짓는 것을 권장한다
사용 이유 : 구체화된 객체가 아닌 interface만으로 메서드를 호출 할 수 있기 때문에 > 변경에 유연한 코드 작성
*/

type Sample interface {
	String() string
	//String(int) string // error : already declared
	//_(x int)           // error : missing method name
}

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Age)
}

func main() {
	student := &Student{Name: "홍길동", Age: 16}
	var stringer Stringer // interface 변수 선언

	stringer = student // interface 변수에 구조체 대입 > student가 String()을 구현하고 있으므로 stringer에 대입 가능

	fmt.Println(stringer.String())
}
