package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Ignore Ctrl + C for 10 seconnd")

	signal.Ignore(syscall.SIGINT, syscall.SIGHUP)

	time.Sleep(10 * time.Second)
}
