package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan int) // Unbuffered channel
	// Start a goroutine to receive values from the channel
	go func() {
	 	for i:=0; i<=10; i++ {
		myChannel <-i
		fmt.Println("Sent:", i)
		time.Sleep(1*time.Second)
		}
		close(myChannel)
	}()
	for {
		data, isOpen := <-myChannel
		if !isOpen {
			fmt.Println("Channel is closed, exiting loop.")
			break
		}
		fmt.Println("Received:", data, isOpen)
	}
}