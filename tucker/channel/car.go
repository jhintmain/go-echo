package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	painCh := make(chan *Car)

	fmt.Println("Start Factory")

	wg.Add(3)

	go MakeBody(tireCh)
	go InstallTire(tireCh, painCh)
	go PainCar(painCh)

	wg.Wait()
}

func MakeBody(tireCh chan *Car) {
	// 생성자
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{Body: "Sport Car"}
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}
func InstallTire(tireCh chan *Car, painCh chan *Car) {
	// 소비자
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter Tire"
		painCh <- car
	}
	close(painCh)
	wg.Done()
}

func PainCar(painCh chan *Car) {
	// 소비자
	for car := range painCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Since(startTime)
		fmt.Printf("%.2f Complete Car : %s / %s / %s\n", duration.Seconds(), car.Tire, car.Color, car.Body)
	}
	wg.Done()
}
