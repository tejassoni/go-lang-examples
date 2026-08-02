package main

import (
	"fmt"
	"sync"
	"time"
)

// channel is unidirectional, it can only be used to send data
func sendWorker(sendChan chan<- int) {
	defer close(sendChan)
	for i := 0; i < 5; i++ {
		fmt.Println("Writing to channel...!", i)
		sendChan <- i
	}
	fmt.Println("Ending of sendWorker function...!")
}

// channel is unidirectional, it can only be used to receive data
func receiveWorker(receiveChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for readValue := range receiveChan {
		fmt.Println("Reading from channel...!", readValue)
	}
	fmt.Println("Ending of receiveWorker function...!")
}

func main() {
	now := time.Now()
	fmt.Println("Starting of main program...!")
	var wg sync.WaitGroup
	// create a bidirectional channel
	ch := make(chan int)
	// start a go routine to write to channel
	go sendWorker(ch)
	// fmt.Println("Reading from channel...!", value)
	wg.Add(1)
	go receiveWorker(ch, &wg)

	wg.Wait()

	fmt.Println("Ending of main program...!")
	fmt.Println("Total time taken: ", time.Since(now))
}

// output:
// Starting of main program...!
// Writing to channel...! 0
// Writing to channel...! 1
// Reading from channel...! 0
// Reading from channel...! 1
// Writing to channel...! 2
// Writing to channel...! 3
// Reading from channel...! 2
// Reading from channel...! 3
// Writing to channel...! 4
// Ending of sendWorker function...!
// Reading from channel...! 4
// Ending of receiveWorker function...!
// Ending of main program...!
// Total time taken:  33.979µs
