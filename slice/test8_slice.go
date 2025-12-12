package main

import "fmt"

// reference copy
func printArray(myArray []int) {
	// _表示匿名变量
	for _, value := range myArray {
		fmt.Println(value)
	}
}

func changevalue(myArray []int) {
	myArray[0] = 100
}

func main() {
	arr := []int{1,2,3,4} // dynamic array

	fmt.Printf("type of slice : %T\n", arr)
	printArray(arr)
	changevalue(arr)
	printArray(arr)
}