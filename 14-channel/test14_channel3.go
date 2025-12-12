/* 
	channel的关闭
		channel不需要像文件一样频繁关闭，只需要在数据发送完毕之后/想要显示地结束range for循环，才去关闭channel
		channel close之后，再去向channel发送数据，会导致panic错误，接收后立即返回零值
		close channel之后，可以继续从channel中读取数据
		对于nil channel，无论收发都会阻塞
 */
package main

import (
	"fmt"
	"time"
)

// close channel
func test1() {
	c := make(chan int, 5)

	go func() {
		defer fmt.Println("go routine exit..")
		for i := 1; i <= 5; i++ {
			c <- i
		}

		
		close(c) // 数据传输完毕后，记得close
		fmt.Println("channel c close.")

		// c <- 6 // close channel后再写数据会导致panic: send on closed channel
 	}()

	time.Sleep(2 * time.Second) // close channel后，仍然可以从channel中读取数据

	for {
		// ok表示channel的状态，未关闭：true，关闭:false
		if value, ok := <- c; ok {
			fmt.Println("value : ", value)
		} else {
			break
		}
	}

	fmt.Println("main exit..")
}

// nil channel test
func test2() {
	// nil channel 示例
	var ch chan int // 声明但未初始化，ch 为 nil
	
	// 示例1：从 nil channel 接收会永久阻塞
	go func() {
		fmt.Println("尝试从 nil channel 接收...")
		// 这行代码会永久阻塞，永远不会执行
		value := <-ch
		fmt.Println("接收到:", value) // 永远不会执行到这里
	}()
	
	// 示例2：向 nil channel 发送也会永久阻塞
	go func() {
		fmt.Println("尝试向 nil channel 发送...")
		// 这行代码会永久阻塞，永远不会执行
		ch <- 1
		fmt.Println("发送成功") // 永远不会执行到这里
	}()
	
	time.Sleep(1 * time.Second)
	fmt.Println("主函数继续执行，但 goroutine 被阻塞了")
	
	// 正确的方式：初始化 channel
	ch2 := make(chan int)
	go func() {
		ch2 <- 42
	}()
	
	value := <-ch2
	fmt.Println("从正常 channel 接收到:", value)
	
	// nil channel 的用途：在 select 中用于永久禁用某个 case
	var ch3 chan int
	select {
	case <-ch3: // 这个 case 永远不会被选中
		fmt.Println("不会执行")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("nil channel 在 select 中会永久阻塞，所以这个 case 被选中")
	}
}

// channel range
func test3() {
	c := make(chan int, 5)

	go func() {
		defer close(c) // 发送完数据后关闭 channel，否则 for range 会永久阻塞
		for i := 1; i <= 5; i++ {
			c <- i
		}
		fmt.Println("数据发送完毕，channel 已关闭")
	}()
	
	// for range 会一直从 channel 读取，直到 channel 被关闭
	for data := range c {
		fmt.Println("data : ", data)
	}
	fmt.Println("for range 结束，因为 channel 已关闭")
}

func main() {
	// test1()
	// test2()
	test3()
}