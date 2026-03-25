package main

import (
	"fmt"
)

// main, bir unbuffered channel üzerinde birden fazla alıcı olduğunda ne olabileceğini gösterir.
func main() {
	// Buffer'sız kanal oluşturulur.
	myChannel := make(chan int)

	// İlk goroutine kanaldan gelen değeri okuyup ekrana yazar.
	go func() {
		data := <-myChannel
		fmt.Println("First go routine took data:", data)
	}()

	// İkinci goroutine de aynı kanaldan veri bekler.
	go func() {
		data := <-myChannel
		fmt.Println("Second go routine took data:", data)
	}()

	// Kanala yalnızca tek bir değer gönderildiği için sadece bir alıcı bu değeri alacaktır.
	myChannel <- 10
	fmt.Println("End of the main function")
}