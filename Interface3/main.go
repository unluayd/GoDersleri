package main

import "fmt"

// ICoder: hem şifreleme hem çözme yapabilen türlerin uyması gereken sözleşme.
// Bu iki metodu birden sağlayan her struct, ICoder olarak kullanılabilir.
type ICoder interface {
	Encode(value string) string
	Decode(value string) string
}

// XCodec: X ailesi kodlama uygulaması (Encode + Decode birlikte).
type XCodec struct{}

// YCodec: Y ailesi kodlama uygulaması (aynı arayüz, farklı davranış).
type YCodec struct{}

func (*XCodec) Encode(value string) string {
	fmt.Println("XCodec Encode", value)
	return "XCodec Encode"
}

func (*XCodec) Decode(value string) string {
	fmt.Println("XCodec Decode", value)
	return "XCodec Decode"
}

func (*YCodec) Encode(value string) string {
	fmt.Println("YCodec Encode", value)
	return "YCodec Encode"
}

func (*YCodec) Decode(value string) string {
	fmt.Println("YCodec Decode", value)
	return "YCodec Decode"
}

func main() {
	// coder hep ICoder tipinde; içine farklı codec'ler konur (çok biçimlilik).
	var coder ICoder

	coder = &XCodec{}
	coder.Encode("123456")
	coder.Decode("abcdef")

	// Aynı değişkene YCodec atanınca çağrılar Y'nin metotlarına gider.
	coder = &YCodec{}
	coder.Encode("123456")
	coder.Decode("abcdef")
}
