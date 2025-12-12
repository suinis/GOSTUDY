package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("type = %T, len = %d, slice = %v\n", arr, len(arr), arr)
	if arr == nil {
		fmt.Println("is empty splice")
	} else {
		fmt.Println("splice is not empty")
	}
 }

func main() {
	// 
	var slice1 []int
	
	//
	slice2 := []int{1,2,3,4}

	// 
	var slice3 []int = make([]int, 3)

	//
	slice4 := make([]int, 3)
	
	/* 
		type = []int, len = 0, slice = []
		is empty splice
		type = []int, len = 4, slice = [1 2 3 4]
		splice is not empty
		type = []int, len = 3, slice = [0 0 0]
		splice is not empty
		type = []int, len = 3, slice = [0 0 0]
		splice is not empty
	*/ 
	printSlice(slice1)
	printSlice(slice2)
	printSlice(slice3)
	printSlice(slice4)
}