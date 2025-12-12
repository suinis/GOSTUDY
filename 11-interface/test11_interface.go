 /* 
 	interface 接口 ：本质上是一个指针，实现多态
 		多态基本要素：
 			有一个父类（接口）
			有子类（继承了父类的所有接口方法）
 			父类类型的变量（指针）指向子类的具体数据变量
 */

package main

import "fmt"

// 父类（接口）
type animal interface {
	gettype() string
	sleep()
}

type cat struct {
	kind string
}

type dog struct {
	kind string
}

func (this *cat) sleep() {
	fmt.Println("cat is sleeping...")
}

func (this *cat) gettype() string{
	return this.kind
}

func (this *dog) sleep() {
	fmt.Println("dog is sleeping...")
}

func (this *dog) gettype() string{
	return this.kind
}

func main() {
	var ani animal
	ani = &cat{"大橘猫"}
	fmt.Println(ani.gettype())
	ani.sleep()

	ani = &dog{"萨摩耶"}
	fmt.Println(ani.gettype())
	ani.sleep()
}