# HOW TO USE
```
    go build -o server server.go main.go user.go deadlock_check.go 
    ./server
```

# TIPS
```
    [v0.6] 提供用户超时踢出功能
    v0.6 user.go 
        [v0.6] 1. 用户上线业务封装
        [v0.6] 2. 用户下线业务封装
        [v0.6] 3. 用户超时踢出功能
        [v0.6] 5. 确保 Offline() 只执行一次，防止重复关闭 channel
    v0.6 server.go
        [v0.6] 4. 新增用户超时踢出功能
        
    deadlock_check.go 提供了deadlock排查函数，定时打印所有goroutine堆栈信息
```