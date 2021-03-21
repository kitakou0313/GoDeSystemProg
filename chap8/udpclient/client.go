package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Sendig to server")
	_, err = conn.Write([]byte("Hello, from, Client"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received: %s\n", string(buffer[:length]))
}
