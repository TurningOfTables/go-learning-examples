// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/

// Basic webserver that returns static text

package main

import (
	"fmt"
	"log"
	"net/http"
)

const URL = "localhost:8000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "hi") })
	log.Print("Listening")
	err := http.ListenAndServe(URL, nil)
	if err != nil {
		log.Fatal(err)
	}
}
