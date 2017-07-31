package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 统计相同行个数量
func main() {
	counts := make(map[string]int)
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println("%d\t%s\n", n, line)
		}
	}
}
