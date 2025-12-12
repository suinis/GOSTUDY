/* 
	interface 
		空接口作为万能类型，基本数据类型都实现了interface{}
		提供类型断言机制
*/

package main
import "fmt"

type book struct {
	name string
	price float64
}

// interface{}作为万能类型
// func emptyInterface(arg interface{}) {
// 	// fmt.Println(arg)

// 	value, isok := arg.(string) // type of value is string
// 	if isok {
// 		fmt.Println("arg is string type")
// 		fmt.Printf("value = %v\n", value)
// 	} else {
// 		fmt.Printf("type of arg is : %T\n", arg)
// 		fmt.Printf("type of value is : %T\n", value)
// 	}
// }

// interface{}作为万能类型
func emptyInterface(tmp interface{}) {
	// fmt.Println(tmp)

	value, isok := tmp.(string) // type of value is string
	if isok {
		fmt.Println("tmp is string type")
		fmt.Printf("value = %v\n", value) 
	} else {
		fmt.Printf("type of tmp is : %T\n", tmp)
		fmt.Printf("type of value is : %T\n", value)
	}
	fmt.Println("=======================")
}

func main() {
	book1 := book{"APNP", 79.9}
	emptyInterface(book1)

	emptyInterface(11)
	emptyInterface("abcd")
	emptyInterface(true)
	emptyInterface(33.4)
}