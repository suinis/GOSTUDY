/* 
	slice切片append、截取、Copy
		append：若len == cap，再次append，cap *= 2
		cut：类似CPP浅拷贝，得到的新slice跟原slice同地址空间
		copy：类似CPP中的深拷贝，直接开辟一段新的空间
*/

package main

import "fmt" 

func printSlice(slice []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)
}

func spliceAppend() {
	numbers := []int{1,2,3} // len = 3, cap = 3
	printSlice(numbers)
	// append
	numbers = append(numbers, 4) // len = 4, cap = 6
	printSlice(numbers)

	fmt.Println("--------------")

	numbers2 := make([]int, 3, 5) // len = 3, cap = 5
	printSlice(numbers2)

	// append
	numbers2 = append(numbers2, 4) // len = 4, cap = 5
	printSlice(numbers2)

	numbers2 = append(numbers2, 5) // len = 5, cap = 5
	printSlice(numbers2)

	numbers2 = append(numbers2, 6) // len = 6, cap = 10
	printSlice(numbers2)
}

func spliceCut() {
	strs := []string{"a", "b", "c"}
	strarr := strs[0 : 2] // [0, 2) 左闭右开
	fmt.Println(strarr)

	// 
	strs[0] = "changed"
	fmt.Println(strs)
	fmt.Println(strarr)
}

func spliceCopy() {
	numbers := []int{1,2,3,4}
	printSlice(numbers)
	numbers2 := make([]int, len(numbers) - 1, cap(numbers))
	printSlice(numbers2) // len = 3, cap = 4, slice = [0 0 0]

	copy(numbers2, numbers) // copy numbers into numbers2
	printSlice(numbers2) // len = 3, cap = 4, slice = [1 2 3]
}

func main() {
	fmt.Println("splice append:")
	spliceAppend()

	fmt.Println("=============")
	fmt.Println("splice cut:")
	spliceCut()

	fmt.Println("=============")
	fmt.Println("splice copy:")
	spliceCopy()
}