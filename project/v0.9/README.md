# HOW TO USE

## 编译并启动 server
```bash
go build -o server server.go main.go user.go deadlock_check.go
./server
```

## 观测与分析
- 可视化 pprof：`http://127.0.0.1:6060/debug/pprof/`
- 命令行 pprof：
  - CPU
    ```bash
    go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=20
    ```
  - Heap
    ```bash
    go tool pprof http://127.0.0.1:6060/debug/pprof/heap
    ```
  - Trace
    ```bash
    curl http://127.0.0.1:6060/debug/pprof/trace?seconds=5 -o trace.out
    go tool trace trace.out
    ```
  - Vars
    ```bash
    curl http://127.0.0.1:6060/debug/vars
    ```

## TIPS
- [v0.9] 新增 pprof 性能分析工具
- [v0.9] 新增观测端口（默认 `127.0.0.1:6060`），集成 pprof/trace/expvar
- benchstat 对比示例：
  ```bash
  go test -bench=. -benchmem ./... > old.txt
  # 代码优化后
  go test -bench=. -benchmem ./... > new.txt
  benchstat old.txt new.txt
  ```