/* 
	变量结构：{type, value}的键值对 
		type: statictype , concretetype

	reflect 反射 ，一般在一些对性能要求不高的场景会使用
*/  

package main

import (
	"fmt"
	"os"
	"io"
)

func pairFunc() {
	var str string
	// pair<statictype: string, value: "abcde">
	str = "abcde"

	// pair<type: string, value "abcde">
	var alltype interface{}
	alltype = str

	value, _:= alltype.(string)
	// value := alltype.(string)
	fmt.Println(value)
}

// to do
// func ioFunc() {
// 	// tty: pair<type: *os.File, >
// 	tty, err := os.OpenFile("../file.txt", os.O_RDWR, 0) // "/dev/tty" ： linux终端
// 	if err != nil {
// 		fmt.Println("open file error", err)
// 		return 
// 	} 

// 	var r io.Reader
// 	r = tty

// 	var w io.Writer
// 	w  = r.()
// }

func main()  {
	pairFunc()

	ioFunc()
}