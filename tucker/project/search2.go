package main

//package main
//
//import (
//	"bufio"
//	"os"
//	"path/filepath"
//	"strings"
//)
//
//type LineInfo struct {
//	lineNo int
//	line   string
//}
//
//type FileInfo struct {
//	filename string
//	lines    []LineInfo
//}
//
//func GetFileList(path string) ([]string, error) {
//	return filepath.Glob(path)
//}
//
//func FindWordInFile(word, filename string) FileInfo {
//	f, err := os.Open(filename)
//	if err != nil {
//		println("file open error")
//	}
//	defer f.Close()
//
//	findInfo := FileInfo{filename, []LineInfo{}}
//	lineNo := 1
//
//	scanner := bufio.NewScanner(f)
//	for scanner.Scan() {
//		line := scanner.Text()
//		if strings.Contains(line, word) {
//			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
//		}
//		lineNo++
//	}
//	return findInfo
//}
//
//func FindWordInAllFiles(word, path string) []FileInfo {
//	filelist, err := GetFileList(path)
//	if err != nil {
//		println("파일을 찾을 수 없습니다 err : ", err)
//		return nil
//	}
//
//	var findInfos []FileInfo
//	for _, filename := range filelist {
//		findInfos = append(findInfos, FindWordInFile(word, filename))
//	}
//	return findInfos
//}
//
//func main() {
//	if len(os.Args) < 3 {
//		println("Usage : search2 <word> <file>")
//		return
//	}
//	word := os.Args[1]
//	files := os.Args[2:]
//	var fileInfos []FileInfo
//
//	for _, path := range files {
//		fileInfos = append(fileInfos, FindWordInAllFiles(word, path)...)
//	}
//	for _, fileInfo := range fileInfos {
//		println(fileInfo.filename)
//		for _, line := range fileInfo.lines {
//			println(line.lineNo, line.line)
//		}
//		println("--------------\n\n")
//	}
//}
