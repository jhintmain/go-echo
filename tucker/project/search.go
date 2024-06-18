package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func printAllFiles(files []string, word string) {
	for _, path := range files {
		filelist, err := getFileList(path)
		if err != nil {
			fmt.Println("파일을 찾을 수 없습니다 err : ", err)
			return
		}

		for _, name := range filelist {
			printFile(name, word)
			//fmt.Println(name)
		}
	}
}
func printFile(filename, word string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file err : ", err)
		return
	}
	defer file.Close()
	fmt.Println(filename)
	fmt.Println("----------------------------")
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		row++
		if strings.Contains(scanner.Text(), word) {
			fmt.Println(row, scanner.Text())
		}
	}
	fmt.Println("----------------------------\n")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(os.Args)
		panic("찾을 단어, 경로를 입력해주세오")
	}

	word := os.Args[1]
	files := os.Args[2:]

	fmt.Println("찾으려는 단어 : ", word)
	printAllFiles(files, word)

}
