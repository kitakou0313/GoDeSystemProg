package main

import (
	"fmt"
	"sync"
	"time"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func main() {

	mutex := new(sync.Mutex)

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(mutex))
		}()
	}
	time.Sleep(2 * time.Second)
}
