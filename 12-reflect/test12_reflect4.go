/* 
	结构体标签
*/ 
 package main

 import (
	"fmt" 
	"reflect"
)

type resume struct {
	/* 
		Tag 必须是
			反引号内的 key:"value" 形式，
			键值间不能有空格，
			值也需用引号
		
		原标签写成 info: "name"、info: sex，解析不到，tag.Get 返回空。
	*/
	name string `info:"name" doc:"我的名字"`
	sex  string `info:"sex"`
}

func findTag(res interface{}) {
	// fmt.Println(reflect.TypeOf(res)) //*main.resume
	// fmt.Println(reflect.TypeOf(res).Elem()) //main.resume
	// fmt.Println(reflect.TypeOf(res).Elem().NumField()) // 2
	t := reflect.TypeOf(res).Elem()

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		/* 
			info: "name" doc: "我的名字"
			info: sex
		*/
		fmt.Println(tag)
		fmt.Println("info: ", tag.Get("info"), ", doc : ", tag.Get("doc"))
	}
}

func main() {
	res := resume{"zzz", "male"}

	findTag(&res)
}