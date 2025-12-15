# HOW TO USE
```
    go build -o server server.go main.go user.go deadlock_check.go 
    ./server
```

# TIPS
```
    [v0.5] 提供用户重命名功能
    v0.5 user.go [v0.5] 新增用户重命名功能
    deadlock_check.go 提供了deadlock排查函数，定时打印所有goroutine堆栈信息
```