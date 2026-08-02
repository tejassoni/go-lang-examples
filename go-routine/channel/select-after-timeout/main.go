package main

import (
	"fmt"
	"time"
)

func performDBQuery(dbQueryChan chan<- string) {
	defer close(dbQueryChan)

	fmt.Println("Starts Performing DB query...!")

	// Simulate a long-running DB query
	time.Sleep(3 * time.Second)

	var result string = "DB query result"

	dbQueryChan <- result

	fmt.Println("Ends Performed DB query...!")
}

func main() {
	now := time.Now()
	fmt.Println("Starting of main program...!")

	dbQueryChan := make(chan string)
	go performDBQuery(dbQueryChan)

	select {
	case result, ok := <-dbQueryChan:
		if !ok {
			fmt.Println("DB query channel closed without result")
			return
		}
		fmt.Println("Received DB query result:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout reached")
	}

	fmt.Println("Ending of main program...!")
	fmt.Println("Total time taken: ", time.Since(now))
}

// output:
// Starting of main program...!
// Starts Performing DB query...!
// Timeout reached
// Ending of main program...!
// Total time taken:  1.000865475s
