/* 
	结构体Tag应用：json编解码
		结构体属性名必须大写，否则不会到导出该属性
		若没有Tag，则默认属性名作为json key
*/
package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title string `json:"title"`
	Time float64 `json:"time"`
	Price float64
	Actors []string
}

func main() {
	mov := Movie{"Linux", 113.23, 0.0, []string{"linux教父", "me"}}

	// struct --> json 编码过程
	jsonStr, err := json.Marshal(mov)
	if err != nil {
		fmt.Println("json marshal failed: ", err)
		return 
	}

	// fmt.Println("json: ", jsonStr) //  [123 34 116 105 116 108 101 34 58 34 7...] 形式输出
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// json --> strcut 解码过程
	newmov := Movie{}
	err = json.Unmarshal(jsonStr, &newmov)
	if err != nil {
		fmt.Println("Unmarshal failed: ", err)
		return 
	}

	fmt.Println("struct:", newmov)
}