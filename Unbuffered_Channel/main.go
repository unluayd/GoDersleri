package main

import (
	"fmt"
	"time"
)

// main, unbuffered channel ile senkron veri iletimini gösterir.
func main() {
	// Buffer'sız kanal oluşturulur; gönderme ve alma aynı anda eşleşmelidir.
	myChannel := make(chan string)

	// Ana goroutine'in, alıcı goroutine tamamlanana kadar beklemesi için sinyal kanalı açılır.
	done := make(chan bool)

	// Bu goroutine kanaldan mesajı alır ve işi bittiğinde done kanalına bilgi gönderir.
	go func() {
		message := <-myChannel
		fmt.Println("Received:", message)
		done <- true
	}()

	// Bu goroutine mesajı unbuffered channel üzerinden gönderir.
	go func() {
		myChannel <- "Hello from goroutine 1"
	}()

	// Alıcı taraf işi bitirene kadar burada beklenir.
	<-done
	fmt.Println("Main function is exiting.")

	// Program kapanmadan önce tüm çıktının görünmesi için kısa süre beklenir.
	time.Sleep(1 * time.Second)
}
