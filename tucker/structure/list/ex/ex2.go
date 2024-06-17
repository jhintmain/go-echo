package main

import "fmt"

// 연습 문제 : 404 page

/* 다음은 ID를 키로 학생데이터를 값으로 갖는 맵을 만든뒤 38번 에 해당하는 학생 데이터를 출력하는 프로그램입니다 빈칸을 채우세오 */

type Student struct {
	Name  string
	Score int
}

func main() {
	m := make(map[int]Student)

	m[5] = Student{"최번개", 67}
	m[17] = Student{"송하나", 89}
	m[23] = Student{"화랑", 97}
	m[38] = Student{"백두산", 78}
	m[45] = Student{"김갑환", 56}

	fmt.Println("38번 ", m[38])
}
