/* 
	map的 add, delete, update, traverse
		map通过引用方式传参 
*/

package main

import "fmt"

func mapTraversal(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key : ", key, ", value : ", value)
	}
}

// map传递方式：引用传递
func insertPair(cityMap map[string]string) {
	cityMap["Endland"] = "London"
}

func searchLanguage(language map[string]map[string]string, target string) {
	value, isfind := language[target]
	// fmt.Println(value, isfind) // map[des:c++是世界上最好的语言 id:1] true
	if isfind {
		fmt.Printf(target + " : %v\n", value)
		return
	} 
	fmt.Println(target, " key is empty")
}

func mapmap() {
	fmt.Println("===================")
	language := make(map[string]map[string]string)
	language["c++"] = make(map[string]string)
	language["c++"]["id"] = "1"
	language["c++"]["des"] = "c++是世界上最好的语言"
	language["golang"] = make(map[string]string)
	language["golang"]["id"] = "2"
	language["golang"]["des"] = "golang更简洁"

	fmt.Println(language)

	fmt.Println("-------------------")
	searchLanguage(language, "c++")
	searchLanguage(language, "c")
}

func main() {
	cityMap := make(map[string]string)

	// add（init）
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	fmt.Println(cityMap)

	// delete
	delete(cityMap, "China")

	// update
	cityMap["USA"] = "WashingtonDC"

	// add
	insertPair(cityMap)

	// traversal
	mapTraversal(cityMap)

	// 嵌套map
	mapmap()
}