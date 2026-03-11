package main

import (
	"fmt"
	"time"
)

// main, buffered channel kullanımını temel bir örnekle gösterir.
func main() {
	// 4 eleman kapasiteli bir buffered channel oluşturulur.
	myChannel := make(chan int, 4)

	// Bu goroutine kanala 0'dan 9'a kadar değerler gönderir.
	go func() {
		for i := 0; i < 10; i++ {
			// Kanal dolu değilse değer gönderilir.
			myChannel <- i
			fmt.Println("Sent:", i)
		}
		// Tüm veriler gönderildikten sonra kanal kapatılır.
		close(myChannel)
	}()

	// Gönderici goroutine'in çalışmaya başlaması için kısa bir süre beklenir.
	time.Sleep(3 * time.Second)

	// Bu goroutine kanaldaki verileri sırayla alır ve ekrana yazar.
	go func() {
		for data := range myChannel {
			fmt.Println("Received:", data)
			// Her veri sonrası işlem süresini temsil eden kısa bir bekleme yapılır.
			time.Sleep(1 * time.Second)
		}
	}()
}