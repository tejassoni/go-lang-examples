package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure the context is cancelled when main exits

	signalChan := make(chan os.Signal, 1)

	// Register which signals to catch
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	//	signalChan <- os.Interrupt  // Simulate an OS signal (e.g., Ctrl+C) for demonstration

	dbQueryChan := make(chan string)
	go performDBQuery(ctx, dbQueryChan)

	select {
	case sig := <-signalChan:
		fmt.Printf("\nReceived signal: %v\n", sig)
		fmt.Println("Performing graceful shutdown...")
		return
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
//	Starting of main program...!
//Starts Performing DB query...!
//^C
//Received signal: interrupt
//Performing graceful shutdown...
