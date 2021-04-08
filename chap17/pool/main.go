package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	pool.Put("Manualy added: 1")
	pool.Put("Manualy added: 2")

	runtime.GC()

	fmt.Println(pool.Get())

}
