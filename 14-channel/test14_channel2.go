/* 
	带缓存的channel：
		缓存已满：写入数据会阻塞
		缓存为空，读出数据也会阻塞
*/

package main

import (
	"fmt"
	"time"
)

/* 
	主程序没有sleep，执行顺序有多种可能
	goroutine is running..
	main get i :  0
	goroutine transfers i : 0, len(channel) : 0, cap(channel) : 3
	goroutine transfers i : 1, len(channel) : 0, cap(channel) : 3
	goroutine transfers i : 2, len(channel) : 1, cap(channel) : 3
	goroutine exit..
	main get i :  2
	main exit..

	goroutine is running..
	main get i :  0
	goroutine transfers i : 0, len(channel) : 0, cap(channel) : 3
	main get i :  1
	goroutine transfers i : 1, len(channel) : 0, cap(channel) : 3
	goroutine transfers i : 2, len(channel) : 0, cap(channel) : 3
	goroutine exit..
	main get i :  2
	main exit..

	goroutine is running..
	goroutine transfers i : 0, len(channel) : 0, cap(channel) : 3
	goroutine transfers i : 1, len(channel) : 1, cap(channel) : 3
	goroutine transfers i : 2, len(channel) : 2, cap(channel) : 3
	goroutine exit..
	main get i :  0
	main get i :  1
	main get i :  2
	main exit..
*/
func test1() {
	c := make(chan int, 3) // 有缓存的channel

	go func() {
		fmt.Println("goroutine is running..")

		defer fmt.Println("goroutine exit..")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("goroutine transfers i : %d, len(channel) : %d, cap(channel) : %d\n", i, len(c), cap(c))
		}
	}()

	for i := 0; i < 3; i++ {
		num := <- c
		fmt.Println("main get i : ", num)
	} 

	fmt.Println("main exit..")
}

/* 
	standard test
*/
func test2() {
	c := make(chan int, 3) // 有缓存的channel

	go func() {
		fmt.Println("goroutine is running..")

		defer fmt.Println("goroutine exit..")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Printf("goroutine transfers i : %d, len(channel) : %d, cap(channel) : %d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <- c
		fmt.Println("main get i : ", num)
	} 

	fmt.Println("main exit..")
}

/* 
未等待子协程：
	goroutine is running..
	goroutine transfers i : 1, len(channel) : 1, cap(channel) : 3
	goroutine transfers i : 2, len(channel) : 2, cap(channel) : 3
	goroutine transfers i : 3, len(channel) : 3, cap(channel) : 3
	main get i :  1
	main get i :  2
	main get i :  3
	main get i :  4
	main exit..

等待子协程：
	goroutine is running..
	goroutine transfers i : 1, len(channel) : 1, cap(channel) : 3
	goroutine transfers i : 2, len(channel) : 2, cap(channel) : 3
	goroutine transfers i : 3, len(channel) : 3, cap(channel) : 3
	main get i :  1
	main get i :  2
	main get i :  3
	main get i :  4
	goroutine transfers i : 4, len(channel) : 3, cap(channel) : 3
	goroutine exit..
	main exit..
*/
func test3() {
	c := make(chan int, 3) // 有缓存的channel

	go func() {
		fmt.Println("goroutine is running..")

		defer fmt.Println("goroutine exit..")

		for i := 1; i <= 4; i++ {
			c <- i
			fmt.Printf("goroutine transfers i : %d, len(channel) : %d, cap(channel) : %d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <- c
		fmt.Println("main get i : ", num)
	} 
	// time.Sleep(2 * time.Second) //等待子协程执行完毕

	fmt.Println("main exit..")
}
 
func main() {
	// test1()
	// test2()
	test3()
}