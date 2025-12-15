# HOW TO USE
```
    go build -o server server.go main.go user.go deadlock_check.go 
    ./server
```

# TIPS
```
    [v0.4] 提供用户查询在线用户功能
    v0.4 user.go [v0.4] 
        1. 用户查询在线用户
        2. 封装发送消息函数
    deadlock_check.go 提供了deadlock排查函数，定时打印所有goroutine堆栈信息
```