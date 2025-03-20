// Package client
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 11:06:45
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package client

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

// Client 表示一个客户端连接
type Client struct {
	UserID   string          // 客户端用户唯一标识
	ClientID string          // 客户端唯一标识
	Conn     *websocket.Conn // WebSocket 连接
	Send     chan []byte     // 带缓冲的发送通道
	JoinAt   time.Time       // 客户端接入时间`
}

// IClientLogic 管理所有客户端连接及消息路由
type IClientLogic struct {
	clients     map[string]*Client // 以 clientID 为 key 的客户端对象字典
	userClients map[string]*Client // 以 userID 为 key 的客户端对象字典，避免消息路由时，每次都遍历 clients
	mu          sync.RWMutex       // 读写锁
}

// NewIClientLogic 构造新的 ClientLogic 实例
func NewIClientLogic() *IClientLogic {
	return &IClientLogic{
		clients:     make(map[string]*Client),
		userClients: make(map[string]*Client),
	}
}
