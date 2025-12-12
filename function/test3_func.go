package main

/* 函数的多返回值 */

import "fmt"

func func1(a string, b int) int {
	fmt.Println("===func1()===")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 1000
}

func func2(a string, b bool) (int, string) {
	fmt.Println("===func2()===")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 1000, "ok"
}

// 带名称的返回值，有作用域的，未设定值时默认初始化
func func3(a , b string) (r1 int, r2 bool) {
	fmt.Println("===func3()===")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Printf("r1 = %d, r2 = %t\n", r1, r2) // 0, false
	r1 = 1000
	r2 = true
	return 
}

func main() {
	ret1 := func1("abcd", 123)
	fmt.Println("ret1 = ", ret1)

	ret2_1, ret2_2 := func2("func2", true)
	fmt.Println("ret2_1 = ", ret2_1)
	fmt.Println("ret2_2 = ", ret2_2)

	r1, r2 := func3("func3_1", "func3_2")
	fmt.Printf("r1 = %d, r2 = %t\n", r1, r2)
}