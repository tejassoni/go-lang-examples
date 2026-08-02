package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	defer func() { wg.Done() }()
	fmt.Println("Starting of worker function...! Task : ", i)
	fmt.Println("Ending of worker function...! Task : ", i)
}

func main() {
	now := time.Now()
	fmt.Println("Starting of main program...!")
	// create a wait group
	var wg sync.WaitGroup
	maxWorkers := 5 // number of workers to run concurrently
	// start 5 go routines
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)         // increment the wait group counter
		go worker(i, &wg) // start a go routine to run the worker function
	}

	wg.Wait() // wait for all go routines to finish before exiting the main function
	fmt.Println("Ending of main program...!")
	fmt.Println("Total time taken: ", time.Since(now))
}

// output:
// Starting of main program...!
// Starting of worker function...! Task :  4
// Ending of worker function...! Task :  4
// Starting of worker function...! Task :  0
// Ending of worker function...! Task :  0
// Starting of worker function...! Task :  3
// Ending of worker function...! Task :  3
// Starting of worker function...! Task :  1
// Ending of worker function...! Task :  1
// Starting of worker function...! Task :  2
// Ending of worker function...! Task :  2
// Ending of main program...!
// Total time taken:  138.521µs
