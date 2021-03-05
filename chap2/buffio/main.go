package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer

	buffer.Write([]byte("test"))
	buffer.Write([]byte("siken"))
	fmt.Println(buffer.String())
}
