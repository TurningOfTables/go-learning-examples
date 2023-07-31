// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/

// Basic example of go channels
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go calculate(5, 6, c)
	result := <-c

	fmt.Println(result)
}

func calculate(x, y int, c chan int) {
	sum := x + y
	c <- sum
}
