package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(path string) (promise chan string) {

	go func() {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		} else {
			promise <- string(content)
		}
	}()
	return
}

func printFunc(futureSource chan string) (promise chan []string) {
	go func() {
		var result []string

		for _, line := range strings.Split(<-futureSource, "\n") {
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}

		promise <- result
	}()
	return
}

func main() {
	futureSource := readFile("./main.go")
	futureFuncs := printFunc(futureSource)

	fmt.Println(strings.Join(<-futureFuncs, "\n"))
}
