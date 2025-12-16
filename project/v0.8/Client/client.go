package main

import (
	"fmt"
	"net"
)

type Client struct {
	Ip   string
	Port int
	name string
	conn net.Conn
}

func NewClient(ip string, port int) *Client {
	client := &Client{
		Ip:   ip,
		Port: port,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}

	client.conn = conn

	return client
}

func main() {
	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>>>>建立连接失败...")
		return
	}

	fmt.Println(">>>>>>>建立连接成功")

	// 启动客户端业务
	go func() {
		select {}
	}()

}
