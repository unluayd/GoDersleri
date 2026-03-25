package main

import "fmt"

type ICoder interface {
	Encode(value string) string
	Decode(value string) string
}

type XCodec struct{}

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
	var coder ICoder

	coder = &XCodec{}
	coder.Encode("123456")
	coder.Decode("abcdef")

	coder = &YCodec{}
	coder.Encode("123456")
	coder.Decode("abcdef")
}
