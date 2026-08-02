package main

import (
	"context"
	"fmt"
	"time"
)

func performDBQuery(ctx context.Context, dbQueryChan chan<- string) {
	defer close(dbQueryChan)

	fmt.Println("Starts Performing DB query...!")

	// Simulate a long-running DB query
	time.Sleep(3 * time.Second)

	var result string = "DB query result"

	// added result to channel before checking context
	dbQueryChan <- result

	select {
	case <-ctx.Done():
		fmt.Println("Database query processing cancelled...! Context cancelled or timed out")
	case dbQueryChan <- result:
		fmt.Println("DB query completed successfully", dbQueryChan)
	default:
		fmt.Println("Default case executed, no one is receiving the result from dbQueryChan")
	}

	fmt.Println("Ends Performed DB query...!")
}

func main() {
	now := time.Now()
	fmt.Println("Starting of main program...!")
	ctx, cancel := context.WithCancel(context.Background())

	dbQueryChan := make(chan string)
	go performDBQuery(ctx, dbQueryChan)

	// Simulate a timeout or cancellation after 4 seconds
	go func() {
		time.Sleep(4 * time.Second)
		cancel() // Cancel the context after 4 seconds
	}()

	select {
	case result, ok := <-dbQueryChan:
		if !ok {
			fmt.Println("DB query channel closed without result")
			return
		}
		fmt.Println("Received DB query result:", result)
	case <-ctx.Done():
		fmt.Println("Context cancelled or timed out")
	}

	fmt.Println("Ending of main program...!")
	fmt.Println("Total time taken: ", time.Since(now))
}

// output:
// Starting of main program...!
// Starts Performing DB query...!
// Default case executed, no one is receiving the result from dbQueryChan
// Ends Performed DB query...!
// Received DB query result: DB query result
// Ending of main program...!
// Total time taken:  3.002458478s
