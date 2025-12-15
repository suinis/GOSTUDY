package main

import (
	"net"
)

// 1. user类型
type User struct {
	Name string
	Addr string
	Ch   chan string
	conn net.Conn

	server *Server
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
func (this *User) ListenMessage() {
	for {
		msg := <-this.Ch

		this.conn.Write([]byte(msg + "\n"))
	}
}

// [v0.4] 2. 封装发送消息函数
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// [v0.4] 1. 用户查询在线用户
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
	} else {
		this.server.BroadCast(this, msg)
	}
}
