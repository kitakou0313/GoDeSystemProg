package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("PID :%d\n", os.Getpid())
	fmt.Printf("Parents PID :%d\n", os.Getppid())
}
