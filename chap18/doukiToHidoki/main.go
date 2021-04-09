package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func readFileAsync(filePath string, outChan chan string) {
	fileContents, _ := ioutil.ReadFile(filePath)

	outChan <- string(fileContents)
}

func readFileSync(filePath string) string {
	fileContents, _ := ioutil.ReadFile(filePath)
	return string(fileContents)
}

func main() {
	inputs1 := make(chan string)
	inputs2 := make(chan string)

	start := time.Now()
	// 処理
	go readFileAsync("a.txt", inputs1)
	go readFileAsync("b.txt", inputs2)

	for {
		select {
		case <-inputs1:
			break
		case <-inputs2:
			break
		default:
			break
		}
	}

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

	start = time.Now()
	// 処理
	_ = readFileSync("a.txt")
	_ = readFileSync("b.txt")

	end = time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
