package main

import (
	"fmt"
	"os"
)

// defer : 함수가 종료되기 직전에 실행되는 문장 (지연실행)
// defer : stack 구조로 실행됨

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("file create error")
		return
	}
	defer fmt.Println("반드시 호출 됩니다")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다")

	fmt.Println("파일에 Hellow World를 씁니다")
	fmt.Fprintln(f, "Hello World")
}
