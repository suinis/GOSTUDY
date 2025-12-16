# HOW TO USE
```
    go build -o server server.go main.go user.go deadlock_check.go 
    ./server
```

# TIPS
```
    [v0.7] 提供用户私聊功能
    v0.7 user.go 
        [v0.7] 1. 新增用户私聊功能
        [v0.7] 2. 调整delete OnlineMap和close this.ch的顺序
        [v0.7] 3. 先从在线列表移除，避免广播时向已关闭的 channel 发送
                  使用 range 在 channel 关闭时自动退出循环，避免空写入
        
    deadlock_check.go 提供了deadlock排查函数，定时打印所有goroutine堆栈信息
```