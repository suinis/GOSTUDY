package main

import (
	"fmt"
	"runtime"
	"time"
)

// 死锁排查方法，定时打印所有堆栈信息
func deadlock_check() {
	go func() {
		for range time.Tick(5 * time.Second) { // 或收到某个控制信号时再打印
			buf := make([]byte, 1<<20)
			/*
				false 仅当前 goroutine 堆栈
				true  全部 goroutine 堆栈
			*/
			n := runtime.Stack(buf, true)
			fmt.Printf("=== ALL STACKS ===\n%s\n", buf[:n])
		}
	}()
}
