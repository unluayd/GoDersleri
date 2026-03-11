package main

import (
	"fmt"
	"time"
)

// main, unbuffered channel üzerinde gönderme, alma ve kapatma akışını gösterir.
func main() {
	// Buffer'sız integer kanalı oluşturulur.
	myChannel := make(chan int)

	// Bu goroutine sayıları sırayla kanala gönderir.
	go func() {
		for i := 0; i <= 10; i++ {
			// Unbuffered channel olduğu için her gönderim bir alıcı bekler.
			myChannel <- i
			fmt.Println("Sent:", i)
			time.Sleep(1 * time.Second)
		}

		// Tüm veriler gönderildikten sonra kanal kapatılır.
		close(myChannel)
	}()

	// Kanal açık olduğu sürece veriler manuel olarak alınır.
	for {
		data, isOpen := <-myChannel
		if !isOpen {
			fmt.Println("Channel is closed, exiting loop.")
			break
		}

		// Alınan veri ve kanalın açık olduğu bilgisi ekrana yazdırılır.
		fmt.Println("Received:", data, isOpen)
	}
}