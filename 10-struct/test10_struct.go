/* 
	struct
		传参方式：值传递
		加* 引用传递，goi内置对指针的解引用
*/
package main

import "fmt"

type myint int

type BOOK struct {
	name string
	author string
}

func valueChangeBook(book BOOK) {
	book.author = "hb"
}

/* 
	在 Go 里，选择器（.）对指针会自动解引用：
	book.author 会被编译器视作 (*book).author。
	因此即便 book 是 *BOOK，也能直接用 . 访问字段，不需要像 C++ 那样 ->。这是语言内置的语法糖；
	同理，调用方法时也会自动在值和指针间做必要的取/解引用。
*/
func referenceChangeBook(book *BOOK) {
	book.author = "hb"
}
  
func main() {
	/* var a myint = 10
	fmt.Printf("%d, type : %T\n", a, a) */

	var book1 BOOK
	book1.name = "c++"
	book1.author = "谭浩强"

	fmt.Println(book1) 
	// struct 值传递
	valueChangeBook(book1)
	fmt.Println(book1)

	// 加* 转为引用传递
	referenceChangeBook(&book1)
	fmt.Println(book1)
}