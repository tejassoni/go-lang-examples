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
	time.Sleep(5 * time.Second)

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
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() // Ensure the context is cancelled when main exits

	dbQueryChan := make(chan string)
	go performDBQuery(ctx, dbQueryChan)

	select {
	case result, ok := <-dbQueryChan:
		if !ok {
			fmt.Println("DB query channel closed without result")
			return
		}
		fmt.Println("Received DB query result:", result)
	case <-ctx.Done():
		fmt.Println("Context cancelled or timed out", ctx.Err())
	}

	fmt.Println("Ending of main program...!")
	fmt.Println("Total time taken: ", time.Since(now))
}

// output:
// Starting of main program...!
// Starts Performing DB query...!
// Context cancelled or timed out context deadline exceeded
// Ending of main program...!
// Total time taken:  4.001550049s
