package main

import (
	"fmt"
	"time"
)

func sub1(exChan chan int) {
	time.Sleep(time.Second)
	exChan <- 1
}

func sub2(exChan chan int) {
	time.Sleep(time.Second)
	exChan <- 2
}

func main() {
	fmt.Println("Start Subs")

	chan1 := make(chan int)
	chan2 := make(chan int)
	go sub1(chan1)
	go sub2(chan2)

	for {
		select {
		case <-chan1:
			fmt.Println("Fin from sub1")
			break
		case <-chan2:
			fmt.Println("Fin From sub2")
			break
		default:
			break
		}
	}

}
