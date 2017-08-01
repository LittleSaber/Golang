package main

import (
	"fmt"
)

func main() {
	res := gcd(12, 18)
	fmt.Println(res)
	fibres := fib(4)
	fmt.Println(fibres)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
