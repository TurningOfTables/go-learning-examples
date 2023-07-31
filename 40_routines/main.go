// Experiment with wait groups, channels and go routines to create a fake 'mining' process

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mineResult = make(chan int)
var wg sync.WaitGroup

const (
	MININGMACHINES   = 4
	CYCLESPERMACHINE = 2
)

func main() {

	for g := 0; g < 4; g++ {
		wg.Add(1)
		go mine(g)
	}
	go closer()

	for range mineResult {
		fmt.Println("Mining operation complete!")
	}
}

func mine(g int) {
	defer wg.Done()
	for x := 0; x < 2; x++ {
		fmt.Printf("Mining machine %v cycle %v started\n", g, x)
		mineTime := rand.Intn(25)
		for t := 0; t < mineTime; t++ {
			time.Sleep(time.Second)
		}
		fmt.Printf("Mining machine %v cycle %v finished\n", g, x)
	}
	mineResult <- 1
}

func closer() {
	go func() {
		wg.Wait()
		close(mineResult)
	}()
}
