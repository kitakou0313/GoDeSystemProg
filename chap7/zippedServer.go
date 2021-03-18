package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func isGZipAcceptable(req *http.Request) bool {
	return strings.Index(strings.Join(req.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSession(conn net.Conn) {
	fmt.Println("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		req, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Time out")
			} else if err == io.EOF {
				break
			}
		}
		panic(err)

		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))

		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}
		if isGZipAcceptable(req) {
			content := "Hello, world, zipped"

			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()

			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "Hello, world"
			response.Body = ioutil.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is now running")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
