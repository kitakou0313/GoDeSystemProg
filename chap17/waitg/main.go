package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex, wg *sync.WaitGroup) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	fmt.Printf("ID: %d\n", id)
	wg.Done()
	return id
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go generateId(&mutex, &wg)
	}

	wg.Wait()
}
