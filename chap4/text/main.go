package main

import (
	"bufio"
	"fmt"
	"strings"
)

var sourceStr = `1 Line
2 Line
3 Lile`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(sourceStr))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
