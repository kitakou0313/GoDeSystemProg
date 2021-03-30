package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start sub")

	go func() {
		fmt.Println("Sub is running")
		time.Sleep(time.Second)
		fmt.Println("Sub is finished")
	}()

	time.Sleep(2 * time.Second)
}
