package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	var b int = 15
	var a int

	// for a := 0; a < 10; a++ {
	// 	fmt.Printf("a的值为：%d\n", a)
	// }
	fmt.Println(a)
	for a < b {
		a++
		fmt.Printf("a的值为：%d\n", a)
	}
}
