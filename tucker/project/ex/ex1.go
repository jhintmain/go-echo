package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func GetFileList2(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		println("파일을 찾을 수 없습니다 err : ", err)
	}
	return files
}

func FindWordInFile(word, filename string, ch chan FindInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	findInfo := FindInfo{filename, []LineInfo{}}
	f, err := os.Open(filename)
	if err != nil {
		ch <- findInfo
		return
	}
	defer f.Close()

	lineNo := 1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo
}

func FindWordInAllFiles(word, path string) []FindInfo {
	var findInfos []FindInfo
	filelist, err := GetFileList(path)
	if err != nil {
		println("파일을 찾을 수 없습니다 err : ", err)
		return findInfos
	}

	// 채널 생성
	var wg sync.WaitGroup
	ch := make(chan FindInfo)
	//cnt := len(filelist)
	//revCtn := 0

	for _, filename := range filelist {
		wg.Add(1)
		go FindWordInFile(word, filename, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo)
		//revCtn++
		//if revCtn == cnt {
		//	break
		//}
	}
	return findInfos
}

func main() {
	if len(os.Args) < 3 {
		println("Usage : search2 <word> <file>")
		return
	}
	word := os.Args[1]
	files := os.Args[2:]
	var fileInfos []FindInfo

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return

	}
	for _, path := range files {
		fmt.Println(GetFileList2(dir))

		fileInfos = append(fileInfos, FindWordInAllFiles(word, path)...)
	}
	for _, fileInfo := range fileInfos {
		println(fileInfo.filename)
		for _, line := range fileInfo.lines {
			println(line.lineNo, line.line)
		}
		println("--------------\n\n")
	}
}
