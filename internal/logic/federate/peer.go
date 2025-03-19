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
	Addr string
	Conn *websocket.Conn
	// 带缓冲的发送 channel，避免因客户端写入速度较慢导致阻塞
	Send chan []byte
}
