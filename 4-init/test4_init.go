/* 
import导包路径，init方法调用流程
*/

package main

// 导包init执行顺序不依赖import的顺序，Go 在做包初始化时，会按依赖图做拓扑排序；当两个包互不依赖时，编译器会按“包导入路径的字典序”来确定顺序
import (
	// _ "GOSTUDY/init/lib1"  // 匿名
	// . "GOSTUDY/init/lib1" // 将指定包中全部方法导入当前包，无需包名/别名调用 （warning)
	bag1 "GOSTUDY/init/lib1" 
	"GOSTUDY/init/lib2"
)

// main函数只能在package main中
func main() {
	/* 
	导入包后必须需要使用，否则报错：.\test4_init.go:9:2: "GOSTUDY/init/lib1" imported and not used
	 或者
	import中队包做匿名(_ "GOSTUDY/init/lib1")
	*/


	/* 
	通过包名/别名调用
	*/
	// lib1.Lib1Test()
	bag1.Lib1Test()
	lib2.Lib2Test()
}