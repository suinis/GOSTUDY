/* 
	1. 函数内部定义方式： func() {}
	若要执行：
		不带参函数执行：func() {}()，末尾加()执行
		带参函数执行： func(a int, b int){...}(4, 5) 末尾加(......)执行

	2. 主go程结束，则协程终止
*/

package main

import (
	"fmt"
	// "time"
)

func main() {
	// go func() {
	// 	i := 0
	// 	for {
	// 		i++
	// 		fmt.Printf("goroutine print i : %d\n", i)
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("goroutine print i : %d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }

	go func() {
		defer fmt.Println("defer A")

		func() {
			defer fmt.Println("defer B")

			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	go func(a int, b int) {
		fmt.Println("a = ", a, ", b = ", b)
	}(1, 2)

	// for {
	// 	time.Sleep(1 * time.Second)
	// }
	fmt.Println("main func exit")
}