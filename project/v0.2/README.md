# HOW TO USE
```
    go build -o server server.go main.go user.go deadlock_check.go 
    ./server
```

# TIPS
```
    deadlock_check.go 提供了deadlock排查函数，定时打印所有goroutine堆栈信息
```