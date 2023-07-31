// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/

// Finds the base name for a given file path and extension using the strings library

package main

import (
	"strings"
)

var s = "a/b/c.go.blah"

func main() {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:] // Remove everything up to and including the slash. Since LastIndex() returns -1 if no slash is found, this leaves the string intact.
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
}
