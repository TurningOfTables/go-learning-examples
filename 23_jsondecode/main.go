// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/

// Makes a request to the Postman Echo service, decodes and prints the headers in the response

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const requestURL = "https://postman-echo.com/get?foo1=bar1&foo2=bar2"

type Response struct {
	Args    map[string]string `json:"args"`
	Headers struct {
		Host string `json:"host"`
	} `json:"headers"`
}

func main() {
	var output Response
	resp, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	json.Unmarshal(b, &output)

	resp.Body.Close()
	fmt.Println(output.Headers.Host)
}
