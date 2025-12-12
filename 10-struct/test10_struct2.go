/* 
	struct used as class 
		类名 属性名 方法名 首字母大写表示对外（其他包）可见，否则只能在本包内访问
		函数中的this是对象的副本，需要修改原对象，需要加指针
*/

package main

import "fmt"

// 类名大写：对其他包可见
// 属性名大写，对其他包可见
type BOOK struct {
	id int
	name string
	price float64
}

func (this BOOK) getName() string{
	return this.name
}

func (this BOOK) setName(name string) {
	this.name = name
}

// 指针传递
func (this *BOOK) setNameRefer(name string) {
	this.name = name
}

func main() {
	book := BOOK{id: 1, name: "APUE", price: 60.4}
	fmt.Println(book)

	// value 
	book.setName("NEWAPUE")
	fmt.Println(book)

	// reference
	book.setNameRefer("NEWAPUE")
	fmt.Println(book)
}