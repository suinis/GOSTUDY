package main

import (
	observability "GOSTUDY/project/v0.9/tools"
	"context"
	"log"

	"github.com/google/gops/agent"
)

func main() {
	// gops
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	// pprof
	stopObs, err := observability.Start(observability.Config{
		Addr: "127.0.0.1:6060", // 可根据环境调整或通过配置/环境变量注入
	})
	if err != nil {
		log.Fatalf("启动观测端口失败: %v", err)
	}
	defer stopObs(context.Background())

	// server
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
