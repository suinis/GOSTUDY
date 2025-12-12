/* 
	map三种定义方式
		map使用前需要先make，make作用：给map分配空间，否则报错：panic: assignment to entry in nil map
*/

package main

import "fmt"

func printMap(hashmap map[string]string) {
	fmt.Println(hashmap)
}

func main() {
	// one
	var hashmap map[string]string 
	if hashmap == nil {
		fmt.Println("map is empty")
	}
	// map使用前需要先make，make作用：给map分配空间，否则报错：panic: assignment to entry in nil map
	hashmap = make(map[string]string, 10) // essenitial
	hashmap["one"] = "apple"
	hashmap["two"] = "banana"
	hashmap["three"] = "peach"
	hashmap["any"] = "stdc++"
	hashmap["abc"] = "first"
	hashmap["abcd"] = "second"
	printMap(hashmap) // map[abc:first abcd:second any:stdc++ one:apple three:peach two:banana]

	// two
	hashmap2 := make(map[int]string)
	hashmap2[1] = "one"
	hashmap2[2] = "two"
	hashmap2[3] = "three"
	fmt.Println(hashmap2)

	// three
	hashmap3 := map[float64]string {
		1.0 : "first version",
		1.1 : "feature version",
		2.0 : "second version", // 最后一对，也需要加','
	}

	fmt.Println(hashmap3)

	hashmap4 := make(map[float64]string, len(hashmap3))
	// copy(hashmap4, hashmap3) // invalid copy: argument must be a slice
	fmt.Println(hashmap4)
}