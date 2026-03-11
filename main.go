package main

import (
	"fmt"
	"sync"
	"time"
)

// main, WaitGroup ile birden fazla goroutine'in bitmesini beklemeyi gösterir.
func main() {
	// Toplam çalışma süresini ölçmek için başlangıç zamanı alınır.
	startNow := time.Now()

	// Üç goroutine'in tamamlanmasını takip edecek sayaç oluşturulur.
	wg := sync.WaitGroup{}
	wg.Add(3)

	// İlk goroutine kendi işini yaptıktan sonra WaitGroup sayacını azaltır.
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
		time.Sleep(3 * time.Second)
	}()

	// İkinci goroutine örnek amaçlı benzer bir işi paralel yürütür.
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
		time.Sleep(3 * time.Second)
	}()

	// Üçüncü goroutine de tamamlandığında sayaç bir kez daha azaltılır.
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 3")
		time.Sleep(3 * time.Second)
	}()

	// Ana goroutine, diğer tüm goroutine'ler bitene kadar burada bekler.
	wg.Wait()

	// Tüm işlemler tamamlandığında toplam geçen süre ekrana yazdırılır.
	fmt.Println("All goroutines finished in", time.Since(startNow))
}
