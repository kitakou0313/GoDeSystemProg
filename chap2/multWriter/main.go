package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)

	}

	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "multi write example")
}
