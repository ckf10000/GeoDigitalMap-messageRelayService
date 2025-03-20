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
	"GeoDigitalMap-messageRelayService/internal/common/utils"
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/logic/manager"
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

// WritePump 方法负责从 Send 通道中读取消息并发送给客户端，带有心跳检测
func (c *Client) WritePump(ctx context.Context) {
	asyncCtx := context.WithoutCancel(ctx)
	ticker := time.NewTicker(consts.ClientHeartbeatDuration)
	defer func() {
		ticker.Stop()
		err := c.Conn.Close()
		if err != nil {
			g.Log(consts.SocketLogger).Error(asyncCtx, err)
			return
		} // 关闭连接
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 通道关闭，退出循环
				return
			}

			// 发送消息给客户端
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				g.Log(consts.SocketLogger).Errorf(ctx, "Write message error: %+v", err)
				return
			}
		}
	}
}

// AddClient 添加新的客户端，clientID 建议使用全局唯一标识
func (l *IClientLogic) AddClient(ctx context.Context, clientID string, conn *websocket.Conn) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.clients) >= consts.SocketClientCount {
		str := fmt.Sprintf("Exceeded the maximum connection limit of %d client devices", consts.SocketClientCount)
		g.Log(consts.SocketLogger).Error(ctx, str)
		return gerror.New(str)
	}

	// 创建带缓冲的发送通道
	client := &Client{
		ClientID: clientID,
		Conn:     conn,
		Send:     make(chan []byte, consts.MessageChannelSize), // 缓冲区大小可根据实际需求调整
	}

	l.clients[clientID] = client

	// 启动一个 goroutine 处理消息发送
	go client.WritePump(ctx)

	g.Log(consts.SocketLogger).Infof(ctx, "New client connected: %s", clientID)
	return nil
}

// HandleMessages 循环读取并解析客户端消息，交由底层逻辑进行路由
func (l *IClientLogic) HandleMessages(ctx context.Context, conn *websocket.Conn) {
	asyncCtx := context.WithoutCancel(ctx)
	clientID := conn.RemoteAddr().String()

	// 客户端断开时移除
	defer func() {
		l.RemoveClient(asyncCtx, clientID)
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			// 检查是否是Peer主动关闭连接
			var closeErr *websocket.CloseError
			if errors.As(err, &closeErr) {
				switch closeErr.Code {
				case websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived:
					g.Log(consts.SocketLogger).Warningf(asyncCtx, "Peer closed the connection")
				default:
					g.Log(consts.SocketLogger).Errorf(asyncCtx, "client closed the connection with unexpected code: %+v", closeErr.Code)
				}
			} else {
				g.Log(consts.SocketLogger).Infof(asyncCtx, "received handle: %s", message)
			}
			break
		}

		// 此时读出来的数据，还是[]byte类型数据
		//g.Log(consts.SocketLogger).Infof(asyncCtx, "received message: %+v", message)

		// 处理消息逻辑
		var msg *dto.MessageIutputDTO
		if err = json.Unmarshal(message, &msg); err != nil {
			g.Log(consts.SocketLogger).Errorf(ctx, "Failed to parse message: %+v", err)
			continue
		}

		// 服务端生成 UUID
		msg.UID = uuid.New().String()

		infoData, _ := gjson.Marshal(msg)
		//infoData, _ := gjson.MarshalIndent(msg, "", "\t")
		g.Log(consts.SocketLogger).Infof(asyncCtx, "received message: %s", string(infoData))

		l.RouteMessage(ctx, msg)
		// TODO 消息持久化

	}
}

// RemoveClient 根据 websocket 连接移除客户端
func (l *IClientLogic) RemoveClient(ctx context.Context, clientID string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if client, exists := l.clients[clientID]; exists {
		close(client.Send) // 关闭通道
		delete(l.clients, clientID)
		g.Log(consts.SocketLogger).Infof(ctx, "Client disconnected: %s", clientID)
	}
}

// RouteMessage 根据消息类型进行路由分发
func (l *IClientLogic) RouteMessage(ctx context.Context, msg *dto.MessageIutputDTO) {
	messageOutputDTO := &dto.MessageOutputDTO{
		MessageID:   msg.MessageID,
		Sender:      msg.Sender,
		Receivers:   msg.Receivers,
		Content:     msg.Content,
		CreatedAt:   msg.CreatedAt.Format(time.RFC3339),
		MessageType: msg.MessageType,
	}
	data, err := json.Marshal(messageOutputDTO)
	if err != nil {
		g.Log(consts.SocketLogger).Error(ctx, err)
		return
	}
	switch msg.MessageType {
	case consts.PointToPoint, consts.PointToGroup:
		l.SendP2GMessage(ctx, msg.Receivers, data)
	case consts.Broadcast:
		l.SendBroadcastMessage(ctx, data)
		// 判断Federate的peer列表是否为空
		if len(manager.GetAllPeerAddrs()) != 0 {
			// 将消息广播给所有 federate 的 peer
			localIP := utils.GetFederateLocalIP(ctx)
			if len(localIP) > 0 {
				manager.GetFederatePeerLogic().SendRelayMessage(ctx, data, []string{localIP})
			}
		}
	// TODO 目前仅支持广播消息的中继
	default:
		g.Log(consts.SocketLogger).Errorf(ctx, "Unsupported message types: %s", msg.MessageType)
	}

	// 持久化消息
	//SaveMessage(ctx, sender, receiver, "private", message)
}

// SendP2PMessage 发送私聊
func (l *IClientLogic) SendP2PMessage(ctx context.Context, receiver string, message []byte) error {
	l.mu.RLock()
	receiverClient, exists := l.clients[receiver]
	l.mu.RUnlock()

	if !exists {
		str := fmt.Sprintf("receiver <%s> not found", receiver)
		g.Log(consts.SocketLogger).Error(ctx, str)
		return gerror.New(str)
	}

	// 将消息写入客户端的 Send 通道
	select {
	case receiverClient.Send <- message:
		// 消息成功写入通道
	default:
		// 通道已满，丢弃消息或记录日志
		g.Log(consts.SocketLogger).Warningf(ctx, "Send channel is full, message dropped for client: %s", receiver)
	}

	return nil
}

// SendP2GMessage 发送群聊
func (l *IClientLogic) SendP2GMessage(ctx context.Context, receivers []string, message []byte) {
	for _, receiver := range receivers {
		_ = l.SendP2PMessage(ctx, receiver, message)
	}
	return
}

// SendBroadcastMessage 发送广播消息
func (l *IClientLogic) SendBroadcastMessage(ctx context.Context, message []byte) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, client := range l.clients {
		// 将消息写入客户端的 Send 通道
		select {
		case client.Send <- message:
			// 消息成功写入通道
		default:
			// 通道已满，丢弃消息或记录日志
			g.Log(consts.SocketLogger).Warningf(ctx, "Send channel is full, message dropped for client: %s", client.ClientID)
		}
	}

	return
}

// GetAllClients 返回所有在线客户端的 map，便于管理监控
func (l *IClientLogic) GetAllClients() map[string]*Client {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.clients
}

// GetAllClientIDs 返回所有在线客户端的 ID 列表
func (l *IClientLogic) GetAllClientIDs() []string {
	l.mu.RLock()
	defer l.mu.RUnlock()

	clients := make([]string, 0, len(l.clients))
	for id := range l.clients {
		clients = append(clients, id)
	}
	return clients
}
