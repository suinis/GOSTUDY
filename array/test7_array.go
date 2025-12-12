package main

import "fmt"

// value copy
func printArray(arr [4]string) {

	for index, value := range arr {
		fmt.Println("index= ", index, ", value= ", value)
	}

	arr[0] = "abcd"
}

func main() {
	// 定长数组
	var arr1 [10]int
	arr2 := [10]int{1,2,3,4} // 1 2 3 4 0 ...
	arr3 := [4]string{"a", "b", "c", "d"}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for index, value := range arr2 {
		fmt.Println("index = ", index , ", value = ", value)
	}

	printArray(arr3)
	fmt.Printf("type of arr3: %T\n", arr3)
	printArray(arr3)
}