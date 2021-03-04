package main

import "os"

func main() {
	os.Stdout.Write([]byte("Out from writer\n"))
}
