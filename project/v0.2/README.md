# HOW TO USE
```
    go build -o server server.go main.go user.go deadlock_check.go 
    ./server
```

# TIPS
```
    [v0.2] 提供用户私聊功能
    v0.2 user.go server.go deadlock_check.go
        [v0.2] 1. user类型
        [v0.2] 2. 创建用户的API NewUser
        [v0.2] 3. 监听当前User channel的方法，一旦有消息，就直接发送给客户端 ListenMessage
        [v0.2] 4. 创建一个go程，监听user channel
        [v0.2] 5. 新增OnlineMap和Message channel, 涉及共享变量的注意需要加锁
        [v0.2] 6. 新增属性初始化
        [v0.2] 7. 处理建立的客户端连接(广播上线消息)，并添加用户到对应map
        [v0.2] 8. 新增广播消息函数BroadCast(User, msg)
        [v0.2] 9. 与8对应，channel有发就有收，新增监听广播的方法ListenMessager

    deadlock_check.go 提供了deadlock排查函数，定时打印所有goroutine堆栈信息
```