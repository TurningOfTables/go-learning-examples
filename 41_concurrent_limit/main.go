// Example using a buffered channel to limit the number of concurrent 'someTasks' that can run at once

package main

import (
	"fmt"
	"time"
)

var tokens = make(chan struct{}, 1)

func main() {
	for x := 1; x <= 10; x++ {
		tokens <- struct{}{}
		go someTask(x)
	}
}

func someTask(x int) {

	fmt.Printf("Starting task %v\n", x)
	time.Sleep(time.Second * 2)
	fmt.Printf("Ending task %v\n", x)
	<-tokens
}
