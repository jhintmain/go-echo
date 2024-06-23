package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//type LineInfo struct {
//	row  int
//	line string
//}
//type FileInfo struct {
//	filename string
//	lines    []LineInfo
//}

func getFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func printAllFiles(files []string, word string) {
	//var wg sync.WaitGroup
	ch := make(chan FileInfo)
	revCtn := 0
	cnt := len(files)

	for _, path := range files {
		filelist, err := getFileList(path)
		if err != nil {
			fmt.Println("파일을 찾을 수 없습니다 err : ", err)
			return
		}

		for _, name := range filelist {
			go printFile(name, word, ch)
			//fmt.Println(name)
		}
	}
	for fileinfo := range ch {
		fmt.Println(fileinfo)
		revCtn++
		if revCtn == cnt {
			break
		}
	}
}
func printFile(filename, word string, ch chan FileInfo) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file err : ", err)
		return
	}
	defer file.Close()

	lineinfo := []LineInfo{}
	fileinfo := FileInfo{filename, lineinfo}

	fmt.Println(filename)
	fmt.Println("----------------------------")
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		row++
		if strings.Contains(scanner.Text(), word) {
			lineinfo = append(fileinfo.lines, LineInfo{row, scanner.Text()})
			fmt.Println(row, scanner.Text())
		}
	}
	ch <- fileinfo
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
