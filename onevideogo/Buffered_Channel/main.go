package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan int, 4) // Buffered channel with capacity of 2

	// Start a goroutine to send values to the channel
	go func() {
		for i := 0; i < 10; i++ {
			myChannel <- i
			fmt.Println("Sent:", i)
		}
		close(myChannel) // Close the channel after sending all values
	}()
	time.Sleep(3 * time.Second) // Just to ensure the sender goroutine has started
	// Start a goroutine to receive values from the channel
	go func() {
		for data := range myChannel {
			fmt.Println("Received:", data)
			time.Sleep(1 * time.Second) // Simulate processing time
		}
	}()



}