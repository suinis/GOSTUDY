/* 
	channel & select
*/

package main

import (
	"fmt"
	"time"
)

// channel 传递的是引用
func fib(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x :
			temp := x
			x = y
			y = temp + y
		case <- quit :
			fmt.Println("quit")
			return // ✅ 退出整个函数和 for 循环
		// 移除了 default case，让 select 阻塞等待，直到某个 case 可以执行
		// 如果保留 default，当 channel 操作不能立即执行时会立即 return，导致函数提前退出
		}
	}
}

// channel 传递的是引用
func fib2(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x :
			temp := x
			x = y
			y = temp + y
		case <- quit :
			fmt.Println("quit")
			break // ❌ 如果使用 break，会发生死锁：
			// 1. break 只退出 select，for 循环继续执行
			// 2. 再次进入 select 时，quit channel 已经读取过了（无缓冲 channel，数据已被消费）
			// 3. case <- quit 不会再触发（没有新数据）
			// 4. case c <- x 也无法执行（接收者 goroutine 已退出，不再接收）
			// 5. 两个 case 都无法执行，select 永久阻塞 → 死锁！fatal error: all goroutines are asleep - deadlock!
		}
	}
}

/* 
	1
	1
	2
	3
	5
	8
	quit
*/
func test1() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<- c)
		}

		quit <- 0
	}()

	fib(c, quit)
}

/* 
	1
	1
	2
	3
	5
	8
	quit
	13
	21
	34
	55
	89
	144
	fatal error: all goroutines are asleep - deadlock!
*/
func test2() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<- c)
		}

		quit <- 0

		/* 
			测试select两个都阻塞，其中一个退出阻塞态的情况
		*/
		time.Sleep(2 * time.Second)
		for i := 0; i < 6; i++ {
			fmt.Println(<- c)
		}
	}()

	fib2(c, quit)
}

func main() {
	// test1()
	test2()
}