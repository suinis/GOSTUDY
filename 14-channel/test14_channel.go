/* 
	channel无缓存设定：管道两端都有阻塞机制
*/

package main

import "fmt"
import "time"

func main() {
	// 未设置cap，channel是无缓冲的管道
	c := make(chan int)

	go func() {
		defer fmt.Println("goorutine exit..")

		fmt.Println("goroutine is running..")
		
		// 如果子协程不传入c管道内对应数据，则主协程会被阻塞
		// for {
		// 	time.Sleep(5 * time.Second)
		// }
		c <- 555 
	}()

	// 如果主协程不获取c管道内对应数据，由于管道无缓冲无法暂存子协程传入的数据，则子协程也会被阻塞
	for {
		time.Sleep(1 * time.Second)
		
	}
   	num := <- c 

	fmt.Println("num : ", num)
	fmt.Println("main goroutine exit..")
}