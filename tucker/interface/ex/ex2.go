package main

/* 다음 패키지의 OurDB 구조체의 모든 공개된 메서드를 이용하는 인터페이스를 만들어보세요 */

type OurDB struct {
	Name string
}

func (db *OurDB) GetData() string {

}
func (db *OurDB) WriteData(data string) {

}
func (db *OurDB) Close() error {

}

type DB interface {
	GetData() string
	WriteData(data string)
	Close() error
}
