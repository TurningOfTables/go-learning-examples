// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/

// Spinner uses a go routine to display a spinner while computing a Fibonacci number

package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
