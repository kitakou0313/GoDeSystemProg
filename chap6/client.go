package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	sendMsg := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}

	current := 0

	var conn net.Conn = nil

	for {
		var err error

		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		req, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMsg[current]))
		if err != nil {
			panic(err)
		}

		err = req.Write(conn)
		if err != nil {
			panic(err)
		}

		response, err := http.ReadResponse(bufio.NewReader(conn), req)

		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		dump, err := httputil.DumpResponse(response, true)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))

		current++
		if current == len(sendMsg) {
			break
		}

	}
}
