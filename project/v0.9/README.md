# 使用指南

## 编译并启动 server
```bash
cd Server
go build -o server .
./server
```

## 模拟连接
```bash
  # 模拟1个线程使用1个连接，持续30s
  wrk -t1 -c1 -d30s http://localhost:8888/
```

## 观测与分析

### 方式一：浏览器访问（推荐，无需额外安装）

#### 传统 pprof UI（推荐）
直接在浏览器中打开：`http://127.0.0.1:6060/debug/pprof/`

### 方式二：命令行 pprof

#### 安装 Graphviz（如需使用可视化图形）
如果需要在 pprof 交互式命令行中使用 `web` 命令生成可视化图形，需要安装 Graphviz：
```bash
# Windows
#下载对应版本: 
https://graphviz.org/download/
#勾选 add to path，安装完后重启终端即可

# Ubuntu/Debian
sudo apt-get install graphviz

# macOS
brew install graphviz

# 安装后即可在 pprof 交互式命令行中使用 web 命令
```

#### 使用 pprof 命令行工具
- **CPU 分析**
  ```bash
  go tool pprof http://127.0.0.1:6060/debug/pprof/profile?seconds=20
  # 进入交互式命令行后可以使用：
  # - top：查看占用 CPU 最多的函数
  # - list 函数名：查看函数详细代码
  # - web：生成调用图（需要 Graphviz）
  # - png/gif：导出为图片（需要 Graphviz）
  # - exit：退出
  ```
- **内存分析**
  ```bash
  go tool pprof http://127.0.0.1:6060/debug/pprof/heap
  # 交互式命令同上
  ```
- **其他分析**
  ```bash
  # Goroutine 分析
  go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine
  
  # Mutex 分析
  go tool pprof http://127.0.0.1:6060/debug/pprof/mutex
  
  # Block 分析
  go tool pprof http://127.0.0.1:6060/debug/pprof/block
  ```

### 方式三：交互式 Web UI（推荐用于可视化分析）

使用 `go tool pprof -http` 启动本地 HTTP 服务器，提供功能丰富的交互式 web 界面，支持火焰图、调用图等可视化：

- **内存分析**
  ```bash
  go tool pprof -http=:8080 http://127.0.0.1:6060/debug/pprof/heap
  # 然后在浏览器中打开 http://localhost:8080
  ```

- **CPU 分析**
  ```bash
  go tool pprof -http=:8080 http://127.0.0.1:6060/debug/pprof/profile?seconds=30
  # 然后在浏览器中打开 http://localhost:8080
  ```

- **其他分析类型**
  ```bash
  # Goroutine 分析
  go tool pprof -http=:8080 http://127.0.0.1:6060/debug/pprof/goroutine
  
  # Mutex 分析
  go tool pprof -http=:8080 http://127.0.0.1:6060/debug/pprof/mutex
  
  # Block 分析
  go tool pprof -http=:8080 http://127.0.0.1:6060/debug/pprof/block
  ```

- **使用已导出的 pprof 文件**
  ```bash
  go tool pprof -http=:8080 heap.pprof
  # 或
  go tool pprof -http=:8080 cpu.pprof
  ```

**Web UI 功能**：
- `/` - 概览页面，显示 top 函数列表
- `/top` - 查看占用资源最多的函数
- `/graph` - 查看函数调用关系图（需要 Graphviz）
- `/flamegraph` - 查看火焰图
- `/source` - 查看源代码级别的分析
- `/peek` - 查看函数调用树

**注意**：如果遇到**背景全黑、字体也全黑看不清**的问题（Go 1.21+ 新 UI 的暗色主题显示问题），可以在浏览器开发者工具（F12）的 Console 中执行以下代码临时修复样式：
```javascript
document.body.style.backgroundColor = '#ffffff';
document.body.style.color = '#000000';
document.querySelectorAll('*').forEach(el => {
  if (el.style.color === 'rgb(0, 0, 0)' || el.style.color === 'black') {
    el.style.color = '#000000';
  }
});
```

### 方式四：导出并离线分析
- **Trace 分析**
  ```bash
  curl http://127.0.0.1:6060/debug/pprof/trace?seconds=5 -o trace.out
  go tool trace trace.out
  ```
- **导出 pprof 数据**
  ```bash
  # 导出 CPU 数据
  curl http://127.0.0.1:6060/debug/pprof/profile?seconds=30 > cpu.pprof
  go tool pprof cpu.pprof
  
  # 导出内存数据
  curl http://127.0.0.1:6060/debug/pprof/heap > heap.pprof
  go tool pprof heap.pprof
  ```

### 其他
- **查看运行时变量**
  ```bash
  curl http://127.0.0.1:6060/debug/vars
  ```

# 更改日志
- [v0.9] 新增 pprof 性能分析工具
- [v0.9] 新增观测端口（默认 `127.0.0.1:6060`），集成 pprof/trace/expvar
- benchstat 对比示例：
  ```bash
  go test -bench=. -benchmem ./... > old.txt
  # 代码优化后
  go test -bench=. -benchmem ./... > new.txt
  benchstat old.txt new.txt
  ```