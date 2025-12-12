/* 
1. 四种变量声明方式 := 
2. 全局变量声明 
3. 多变量声明 var ()
*/


package main

import "fmt"

var x, y int 
// global variables could not use :=
// z := 1000 // not allowed 
var (
	xx int = 202303
	yy = "absdgdgdh"
)

func main() {
	// 
	var a int 
	fmt.Println("a = ", a)

	//
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	//
	var c = 1000
	fmt.Printf("c = %d, type of c = %T\n", c, c)

	// := (common)
	d := "1000d"
	fmt.Print("d = ", d)
	fmt.Printf(", type of d = %T\n", d)

	// 多种变量声明
	var e, f, g = 0.23, false, "dstring"
	h, i := 3.41323123213123121, true
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("h = ", h)
	fmt.Printf("type of h = %T\n", h)

	fmt.Println("i = ", i)
	fmt.Printf("type of i = %T\n", i)

	// global variables 
	fmt.Println("x = ", x)
	fmt.Printf("type of x = %T\n", x)

	fmt.Println("y = ", y)
	fmt.Printf("type of y = %T\n", y)

	fmt.Println("xx = ", xx)
	fmt.Printf("type of xx = %T\n", xx)

	fmt.Println("yy = ", yy)
	fmt.Printf("type of yy = %T\n", yy)
}
