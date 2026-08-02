package main

import (
	"fmt"
	"time"
)

// channel is Bidirectional, it can only be used to send data
func writer(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		fmt.Println("Writing to channel...!", i)
		ch <- i
	}
	fmt.Println("Ending of writer function...!")
}

// reader function that reads from a Bidirectional channel
func reader(ch <-chan int) {
	for readValue := range ch {
		fmt.Println("Reading from channel...!", readValue)
	}
	fmt.Println("Ending of reader function...!")
}

func main() {
	now := time.Now()
	fmt.Println("Starting of main program...!")

	// create a Bidirectional channel
	ch := make(chan int)
	// start a go routine to write to channel
	go writer(ch)

	// read from channel until it is closed
	reader(ch)

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
// Ending of writer function...!
// Reading from channel...! 4
// Ending of main program...!
// Total time taken:  158.136µs
