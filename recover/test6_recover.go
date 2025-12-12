/* 
	recover错误拦截
*/

package main

import "fmt"

// func deferFunc() {
// 	err := recover()
// 	if err != nil {
// 		fmt.Println("捕获到 panic:", err)
// 	}
// }

func test(a int) {
	// defer deferFunc()

	var arr [5]int
	arr[a] = 10
}

func main() {
	// recover 是内置函数，用于从 panic 中恢复
	// 必须在 defer 函数中调用
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic:", r)
		}
	}()

	test(5)
	
	// 触发 panic
	// panic("这是一个测试 panic")
}
 