package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建服务器
func NewServer(ip string, port int) *Server {
	server := &Server{ip, port}
	return server
}

// 连接处理函数
func (this *Server) handleConnection(conn net.Conn) {
	// 处理连接的业务
	fmt.Println(("连接建立成功"))
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
