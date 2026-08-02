package main

import (
	"fmt"
	"time"
)

func emailWorker(emailChan chan<- int) {
	defer close(emailChan)
	for i := 0; i < 5; i++ {
		fmt.Println("Sending email...!", i)
		emailChan <- i
	}
	fmt.Println("Ending of emailWorker function...!")
}

func whatsAppWorker(whatsAppChan chan<- int) {
	defer close(whatsAppChan)
	for i := 0; i < 5; i++ {
		fmt.Println("Sending WhatsApp message...!", i)
		whatsAppChan <- i
	}
	fmt.Println("Ending of whatsAppWorker function...!")
}

func smsWorker(smsChan chan<- int) {
	defer close(smsChan)
	for i := 0; i < 5; i++ {
		fmt.Println("Sending SMS...!", i)
		smsChan <- i
	}
	fmt.Println("Ending of smsWorker function...!")
}

func main() {
	now := time.Now()
	fmt.Println("Starting of main program...!")

	emailChan := make(chan int)
	whatsAppChan := make(chan int)
	smsChan := make(chan int)

	go emailWorker(emailChan)
	go whatsAppWorker(whatsAppChan)
	go smsWorker(smsChan)

	for emailChan != nil || whatsAppChan != nil || smsChan != nil {
		select {
		case emailMsg, ok := <-emailChan:
			if !ok {
				emailChan = nil
			}
			fmt.Println("Received email message: ", emailMsg)
		case whatsAppMsg, ok := <-whatsAppChan:
			if !ok {
				whatsAppChan = nil
			}
			fmt.Println("Received from WhatsApp channel", whatsAppMsg)
		case smsMsg, ok := <-smsChan:
			if !ok {
				smsChan = nil
			}
			fmt.Println("Received from SMS channel", smsMsg)
		}
	}

	fmt.Println("Ending of main program...!")
	fmt.Println("Total time taken: ", time.Since(now))
}

// output:
// Starting of main program...!
// Received from channel 2
// Ending of main program...!
// Total time taken:  1.000335241s
