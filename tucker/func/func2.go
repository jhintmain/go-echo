package main

/* func type 변수
-  이 예제로는 왜 func type 변수를 사용하는지 이해가 안됨...
*/

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	operator := getOperator("*")
	println(operator(3, 4))

}
