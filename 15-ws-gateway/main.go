package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

// Hub 管理所有在线连接。
type Hub struct {
	mu      sync.RWMutex
	clients map[string]*Client
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*Client),
	}
}

func (h *Hub) Add(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if old, ok := h.clients[c.uid]; ok {
		old.Close("replaced by new connection")
	}
	h.clients[c.uid] = c
	log.Printf("uid=%s connected, online=%d", c.uid, len(h.clients))
}

func (h *Hub) Remove(uid string, reason string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.clients[uid]; ok {
		delete(h.clients, uid)
		log.Printf("uid=%s removed, reason=%s, online=%d", uid, reason, len(h.clients))
	}
}

func (h *Hub) Kick(uid, reason string) bool {
	h.mu.RLock()
	c, ok := h.clients[uid]
	h.mu.RUnlock()
	if !ok {
		return false
	}
	c.Close(reason)
	return true
}

func (h *Hub) Online() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.clients)
}

type Client struct {
	uid        string
	conn       *websocket.Conn
	hub        *Hub
	send       chan []byte
	lastPongMs atomic.Int64
	ctx        context.Context
	cancel     context.CancelFunc
	once       sync.Once
	remoteAddr string
}

func NewClient(uid string, conn *websocket.Conn, hub *Hub) *Client {
	ctx, cancel := context.WithCancel(context.Background())
	c := &Client{
		uid:        uid,
		conn:       conn,
		hub:        hub,
		send:       make(chan []byte, 16),
		ctx:        ctx,
		cancel:     cancel,
		remoteAddr: conn.RemoteAddr().String(),
	}
	c.lastPongMs.Store(time.Now().UnixMilli())
	return c
}

func (c *Client) Start() {
	c.conn.SetPongHandler(func(appData string) error {
		c.lastPongMs.Store(time.Now().UnixMilli())
		return nil
	})
	go c.readLoop()
	go c.writeLoop()
}

func (c *Client) Close(reason string) {
	c.once.Do(func() {
		c.cancel()
		_ = c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, reason), time.Now().Add(2*time.Second))
		_ = c.conn.Close()
		c.hub.Remove(c.uid, reason)
	})
}

func (c *Client) readLoop() {
	defer c.Close("read exit")
	c.conn.SetReadLimit(2 << 20) // 2MB
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if !errors.Is(err, websocket.ErrCloseSent) && !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Printf("uid=%s read error=%v", c.uid, err)
			}
			return
		}
		// 简单 echo，实际可改成路由到业务逻辑
		select {
		case c.send <- msg:
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) writeLoop() {
	defer c.Close("write exit")
	pingTicker := time.NewTicker(15 * time.Second)
	defer pingTicker.Stop()

	for {
		select {
		case msg := <-c.send:
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Printf("uid=%s write error=%v", c.uid, err)
				return
			}
		case <-pingTicker.C:
			if time.Since(time.UnixMilli(c.lastPongMs.Load())) > 30*time.Second {
				log.Printf("uid=%s heartbeat timeout", c.uid)
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("uid=%s ping error=%v", c.uid, err)
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	token := flag.String("token", "secret-token", "simple static token for demo")
	flag.Parse()

	hub := NewHub()
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]any{
			"online": hub.Online(),
			"time":   time.Now().Format(time.RFC3339),
		}
		_ = json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/kick", func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("uid")
		if uid == "" {
			http.Error(w, "uid required", http.StatusBadRequest)
			return
		}
		if ok := hub.Kick(uid, "manual kick"); !ok {
			http.Error(w, "uid not found", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "kicked %s\n", uid)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("uid")
		tk := r.URL.Query().Get("token")
		if uid == "" || tk == "" {
			http.Error(w, "uid and token required", http.StatusBadRequest)
			return
		}
		if tk != *token {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("upgrade error=%v", err)
			return
		}

		client := NewClient(uid, conn, hub)
		hub.Add(client)
		client.Start()
		log.Printf("uid=%s remote=%s handshake ok", uid, client.remoteAddr)
	})

	log.Printf("listening on %s (token=%s)", *addr, *token)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
