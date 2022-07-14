package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, n := range os.Args[1:] {
		nbr, _ := strconv.Atoi(n)
		fmt.Println(fib(nbr))
	}
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
