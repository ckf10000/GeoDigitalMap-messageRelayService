// Package websocket
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     hub.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-11 21:26:01
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package websocket

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

// upgrader WebSocket连接升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源连接
	},
}

// Client 客户端连接结构
type Client struct {
	ID      string
	Connect *websocket.Conn
	Send    chan []byte
}

// Hub 连接管理 Hub
type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mu         sync.Mutex
}

func NewWebSocketHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

// Run 处理 WebSocket 连接
func (h *Hub) Run(ctx context.Context) {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.ID] = client
			h.mu.Unlock()
			g.Log().Info(ctx, "新连接加入")
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.ID]; ok {
				delete(h.clients, client.ID)
				close(client.Send)
				err := client.Connect.Close()
				if err != nil {
					g.Log().Error(ctx, err)
					return
				}
			}
			h.mu.Unlock()
			g.Log().Info(ctx, "连接断开")
		case message := <-h.broadcast:
			h.mu.Lock()
			for _, client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client.ID)
				}
			}
			h.mu.Unlock()
		}
	}
}

func (h *Hub) ServerWs(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	connect, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		str := fmt.Sprintf("WebSocket 连接升级失败: %s", err)
		g.Log().Error(ctx, str)
		return
	}

	// 从 URL 获取客户端 ID
	clientID := r.URL.Query().Get("id")
	if clientID == "" {
		clientID = connect.RemoteAddr().String()
	}

	client := &Client{
		ID:      clientID,
		Connect: connect,
		Send:    make(chan []byte, 256),
	}
	h.register <- client

	go h.ReadPump(ctx, client)
	go h.WritePump(ctx, client)

}

// ReadPump 读取客户端消息
func (h *Hub) ReadPump(ctx context.Context, client *Client) {
	defer func() {
		h.unregister <- client
		err := client.Connect.Close()
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}()

	for {
		_, message, err := client.Connect.ReadMessage()
		if err != nil {
			str := fmt.Sprintf("读取消息错误: %s", err)
			g.Log().Error(ctx, str)
			break
		}
		// 广播消息
		h.broadcast <- message
	}
}

// WritePump 向客户端发送消息
func (h *Hub) WritePump(ctx context.Context, client *Client) {
	for msg := range client.Send {
		err := client.Connect.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			str := fmt.Sprintf("发送消息错误: %s", err)
			g.Log().Error(ctx, str)
			break
		}
	}
}
