package main

import "fmt"

// Shippable: kargo ücreti hesaplanabilen her ürün türü için ortak sözleşme.
// Farklı struct'lar aynı metot imzasını sağlayarak bu arayüzü karşılar.
type IShippable interface {
	ShippingCost() int
}

type Books struct {
	desi int // hacim / ağırlık birimi (örnek)
}

type Electronics struct {
	desi int
}

func (b *Books) ShippingCost() int {
	return b.desi * 10
}

func (e *Electronics) ShippingCost() int {
	return e.desi * 20
}

func main() {
	// Aynı Shippable değişkeni önce kitap, sonra elektronik tutabilir (çok biçimlilik).
	var product IShippable
	product = &Books{desi: 10}
	fmt.Println("Kitap kargo:", product.ShippingCost())
	product = &Electronics{desi: 10}
	fmt.Println("Elektronik kargo:", product.ShippingCost())

	// Dilimde farklı türler; hepsi Shippable olduğu için birlikte işlenir.
	cart := []IShippable{&Books{desi: 2}, &Electronics{desi: 3}}
	fmt.Println("Sepet toplamı:", CalculateShippingCost(cart))
}

// CalculateShippingCost: arayüz dilimini bilmezden gelir; somut türleri import etmez.
func CalculateShippingCost(products []IShippable) int {
	totalCost := 0
	for _, product := range products {
		totalCost += product.ShippingCost()
	}
	return totalCost
}
