/* 
	goroutine 

		嵌套函数执行，return只会终止当前调用栈函数的执行
			若想要直接终止：采用runtime.Goexit()
*/

package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("defer A")

		func() {
			defer fmt.Println("defer B")
			// return  // 
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}