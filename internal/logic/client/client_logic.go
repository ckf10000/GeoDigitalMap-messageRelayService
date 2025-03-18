// Package client
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client_logic.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 11:06:56
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package client

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"time"
)

// AddClient 添加新的客户端，clientID 建议使用全局唯一标识
func (l *IClientLogic) AddClient(ctx context.Context, clientID string, ws *websocket.Conn) error {
	if ws == nil || clientID == "" {
		return errors.New("invalid client connection or ID")
	}
	client := &Client{
		ID:   clientID,
		Conn: ws,
		Send: make(chan []byte, 1024),
	}
	l.clients.Store(clientID, client)
	g.Log(consts.SocketLogger).Infof(ctx, "Client: %s has joined", clientID)
	// 启动写入 goroutine 实现非阻塞消息发送
	go l.WritePump(ctx, client)
	return nil
}

// HandleMessages 循环读取并解析客户端消息，交由底层逻辑进行路由
func (l *IClientLogic) HandleMessages(ctx context.Context, ws *websocket.Conn) {
	asyncCtx := context.WithoutCancel(ctx)
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			g.Log(consts.SocketLogger).Error(asyncCtx, err)
		}
		//g.Log(consts.SocketLogger).Info(ctx, "websocket connect closed")
	}(ws)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			// 检查是否是客户端主动关闭连接
			var closeErr *websocket.CloseError
			if errors.As(err, &closeErr) {
				switch closeErr.Code {
				case websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived:
					g.Log(consts.SocketLogger).Warningf(ctx, "Client closed the connection")
				default:
					g.Log(consts.SocketLogger).Errorf(ctx, "client closed the connection with unexpected code: %+v", closeErr.Code)
				}
			}
			l.RemoveClient(ws)
			break
		}
		g.Log(consts.SocketLogger).Infof(ctx, "received handle: %s", message)

		// 将收到的消息原样回传
		//if err = ws.WriteMessage(msgType, message); err != nil {
		//	break
		//}

		var msg dto.MessageOutputDTO
		if err = json.Unmarshal(message, &msg); err != nil {
			// 解析出错，记录日志或返回错误响应
			g.Log(consts.SocketLogger).Error(ctx, err)
			continue
		}
		l.RouteMessage(ctx, msg)
	}
}

// RemoveClient 根据 websocket 连接移除客户端
func (l *IClientLogic) RemoveClient(ws *websocket.Conn) {
	var targetID string
	l.clients.Range(func(key, value interface{}) bool {
		client := value.(*Client)
		if client.Conn == ws {
			targetID = client.ID
			return false
		}
		return true
	})
	if targetID != "" {
		l.clients.Delete(targetID)
	}
}

// RouteMessage 根据消息类型进行路由分发
func (l *IClientLogic) RouteMessage(ctx context.Context, msg dto.MessageOutputDTO) {
	data, err := json.Marshal(msg)
	if err != nil {
		g.Log(consts.SocketLogger).Error(ctx, err)
		return
	}
	switch msg.MessageType {
	case consts.PointToPoint, consts.PointToGroup:
		for _, targetID := range msg.Receivers {
			if v, ok := l.clients.Load(targetID); ok {
				client := v.(*Client)
				select {
				case client.Send <- data:
				default:
					// 如果 channel 满则丢弃该消息
				}
			}
		}
	case consts.Broadcast:
		l.clients.Range(func(key, value interface{}) bool {
			client := value.(*Client)
			select {
			case client.Send <- data:
			default:
				// 丢弃消息以保证广播不被阻塞
			}
			return true
		})
	}
}

// WritePump 将消息写入客户端连接，附带心跳检测
func (l *IClientLogic) WritePump(ctx context.Context, client *Client) {
	asyncCtx := context.WithoutCancel(ctx)
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		err := client.Conn.Close()
		if err != nil {
			g.Log(consts.SocketLogger).Error(asyncCtx, err)
			return
		}
	}()
	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				if err := client.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					g.Log(consts.SocketLogger).Error(asyncCtx, err)
					return
				}
				return
			}
			if err := client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
				g.Log(consts.SocketLogger).Error(asyncCtx, err)
				return
			}
			if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				g.Log(consts.SocketLogger).Error(asyncCtx, err)
				return
			}
		case <-ticker.C:
			if err := client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
				g.Log(consts.SocketLogger).Error(asyncCtx, err)
				return
			}
			if err := client.Conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				//g.Log(consts.SocketLogger).Error(asyncCtx, err)
				return
			}
		}
	}
}

// GetAllClients 返回所有在线客户端的 ID 列表，便于管理监控
func (l *IClientLogic) GetAllClients() []string {
	var clientIDs []string
	l.clients.Range(func(key, _ interface{}) bool {
		clientIDs = append(clientIDs, key.(string))
		return true
	})
	return clientIDs
}
