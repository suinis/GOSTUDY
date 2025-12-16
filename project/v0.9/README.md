# HOW TO USE
1. 编译启动server:
```
    go build -o server server.go main.go user.go deadlock_check.go 
    ./server
```
2. 可视化pprof地址：http://127.0.0.1:6060/debug/pprof/
   or 命令行pprof
        - CPU:   
        ```
            go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=20
        ```
        - Heap:  
        ```
            go tool pprof http://127.0.0.1:6060/debug/pprof/heap
        ```
        - Trace: 
        ```
            curl http://127.0.0.1:6060/debug/pprof/trace?seconds=5 -o trace.out
            go tool trace trace.out
        ```
        - Vars:  
        ```
            curl http://127.0.0.1:6060/debug/vars
        ```

# TIPS
```
    [v0.9] 新增pprof性能分析工具
    [v0.9] 新增观测端口 (默认 127.0.0.1:6060)，集成 pprof/trace/expvar：

    benchstat 示例（对比基准测试优化前后效果）：
        go test -bench=. -benchmem ./... > old.txt
        # 代码优化后
        go test -bench=. -benchmem ./... > new.txt
        benchstat old.txt new.txt 
```