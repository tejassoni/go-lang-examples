package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Starting of main program...!")

	ch1 := make(chan int)
	ch2 := make(chan int)

	// write to channel 1 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()
	// write to channel 2 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 1
	}()

	select {

	case <-ch1:
		fmt.Println("Received from channel 1")
	case <-ch2:
		fmt.Println("Received from channel 2")
	}

	fmt.Println("Ending of main program...!")
	fmt.Println("Total time taken: ", time.Since(now))
}

// output:
// Starting of main program...!
// Received from channel 2
// Ending of main program...!
// Total time taken:  1.000335241s
