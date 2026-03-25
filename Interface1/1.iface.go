package main

import "fmt"

// iletisim: topla() metodunu tanımlayan arayüz (interface).
// Bu imzayı sağlayan her tür, iletisim olarak kullanılabilir.
type iletisim interface {
	topla() int
}

// sayilar: üç tamsayı tutan yapı; *sayilar üzerinde topla() tanımlıdır.
type sayilar struct {
	bir, iki, uc int
}

func main() {
	// a: somut tür yerine arayüz tipinde değişken (çok biçimlilik).
	var a iletisim
	v := sayilar{3, 4, 5}
	// &v: *sayilar, iletisim'i topla() ile gerçekleştirir.
	a = &v
	fmt.Println("sayilar üzerinden:", a.topla())

	// sayi: değer alıcılı topla(); aynı iletisim değişkenine atanabilir.
	a = sayi(10)
	fmt.Println("sayi üzerinden:", a.topla())
}

// sayi: int için takma ad; topla() bu değeri olduğu gibi döndürür.
type sayi int

func (f sayi) topla() int {
	return int(f)
}

// (*sayilar).topla: üç alanın toplamını döndürür; iletisim'i karşılar.
func (sayi *sayilar) topla() int {
	return sayi.bir + sayi.iki + sayi.uc
}
