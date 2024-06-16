package main

type Product struct {
	Name        string
	Price       int
	ReviewScore float64
}
type Padding struct {
	A int8
	G int8
	F float32
	D uint16
	C float64
	B int
	E int
}

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	return &Actor{name, hp, speed}
}

func main() {

}
