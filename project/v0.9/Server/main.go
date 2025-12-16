package main

import (
	observability "GOSTUDY/project/v0.9/tools"
	"context"
	"log"
)

func main() {
	stopObs, err := observability.Start(observability.Config{
		Addr: "127.0.0.1:6060", // 可根据环境调整或通过配置/环境变量注入
	})
	if err != nil {
		log.Fatalf("启动观测端口失败: %v", err)
	}
	defer stopObs(context.Background())

	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
