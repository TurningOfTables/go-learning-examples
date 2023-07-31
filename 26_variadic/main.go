// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/

// Uses a variadic function (one that can accept any number of arguments using ...T) to sum any number of ints

package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)
	totalTwo := sum(29, 5)
	fmt.Println(totalTwo)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}

	return total
}
