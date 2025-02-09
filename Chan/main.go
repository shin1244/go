package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	wg.Add(3)
	var TireCh = make(chan *Car)
	var paintCh = make(chan *Car)

	fmt.Println("Start!")

	go MakeBody(TireCh)
	go InstallTire(TireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
}

func square1() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}

func square2(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println(n * n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func square3(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminate")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println(n)
			time.Sleep(time.Second)
		}
	}
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(time.Second * 10)

	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports Car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "winter Tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.color = "red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.color)
	}
	wg.Done()
}
