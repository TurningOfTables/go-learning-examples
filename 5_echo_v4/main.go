// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/

// prints command line arguments and their index

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d\t%s", i, arg)
	}
}
