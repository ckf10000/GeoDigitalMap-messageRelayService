// Package federate
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     peer.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-18 17:44:08
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package federate

import "github.com/gorilla/websocket"

// Peer 表示一个级联对端连接
type Peer struct {
	Addr string          // 用IP地址作为对端的唯一标识
	Conn *websocket.Conn // webksocket连接通道
	Send chan []byte     // 带缓冲的发送 channel，避免因客户端写入速度较慢导致阻塞
}
