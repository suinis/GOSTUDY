/* 
	goroutine 协程
*/

package main
import (
	"fmt"
	"time"
)

// 主go程和协程并发打印
func anotherPrint() {
	i := 0

	for {
		i++
		fmt.Printf("goroutine print i : %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go anotherPrint()

	i := 0
	for {
		i++
		fmt.Printf("main print i : %d\n", i)
		time.Sleep(1 * time.Second)
	}	
}