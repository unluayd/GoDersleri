package main

import (
	"fmt"
	"time"
)

// main, select yapısı ile birden fazla kanalı aynı anda dinlemeyi gösterir.
func main() {
	// İki ayrı string kanalı oluşturulur.
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Kanallardan gelecek veriler bu değişkenlerde tutulur.
	var data1 string
	var data2 string

	// İlk goroutine 10 saniye sonra birinci kanala veri gönderir.
	go func() {
		time.Sleep(10 * time.Second)
		channel1 <- "Hello"
	}()

	// İkinci goroutine 5 saniye sonra ikinci kanala veri gönderir.
	go func() {
		time.Sleep(5 * time.Second)
		channel2 <- "World"
	}()

	// Her iki kanaldan da veri gelene kadar select döngüsü çalıştırılır.
	for len(data1) == 0 || len(data2) == 0 {
		select {
		case data1 = <-channel1:
			fmt.Println("Received data from channel1 : ", data1)
		case data2 = <-channel2:
			fmt.Println("Received data from channel2 : ", data2)
		default:
			// Henüz hiçbir kanalda veri yoksa default bloğu çalışır.
			fmt.Println("Veri Gelmedi.")
		}

		// Döngünün çok hızlı çalışmasını önlemek için kısa bir bekleme yapılır.
		time.Sleep(1 * time.Second)
	}
}
