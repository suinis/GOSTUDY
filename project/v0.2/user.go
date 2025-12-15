package main

import "net"

// 1. user类型
type User struct {
	Name string
	Addr string
	Ch   chan string
	Conn net.Conn
}

// 2. 创建用户的API NewUser
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	// 注意是指针
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		Ch:   make(chan string),
		Conn: conn,
	}

	// 4. 创建一个go程，监听user channel
	go user.ListenMessage()

	return user
}

// 3. 监听当前User channel的方法，一旦有消息，就直接发送给客户端 ListenMessage
func (this *User) ListenMessage() {
	for {
		msg := <-this.Ch

		this.Conn.Write([]byte(msg + "\n"))
	}
}
