package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type BOOK struct {

}

func (this *BOOK) Write() {
	fmt.Println("writing")
}

func (this *BOOK) Read() {
	fmt.Println("reading")
}

func main() {
	// pair<type: BOOK, value: BOOK{}地址>
	book1 := &BOOK{}
	
	// reader : pair<type: , value: >
	var reader Reader
	// reader : pair<type: BOOK, value: BOOK{}地址>
	reader = book1
	reader.Read()

	// writer: pair<type: , value: >
	var writer Writer
	// reader : pair<type: BOOK, value: BOOK{}地址>
	writer = reader.(Writer) // 此处的断言为什么能成功？因为reader和writer的具体type一致
	writer.Write()
}