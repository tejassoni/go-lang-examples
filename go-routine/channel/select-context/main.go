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
	case <-time.after(4 * time.Second):
		fmt.Println("Timeout reached")
	}

	fmt.Println("Ending of main program...!")
	fmt.Println("Total time taken: ", time.Since(now))
}

// output:
// Starting of main program...!
// Starts Performing DB query...!
// DB query completed successfully 0xc0001000e0
// Received DB query result: DB query result
// Ending of main program...!
// Total time taken:  3.000230671s
