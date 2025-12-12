package main

import "fmt"


/* 
	defer：类似CPP析构函数
	内部保存调用栈，先进后出

	作用：
		类似CPP析构函数，释放资源
		捕获处理异常
		print log
*/
func deferFunc() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

func returnFunc() int {
	fmt.Println("returnFunc()...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()

	return returnFunc()
}

func main() {
	/* 
		returnFunc()...
		4
		3
		2
		1
	*/
	returnAndDefer()
}