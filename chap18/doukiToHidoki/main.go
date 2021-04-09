package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

func readFileAsync(filePath string, outChan chan string, wg *sync.WaitGroup) {
	_, _ = ioutil.ReadFile(filePath)

	// outChan <- string(fileContents)

	wg.Done()
}

func readFileSync(filePath string) string {
	fileContents, _ := ioutil.ReadFile(filePath)
	return string(fileContents)
}

func main() {
	inputs1 := make(chan string)
	inputs2 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()
	// 処理
	go readFileAsync("a.txt", inputs1, &wg)
	go readFileAsync("b.txt", inputs2, &wg)

	wg.Wait()

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

	start = time.Now()
	// 処理
	_ = readFileSync("a.txt")
	_ = readFileSync("b.txt")

	end = time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())

}
