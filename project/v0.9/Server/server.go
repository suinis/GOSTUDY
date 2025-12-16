package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// 5. 新增OnlineMap和Message channel, 涉及共享变量的注意需要加锁
type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	MapMutex  sync.RWMutex

	ChMessage chan string
}

// 创建服务器
// 6. 新增属性初始化
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		ChMessage: make(chan string),
	}
	return server
}

// 8. 新增广播消息函数BroadCast(User, msg)
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + " : " + msg

	this.ChMessage <- sendMsg
}

// 连接处理函数
// 7. 处理建立的客户端连接(广播上线消息)，并添加用户到对应map
// [v0.6] 4. 新增用户超时踢出功能
func (this *Server) handleConnection(conn net.Conn) {
	// 处理连接的业务
	// fmt.Println(("连接建立成功"))
	user := NewUser(conn, this)

	user.Online()

	isLive := make(chan bool)

	// [v0.3] 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err : ", err)
				return
			}

			// 提取用户消息(去掉'\n')
			msg := buf[:n-1]
			user.DoMessage(string(msg))

			isLive <- true
		}
	}()

	// 当前handler阻塞
	for {
		select {
		case <-isLive:
			// select执行时，会把所有条件都计算一遍，所以每次select都会重新创建一个time.After管道，
			// 超时后，可以从time.After管道中读取出数据，表示已经超时
		case <-time.After(300 * time.Second):
			// 超时处理
			user.TimeOutOffline()
			return
		}
	}
}

// 9. 与8对应，channel有发就有收，新增监听广播的方法ListenMessager
func (this *Server) ListenChMessage() {
	// deadlock_check()
	for {
		msg := <-this.ChMessage

		this.MapMutex.Lock()
		for _, clients := range this.OnlineMap {
			clients.Ch <- msg
		}
		this.MapMutex.Unlock()
	}
}

// 启动服务器
func (this *Server) Start() {

	// 监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}

	// listener close
	defer listener.Close()

	// 10. 服务器启动后就监听ChMessage管道
	go this.ListenChMessage()

	for {
		// accept
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("listener.Accept error: ", err)
			continue
		}

		// handler 连接
		go this.handleConnection(conn)
	}
}
