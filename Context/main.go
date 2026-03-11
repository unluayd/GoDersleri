package main

import (
	"context"
	"fmt"
	"time"
)

// Product, dış servisten alınan örnek ürün bilgisini temsil eder.
type Product struct {
	id   int64
	Name string
}

// Ürün detaylarını örneklemek için paylaşılan kanal tanımlanır.
var productChannel = make(chan Product)

// main, context ile veri taşıma ve iptal senaryolarını göstermeyi amaçlar.
func main() {
	// Uygulama için temel context oluşturulur.
	ctx := context.Background()

	// Context içine örnek bir correlation id bilgisi eklenir.
	ctx = context.WithValue(ctx, "correlation-id", "abc123")

	// Aynı context değeri fonksiyon zinciri boyunca taşınır.
	F1(ctx)

	// Aşağıdaki örnek kod, timeout ile çalışan alternatif bir senaryoyu gösterir.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()

	// go getProductDetailsFromExternalService(10)

	// select {
	// case productDetail := <-productChannel:
	// 	fmt.Println("Ürün detayları getirildi.", productDetail)
	// case <-ctx.Done():
	// 	fmt.Println("Timeout occurred, context cancelled")
	// }
	// fmt.Println("End of the main")
}

// F1, context içindeki correlation id bilgisini okuyup bir sonraki adıma aktarır.
func F1(ctx context.Context) {
	fmt.Println("F1", ctx.Value("correlation-id"))
	F2(ctx)
}

// F2, context'in fonksiyonlar arasında taşınabildiğini gösterir.
func F2(ctx context.Context) {
	fmt.Println("F2", ctx.Value("correlation-id"))
	F3(ctx)
}

// F3, zincirin son halkasında aynı context değerine erişir.
func F3(ctx context.Context) {
	fmt.Println("F3", ctx.Value("correlation-id"))
}

// getProductDetailsFromExternalService, gecikmeli bir dış servis çağrısını taklit eder.
func getProductDetailsFromExternalService(productId int64) {
	// Dış servis yanıt süresi örneklenir.
	time.Sleep(10 * time.Second)

	// İşlem tamamlandığında ürün bilgisi kanal üzerinden gönderilir.
	productChannel <- Product{
		id: productId,
		Name: "Çamaşır Makinesi",
	}
}
