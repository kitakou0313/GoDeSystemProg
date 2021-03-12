package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	fileOut, err := os.Create("out.txt")

	if err != nil {
		panic(err)
	}
	defer fileOut.Close()

	_, err = io.CopyN(fileOut, rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
}
