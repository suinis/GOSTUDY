package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8080/ws?uid=U1&token=secret-token"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", msg)
		}
	}()

	for i := 0; i < 3; i++ {
		err := c.WriteMessage(websocket.TextMessage, []byte("hello"))
		if err != nil {
			log.Println("write:", err)
			return
		}
		time.Sleep(time.Second)
	}
	time.Sleep(5 * time.Second)
}
