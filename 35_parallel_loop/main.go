// Taken from or adapted from "The Go Programming Language" (Alan A. A. Donovan and Brian W. Kernighan.) source code at https://github.com/adonovan/gopl.io which is licensed under https://creativecommons.org/licenses/by-nc-sa/4.0/
// Parallel looping using sync.Waitgroup

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const jobCount int = 5

var wg sync.WaitGroup

func main() {
	startTime := time.Now()
	for job := 0; job <= jobCount; job++ {
		wg.Add(1)
		go work(job)
	}

	wg.Wait()
	totalDuration := time.Since(startTime)
	fmt.Printf("Total run time %v", totalDuration)
}

func work(job int) error {
	defer wg.Done()
	fmt.Printf("Job %v started\n", job)
	maxWorkTime := 7
	minWorkTime := 2
	workTime := rand.Intn(maxWorkTime-minWorkTime) + minWorkTime
	time.Sleep(time.Duration(workTime) * time.Second)
	fmt.Printf("Job %v finished in %vs\n", job, workTime)
	return nil
}
