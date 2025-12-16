package main

import (
	"net"
	"strings"
	"sync"
)

// 1. user类型
type User struct {
	Name string
	Addr string
	Ch   chan string
	conn net.Conn

	server *Server

	// [v0.6] 5. 确保 Offline() 只执行一次，防止重复关闭 channel
	offlineOnce sync.Once
}

// 2. 创建用户的API NewUser
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	// 注意是指针
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		Ch:     make(chan string),
		conn:   conn,
		server: server,
	}

	// 4. 创建一个go程，监听user channel
	go user.ListenMessage()

	return user
}

// 3. 监听当前User channel的方法，一旦有消息，就直接发送给客户端 ListenMessage
// [v0.7] 3. 使用 range 在 channel 关闭时自动退出循环，避免空写入
func (this *User) ListenMessage() {
	// 使用 range 在 channel 关闭时自动退出循环，避免空写入
	// for {
	// 	msg := <-this.Ch
	// }
	for msg := range this.Ch {
		this.conn.Write([]byte(msg + "\n"))
	}
}

// [v0.4] 2. 封装发送消息函数
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// [v0.4] 1. 用户查询在线用户
// [v0.7] 1. 新增用户私聊功能
func (this *User) DoMessage(msg string) {
	if msg == "search" {
		msg := "当前在线用户："

		this.server.MapMutex.Lock()
		for name, _ := range this.server.OnlineMap {
			msg += "[" + name + "] "
		}
		this.server.MapMutex.Unlock()

		msg += "\n"

		this.SendMsg(msg)
	} else if len(msg) > 7 && msg[:7] == "rename|" { // [v0.5] 新增用户重命名功能
		newName := strings.Split(msg, "|")[1]

		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("用户名已存在")
		} else {
			this.server.MapMutex.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.MapMutex.Unlock()

			this.Name = newName
			this.SendMsg("您已修改用户名为：" + newName + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式：to|name|content
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式有误：请按照 \"to|name|content\" \n")
			return
		}

		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("用户不存在\n")
			return
		}

		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重发\n")
			return
		}
		remoteUser.SendMsg(this.Name + "对您说：" + content + "\n")
	} else {
		this.server.BroadCast(this, msg)
	}
}

// [v0.6] 1. 用户上线业务封装
func (this *User) Online() {
	this.server.MapMutex.Lock()
	this.server.OnlineMap[this.Addr] = this
	this.server.MapMutex.Unlock()

	this.server.BroadCast(this, "上线")
}

// [v0.6] 2. 用户下线业务封装
/*
	[v0.7] 2. 调整delete OnlineMap和close this.ch的顺序，
				Offline 里先广播再关闭用户 channel，ListenChMessage 迭代 OnlineMap 时仍能取到已关闭的 clients.Ch，导致发送时 panic。
				现在先移除用户再广播，避免向已关闭 channel 发送。
*/
func (this *User) Offline() {
	// 子协程接收到0，执行Offline和主循环TimeOutOffline会导致重复关闭问题
	// 通过sync.Once确保Offline之执行一次
	this.offlineOnce.Do(func() {
		// [v0.7] 3. 先从在线列表移除，避免广播时向已关闭的 channel 发送
		this.server.MapMutex.Lock()
		delete(this.server.OnlineMap, this.Name)
		this.server.MapMutex.Unlock()

		this.server.BroadCast(this, "下线")

		close(this.Ch)
		this.conn.Close()
	})
}

// [v0.6] 3. 用户超时踢出功能
func (this *User) TimeOutOffline() {
	this.conn.Write([]byte("您被踢了"))
	this.Offline()
}
