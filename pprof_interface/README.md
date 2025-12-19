## pprof/gops性能分析工具使用接口

### gops
gops启动接口
```
// StartGops 启动gops诊断工具
func StartGops() {
    if err := agent.Listen(agent.Options{}); err != nil {
        log.Printf("[observability] 启动gops失败: %v", err)
        return
    }
    log.Printf("[observability] gops已启动")
}
```

### pprof 
pprof启动接口：依据业务需求传入addr
```
// StartPProf 启动pprof性能分析
// addr为空时使用默认地址 127.0.0.1:1234
func StartPProf(addr string) {
	if addr == "" {
		addr = "127.0.0.1:1234"
	}

	shutdown, err := startPProfServer(addr)
	if err != nil {
		log.Printf("[observability] 启动pprof失败: %v", err)
		return
	}

	pprofShutdown = shutdown
	log.Printf("[observability] pprof监听在 http://%s", addr)
}
```

pprof资源回收接口
```
// CleanAllObserv 清理所有观测性资源
func CleanAllObserv() {
	// 关闭pprof
	if pprofShutdown != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := pprofShutdown(ctx); err != nil {
			log.Printf("[observability] 关闭pprof失败: %v", err)
		} else {
			log.Printf("[observability] pprof已关闭")
		}
	}
}
```

