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
)

// Client 表示一个客户端连接
type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

// IClientLogic 管理所有客户端连接及消息路由
type IClientLogic struct {
	clients sync.Map // map[string]*Client
}

// NewIClientLogic 构造新的 ClientLogic 实例
func NewIClientLogic() *IClientLogic {
	return &IClientLogic{}
}
