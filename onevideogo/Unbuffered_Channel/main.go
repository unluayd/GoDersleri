package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan string) // Unbuffered channel
	done := make(chan bool)

	go func() {
		message := <-myChannel
		fmt.Println("Received:", message)
		done <- true
	}()

	go func() {
		myChannel <- "Hello from goroutine 1"
	}()

	<-done // Wait for the first goroutine to finish
	fmt.Println("Main function is exiting.")
	time.Sleep(1 * time.Second) // Just to ensure all output is printed before the program exits
}
