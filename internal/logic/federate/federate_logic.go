// Package federate
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate_logic.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 09:01:35
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package federate

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"time"
)

// FederatedPeer 表示一个级联对端连接
type FederatedPeer struct {
	Addr string
	Conn *websocket.Conn
	// 带缓冲的发送 channel，避免因客户端写入速度较慢导致阻塞
	Send chan []byte
}

// AddPeer 添加新的级联对端
func (l *IFederateLogic) AddPeer(ctx context.Context, remoteAddr string, ws *websocket.Conn) error {
	if ws == nil {
		g.Log(consts.SocketLogger).Errorf(ctx, "invalid websocket connection")
		return errors.New("invalid websocket connection")
	}
	peer := &FederatedPeer{
		Addr: remoteAddr,
		Conn: ws,
		Send: make(chan []byte, 1024),
	}
	l.peers.Store(remoteAddr, peer)
	// 启动写入 goroutine 提高并发性能
	go l.WritePump(ctx, peer)
	return nil
}

// HandleFederateConnection 循环读取来自级联对端的消息，并进行广播处理
func (l *IFederateLogic) HandleFederateConnection(ctx context.Context, ws *websocket.Conn) {
	asyncCtx := context.WithoutCancel(ctx)
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			g.Log(consts.SocketLogger).Error(asyncCtx, err)
		}
	}(ws)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			l.RemovePeer(ws)
			break
		}
		// 此处可以对消息进行解包、鉴权、日志记录等处理，示例中直接广播给所有对端
		l.Broadcast(message)
	}
}

// WritePump 将待发送的消息写入到 websocket 连接，带有心跳检测
func (l *IFederateLogic) WritePump(ctx context.Context, peer *FederatedPeer) {
	asyncCtx := context.WithoutCancel(ctx)
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		err := peer.Conn.Close()
		if err != nil {
			g.Log().Error(asyncCtx, err)
			return
		}
	}()
	for {
		select {
		case message, ok := <-peer.Send:
			if !ok {
				if err := peer.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					g.Log().Error(asyncCtx, err)
					return
				}
				return
			}
			if err := peer.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
				g.Log().Error(asyncCtx, err)
				return
			}
			if err := peer.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				g.Log().Error(asyncCtx, err)
				return
			}
		case <-ticker.C:
			if err := peer.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
				g.Log().Error(asyncCtx, err)
				return
			}
			if err := peer.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				g.Log().Error(asyncCtx, err)
				return
			}
		}
	}
}

// Broadcast 向所有级联对端发送消息，采用非阻塞写入
func (l *IFederateLogic) Broadcast(message []byte) {
	l.peers.Range(func(key, value interface{}) bool {
		peer := value.(*FederatedPeer)
		select {
		case peer.Send <- message:
		default:
			// 如果对端发送 channel 满，直接丢弃消息，保证不阻塞整个广播流程
		}
		return true
	})
}

// RemovePeer 根据 websocket 连接移除对端
func (l *IFederateLogic) RemovePeer(ws *websocket.Conn) {
	l.peers.Range(func(key, value interface{}) bool {
		peer := value.(*FederatedPeer)
		if peer.Conn == ws {
			l.peers.Delete(key)
			close(peer.Send)
			return false
		}
		return true
	})
}
