package lib1

import "fmt"

// 对外可见的函数，命名必须要首字母大写
func Lib1Test() {
	fmt.Println("lib1.test()..")
}

// init函数可以在一个包中出现多次
func init() {
	fmt.Println("lib1.init()..")
}

func init() {
	fmt.Println("lib1.init2()..")
}