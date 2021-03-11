package main

import (
	"io"
	"os"
)

func main() {
	fileIn, err := os.Open("old")

	if err != nil {
		panic(err)
	}
	defer fileIn.Close()

	fileOut, err := os.Open("new")
	if err != nil {
		panic(err)
	}
	defer fileOut.Close()

	io.Copy(fileOut, fileIn)
}
