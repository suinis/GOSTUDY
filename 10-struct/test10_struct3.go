/* 
	struct 
		面向对象继承

*/
package main

import "fmt"

type Human struct {
	name string
	age int
}

func (this *Human) walk() {
	fmt.Println("human is walking..")
}

func (this* Human) eat() {
	fmt.Println("human is eating..")
}

type Adult struct {
	Human
	ismarried bool
	hasoffer bool
}

func (this* Adult) eat() {
	fmt.Println("the adult is eating..")
}

func (this *Adult) printAdult() {
	fmt.Println("name: ", this.name)
	fmt.Println("age: ", this.age)
	fmt.Println("ismarried: ", this.ismarried)
	fmt.Println("hasoffer: ", this.hasoffer)
}

func main() {
	human := Human{name: "zhangsan", age: 8}
	fmt.Println(human)
	human.walk()
	human.eat()

	// 方式1：全部使用键值对初始化（推荐）
	adult := Adult{Human: Human{name: "lisi", age: 27}, ismarried: false, hasoffer: true}
	
	// 方式2：全部使用位置初始化
	// adult := Adult{Human{name: "lisi", age: 27}, false, true}
	adult.walk()
	adult.eat()
	adult.printAdult()
}