package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startNow := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 3")
		time.Sleep(3 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All goroutines finished in", time.Now().Sub(startNow))
}
