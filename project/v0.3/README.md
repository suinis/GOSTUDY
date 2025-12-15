# HOW TO USE
```
    go build -o server server.go main.go user.go deadlock_check.go 
    ./server
```

# TIPS
```
    [v0.3] 提供用户消息广播
    v0.3 新增server.go [v0.3] 接收客户端发送的消息
    deadlock_check.go 提供了deadlock排查函数，定时打印所有goroutine堆栈信息
```