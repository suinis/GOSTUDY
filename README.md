## GO 学习笔记（基于个人 GOSTUDY 仓库）

- **仓库地址**：[`https://github.com/suinis/GOSTUDY`](https://github.com/suinis/GOSTUDY)  

---

## 0. 环境与整体结构

- **模块说明**
  - `go.mod` / `go.sum`：Go Modules 依赖与版本管理。
  - 各数字前缀目录（`1-var`、`2-const` … `15-ws-gateway`）代表一个知识专题。
  - `hello`：最基础的 Hello World 示例。
  - `project`：更接近实战的综合项目练习（包含版本子目录，如 `v0.7`、`v0.8`）。

- **学习顺序**
  1. `hello` 入门
  2. 基础语法：`1-var` → `2-const` → `3-function` → `4-init` → `5-defer` → `6-recover`
  3. 容器类型：`7-array` → `8-slice` → `9-map`
  4. 面向对象与抽象：`10-struct` → `11-interface` → `12-reflect`
  5. 并发与通信：`13-goroutine` → `14-channel`
  6. 综合实践：`15-ws-gateway`、`project`

---

## 1. Hello World 与 Go 基本结构

- **对应目录**：[`hello`](https://github.com/suinis/GOSTUDY/tree/main/hello)

### 1.1 Go 程序基本骨架

- **关键点**
  - 每个可执行程序必须有 `package main`。
  - 入口函数固定为 `func main()`，无参数无返回值。
  - 使用 `import` 导入标准库或第三方包。
  - 使用 `go run` 快速运行；使用 `go build` 生成可执行文件。

---

## 2. 变量与类型：`1-var`

- **对应目录**：[`1-var`](https://github.com/suinis/GOSTUDY/tree/main/1-var)

### 2.1 变量声明方式

- **关键语法**
  - 显式声明：`var x int` / `var x int = 10`
  - 类型推断：`var x = 10`
  - 短变量声明：`x := 10`
  - 多变量声明：`var a, b int = 1, 2` 或 `a, b := 1, 2`

### 2.2 基本数据类型

- **整数类型**：`int` / `int8` / `int16` / `int32` / `int64`，以及对应的 `uint` 系列。
- **浮点类型**：`float32` / `float64`
- **布尔类型**：`bool`
- **字符串**：UTF-8 字符串；注意按字节和按 rune（字符）遍历的区别。
- **其它**：`byte`（`uint8` 别名）、`rune`（`int32` 别名）。

### 2.3 变量作用域

- **函数内局部变量**：只在函数体内可见。
- **包级变量**：在同一包内所有文件共享。
- **命名规范**：
  - 首字母 **大写**：对外导出（包外可见）。
  - 首字母 **小写**：包内可见。

---

## 3. 常量：`2-const`

- **对应目录**：[`2-const`](https://github.com/suinis/GOSTUDY/tree/main/2-const)

### 3.1 常量基本用法

- 使用 `const` 声明，编译期确定，运行时不可修改：
  - `const Pi = 3.1415`
  - `const MaxConnections int = 100`

### 3.2 无类型常量与类型安全

- **无类型常量** 可以根据上下文自动适配类型（例如参与不同类型表达式运算）。

### 3.3 iota 枚举

- 常用来定义枚举值：
  - `const ( A = iota; B; C )`，自动递增。
- 在状态机、错误码、类型标识等场景非常实用。

---

## 4. 函数与多返回值：`3-function`

- **对应目录**：[`3-function`](https://github.com/suinis/GOSTUDY/tree/main/3-function)

### 4.1 函数定义与调用

- **关键特性**
  - 强类型、参数与返回值需要显式声明类型。
  - 支持多返回值：例如 `(value int, err error)`。
  - 支持具名返回值：在 `return` 处可省略变量名，但不建议滥用。

### 4.2 不定参数与切片

- 不定参数：`func sum(nums ...int) int`，在函数体中 `nums` 为切片。
- 调用时可以直接传多个参数或使用 `slice...` 进行展开。

### 4.3 函数作为一等公民

- 函数可以：
  - 作为参数传入其他函数；
  - 作为返回值；
  - 赋值给变量（`type Handler func(int) error`）。

---

## 5. init 函数与包初始化：`4-init`

- **对应目录**：[`4-init`](https://github.com/suinis/GOSTUDY/tree/main/4-init)

### 5.1 init 函数语义

- 每个文件可以有 **多个** `init` 函数，编译时按出现顺序执行。
- 整体初始化顺序：
  1. 导入依赖包（深度优先）。
  2. 初始化包级变量。
  3. 执行各文件中的 `init()`。
  4. 最后执行 `main.main()`。

### 5.2 常见使用场景

- 注册日志、监控、路由、驱动等。
- 初始化包级缓存、连接池等。

> **注意**：避免在 `init` 中做过于复杂或有副作用的逻辑，容易降低可测试性与可维护性。

---

## 6. defer 延迟调用：`5-defer`

- **对应目录**：[`5-defer`](https://github.com/suinis/GOSTUDY/tree/main/5-defer)

### 6.1 基本语义

- 在函数返回前 **延迟执行** 已注册的 `defer` 调用。
- 多个 `defer` 以后进先出（LIFO）顺序执行。

### 6.2 典型应用

- 资源清理：`defer file.Close()`、`defer conn.Close()`。
- 锁释放：`defer mu.Unlock()`。
- 指标采集：在函数首部记录开始时间，`defer` 中计算耗时并上报。

### 6.3 性能与坑

- 每次 `defer` 有一定开销，高频路径下需注意。
- `defer` 捕获的是当时的参数值（值传递），而不是之后被修改的变量状态。

---

## 7. 错误处理与 panic/recover：`6-recover`

- **对应目录**：[`6-recover`](https://github.com/suinis/GOSTUDY/tree/main/6-recover)

### 7.1 Go 的错误处理哲学

- 推荐使用 **显式错误返回**：`value, err := foo(); if err != nil { ... }`
- 保持调用链上的错误可见与可追踪，而不是异常抛来抛去。

### 7.2 panic 与 recover

- `panic`：导致当前 goroutine 中断，向上展开调用栈，依次执行 `defer`，最终程序崩溃（如果未被 `recover`）。
- `recover`：
  - 必须在 `defer` 函数中调用才生效。
  - 可以捕获 `panic`，使程序继续运行。
- 使用原则：
  - **只在不可恢复的程序错误 / 严重状态下使用 `panic`**。
  - 对业务错误（如参数非法、用户不存在），一律使用 `error`。

---

## 8. 数组：`7-array`

- **对应目录**：[`7-array`](https://github.com/suinis/GOSTUDY/tree/main/7-array)

### 8.1 数组特性

- **长度是类型的一部分**：`[3]int` 与 `[4]int` 是不同类型。
- 声明与初始化：
  - `var a [3]int`
  - `b := [3]int{1, 2, 3}`
  - `c := [...]int{1, 2, 3}` 编译期自动推断长度。

### 8.2 使用场景

- 固定大小的缓冲区、底层实现、性能敏感的代码中。
- 一般业务逻辑更多使用 **切片**（slice）。

---

## 9. 切片：`8-slice`

- **对应目录**：[`8-slice`](https://github.com/suinis/GOSTUDY/tree/main/8-slice)

### 9.1 切片的本质

- 一个轻量级描述符，包含：
  - 指针（指向底层数组）。
  - 长度 `len`。
  - 容量 `cap`。

### 9.2 创建与扩容

- 使用字面量：`s := []int{1, 2, 3}`
- 使用 `make`：`s := make([]int, 0, 10)`  
- 使用 `append` 扩容：
  - 当 `len < cap`：复用原数组。
  - 当 `len == cap`：分配新数组并拷贝。

### 9.3 切片分享底层数组的坑

- 拷贝切片变量只是拷贝描述符，底层数组仍然共享。
- 在函数间传递、子切片操作时要谨慎，避免“不小心修改到别人”的内存。

---

## 10. Map：`9-map`

- **对应目录**：[`9-map`](https://github.com/suinis/GOSTUDY/tree/main/9-map)

### 10.1 基本用法

- 创建：
  - `m := make(map[string]int)`
  - `m := map[string]int{"a": 1, "b": 2}`
- 访问与插入：
  - `m["key"] = 1`
  - `v, ok := m["key"]`，`ok` 判断是否存在。

### 10.2 注意事项

- 遍历 `range map` 是 **无序** 的，不能依赖遍历顺序。
- `nil map` 上读不会 panic，写会 panic。
- map 不是并发安全的，**多个 goroutine 并发写必须加锁或用 sync.Map**。

---

## 11. 结构体 struct：`10-struct`

- **对应目录**：[`10-struct`](https://github.com/suinis/GOSTUDY/tree/main/10-struct)

### 11.1 定义与实例化

- 定义：`type User struct { ID int; Name string }`
- 实例化：
  - 字面量：`u := User{ID: 1, Name: "Tom"}`
  - `new`：`u := new(User)`，得到指针。
  - 取地址：`u := &User{ID: 1}`。

### 11.2 方法与接收者

- 值接收者：`func (u User) Do()`，调用时会拷贝。
- 指针接收者：`func (u *User) Do()`，可修改原对象，避免拷贝。
- 指针 / 值接收者混用规则与接口实现有密切关系（配合下一节 `interface` 理解）。

### 11.3 嵌套与“伪继承”

- 匿名字段嵌套：`type Admin struct { User; Level int }`
- 外层类型可以直接访问内嵌字段，形成一种“组合 + 继承”风格的复用。

---

## 12. 接口 interface：`11-interface`

- **对应目录**：[`11-interface`](https://github.com/suinis/GOSTUDY/tree/main/11-interface)

### 12.1 隐式实现

- 无需显式 `implements`，只要某类型实现了接口中的 **所有方法**，就被视为实现。
- **小接口原则**：将接口拆分为多个小的、单一职责接口（如 `io.Reader` / `io.Writer`）。

### 12.2 动态分发与多态

- 接口值由“具体类型 + 具体值”两部分构成。
- 方法调用由运行时动态选择具体实现。

### 12.3 接口与指针 / 值接收者

- 使用指针接收者实现接口时，**只有指针类型的值** 才实现该接口。
- 在设计 API 时，要考虑调用方是持有值还是持有指针。

---

## 13. 反射 reflect：`12-reflect`

- **对应目录**：[`12-reflect`](https://github.com/suinis/GOSTUDY/tree/main/12-reflect)

### 13.1 反射基础类型

- `reflect.TypeOf`：获取静态类型信息。
- `reflect.ValueOf`：操作运行时值（读、写、调用方法）。

### 13.2 典型应用

- 通用序列化 / 反序列化（如 JSON）。
- 通用校验（tag 驱动，如 `validate:"required"`）。
- 日志打印、自动诊断等。

> **注意**：反射写法较为复杂，运行时开销也更大，**不要在性能敏感路径滥用**。

---

## 14. Goroutine 并发：`13-goroutine`

- **对应目录**：[`13-goroutine`](https://github.com/suinis/GOSTUDY/tree/main/13-goroutine)

### 14.1 Goroutine 特性

- 使用 `go func()` 启动轻量级线程。
- 由 Go 运行时负责管理调度，而非 OS 线程。
- 非常适合处理 I/O 密集型任务和高并发业务。

### 14.2 常见模式

- **并发任务分发**：循环中 `go` 多个 worker。
- **超时与取消**：结合 `context.Context` 使用。

### 14.3 Goroutine 泄漏与调试

- 忘记退出的 goroutine 会一直占用内存 / 阻塞资源。
- 在复杂项目中，后续会结合 `pprof` 查看 goroutine profile，排查泄漏。

---

## 15. Channel 通信：`14-channel`

- **对应目录**：[`14-channel`](https://github.com/suinis/GOSTUDY/tree/main/14-channel)

### 15.1 Channel 基础

- 无缓冲 channel：`ch := make(chan int)`，发送 / 接收会互相阻塞。
- 有缓冲 channel：`ch := make(chan int, 10)`，缓冲满 / 空时才阻塞。
- 关闭操作：`close(ch)`；读取已关闭的 channel 会立即返回零值并 `ok == false`。

### 15.2 select 多路复用

- `select` 允许同时等待多个 channel 事件。
- 常见用法：
  - 多种输入竞争（谁先到用谁）。
  - 结合 `time.After` 实现超时。
  - 使用 `default` 实现非阻塞操作。

### 15.3 常见并发模式

- **生产者-消费者** 模式。
- **worker pool**：固定数量的 worker 消费任务队列。
- **扇出 / 扇入**：多个 goroutine 并发处理，再汇总结果。

---

## 16. WebSocket 网关实践：`15-ws-gateway`

- **对应目录**：[`15-ws-gateway`](https://github.com/suinis/GOSTUDY/tree/main/15-ws-gateway)

### 16.1 模块角色

- **服务端网关**：维护 WebSocket 连接，做消息转发 / 路由。
- **客户端**：与网关建立长连接，收发消息。
- **协程模型与 channel 设计**：用 goroutine 处理连接，用 channel 在内部传递消息事件。

### 16.2 关键点

- 使用 `net/http` + WebSocket 实现长连接。
- 对每个客户端连接创建 goroutine 进行读写。
- 使用 channel 把读取到的消息交给内部逻辑处理，避免在读写 goroutine 中做复杂业务。

---

## 17. 综合项目：`project`

- **对应目录**：[`project`](https://github.com/suinis/GOSTUDY/tree/main/project)

### 17.1 版本演进

- version：
  - `v0.1`：基础server服务构建
  - `v0.2`：提供用户私聊功能
  - `v0.3`：提供用户消息广播
  - `v0.4`：提供用户查询在线用户功能
  - `v0.5`：提供用户重命名功能
  - `v0.6`：提供用户超时踢出功能
  - `v0.7`：提供用户私聊功能
  - `v0.8`：自定义客户端连接(half to do)
  - `v0.9`：嵌入pprof等性能分析工具（to do）

### 17.2 与工具链结合

- 尝试嵌入性能分析、调试工具，例如：
  - `net/http/pprof`
  - 慢日志 / tracing
  - 自定义 metrics（如 Prometheus）

---

## 18. TODO：Go 服务端工具链与性能分析（pprof 等）

> 本节正在学习中

### 18.1 基础目标

- **TODO-1：掌握 pprof 基础用法**
  - 在简单 HTTP 服务或 `15-ws-gateway` 项目中引入 `net/http/pprof`。
  - 采集：
    - CPU profile
    - Heap profile
    - Goroutine profile
  - 使用：
    - `go tool pprof` 分析热点函数 / 卡点。
    - web 界面查看调用图、火焰图（需要配套工具）。

- **TODO-2：在综合项目中定位性能瓶颈**
  - 在 `project/v0.9` 中选一个典型接口，压测前后对比：
    - QPS / 延迟（p99, p999）
    - CPU 占用
    - 内存使用 / GC 次数
  - 分析：
    - 哪些函数占用 CPU 时间最多。
    - 是否存在 goroutine 泄漏（goroutine profile 明显上升）。
    - 是否有频繁分配内存的热点（heap / allocs profile）。

- **TODO-3：故障排查与稳定性提升**
  - 模拟典型问题：
    - 死锁 / 活锁（可参考 `project/v0.8/Server/deadlock_check.go` 的相关逻辑思路）。
    - 阻塞 channel / 阻塞 I/O。
  - 利用：
    - goroutine dump（`runtime.Stack` / `pprof` goroutine）
    - 日志 + 监控指标，定位问题位置。
  - 结合分析结果，优化：
    - goroutine 生命周期管理。
    - channel 缓冲大小与关闭时机。
    - 锁的粒度和使用方式。

### 18.2 延伸工具链 TODO

- **TODO-4：学习并使用以下工具**
  - 日志：结构化日志（如 zap、logrus）。
  - Metrics：Prometheus 指标采集与监控。
  - Tracing：分布式追踪（OpenTelemetry 等）。
  - 静态分析：
    - `go vet`
    - `golangci-lint` 等工具。

- **TODO-5：将工具链沉淀为模板**
  - 为后续新项目准备一个“服务端脚手架”：
    - 基础 HTTP 服务 / WebSocket 入口。
    - 统一日志组件。
    - pprof / metrics / health check 一体化。
    - 标准的配置管理与优雅关闭（graceful shutdown）逻辑。

---

