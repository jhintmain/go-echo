package main

import (
	"bufio"
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

func GetFileListUnder(pattern string) ([]string, error) {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			_, err := filepath.Match(pattern, info.Name())
			if err == nil {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		println("파일을 찾을 수 없습니다 err : ", err)
		return files, err
	}
	return files, nil
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
	//filelist, err := GetFileList(path)
	filelist, err := GetFileListUnder(path)
	if err != nil {
		println("파일을 찾을 수 없습니다 err : ", err)
		return findInfos
	}

	// 채널 생성
	var wg sync.WaitGroup
	ch := make(chan FindInfo)

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

	for _, path := range files {
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
