package main

// 연습 문제 : 361 page

/* 다음 예쩨에서 컴파일 에러가 발생하는 이유 */

type ReadWriter interface {
	Read()
	Write()
}

type File struct {
}

func (f *File) Read() {

}

func ReadWrite(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	f := &File{}
	ReadWrite(f) // answer : File 구조체는 Read()만 구현하고 있으므로 ReadWriter 의 interface가 아님
}
