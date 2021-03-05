package main

import "os"

//stdoutのwriteインターフェイスを使う
func main() {
	os.Stdout.Write([]byte("Out from writer\n"))
}
