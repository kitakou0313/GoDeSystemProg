package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("Start sub")

	ctx, calcel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("Sub is finished")
		calcel()
	}()

	<-ctx.Done()
	fmt.Println("All tasks are finished")
}
