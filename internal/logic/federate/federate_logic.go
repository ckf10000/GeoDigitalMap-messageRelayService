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

func (l *IFederateLogic) checkAddPeer(ctx context.Context, addr string) error {
	// 检查是否超过最大数量
	if len(l.peers) >= consts.FederatePeerCount {
		str := fmt.Sprintf("Exceeded the maximum cascading limit of %d peer devices", consts.FederatePeerCount)
		g.Log(consts.FederateLogger).Error(ctx, str)
		return gerror.New(str)
	}

	// 先加读锁检查是否已存在
	l.mu.RLock()
	_, exists := l.peers[addr]
	l.mu.RUnlock()

	if exists {
		str := fmt.Sprintf("Peer <%s> already exists", addr)
		g.Log(consts.FederateLogger).Error(ctx, str)
		return gerror.New(str)
	}

	// 再加写锁修改数据
	l.mu.Lock()
	defer l.mu.Unlock()

	// 再次检查，防止并发情况下同一地址重复添加
	if _, exists = l.peers[addr]; exists {
		str := fmt.Sprintf("Peer <%s> already exists", addr)
		g.Log(consts.FederateLogger).Error(ctx, str)
		return gerror.New(str)
	}
	return nil
}

func (l *IFederateLogic) pushPeer(ctx context.Context, addr string, conn *websocket.Conn) {
	// 创建带缓冲的发送通道
	peer := &Peer{
		Addr:   addr,
		Conn:   conn,
		Send:   make(chan []byte, consts.MessageChannelSize), // 缓冲区大小可根据实际需求调整
		JoinAt: time.Now(),                                   // 记录当前时间
	}

	l.peers[addr] = peer

	// 启动一个 goroutine 处理消息发送
	go peer.WritePump(ctx)

	g.Log(consts.FederateLogger).Infof(ctx, "New peer connected: %s", addr)
}

// AddPeer 添加新的级联对端
func (l *IFederateLogic) AddPeer(ctx context.Context, addr string, conn *websocket.Conn) error {

	err := l.checkAddPeer(ctx, addr)
	if err != nil {
		return err
	}

	l.pushPeer(ctx, addr, conn)
	return nil
}

// HandleMessages 循环读取来自级联对端的消息，并进行广播处理
func (l *IFederateLogic) HandleMessages(ctx context.Context, conn *websocket.Conn) {
	asyncCtx := context.WithoutCancel(ctx)
	addr := conn.RemoteAddr().String()

	// 客户端断开时移除
	defer func() {
		l.RemovePeer(asyncCtx, addr)
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			// 检查是否是Peer主动关闭连接
			var closeErr *websocket.CloseError
			if errors.As(err, &closeErr) {
				switch closeErr.Code {
				case websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived:
					g.Log(consts.FederateLogger).Warningf(asyncCtx, "Peer closed the connection")
				default:
					g.Log(consts.FederateLogger).Errorf(asyncCtx, "client closed the connection with unexpected code: %+v", closeErr.Code)
				}
			} else {
				g.Log(consts.FederateLogger).Error(asyncCtx, err)
			}
			break
		}

		// 此时读出来的数据，还是[]byte类型数据
		//g.Log(consts.FederateLogger).Infof(asyncCtx, "received message: %+v", message)

		// 处理消息逻辑
		var msg *dto.RelayMessageIutputDTO
		if err = json.Unmarshal(message, &msg); err != nil {
			g.Log(consts.FederateLogger).Errorf(asyncCtx, "Failed to parse federate relay message: %+v", err)
			continue
		}

		// 服务端生成 UUID
		msg.MessageIutput.MessageID = uuid.New().String()

		infoData, _ := gjson.Marshal(msg)
		//infoData, _ := gjson.MarshalIndent(msg, "", "\t")
		g.Log(consts.SocketLogger).Infof(asyncCtx, "received message: %s", string(infoData))

		l.RouteMessage(asyncCtx, msg)
	}
}

// RouteMessage 将消息发送给自己的所有客户端，以及广播给federate peer
func (l *IFederateLogic) RouteMessage(ctx context.Context, message *dto.RelayMessageIutputDTO) {

	messageOutputDTO := &dto.MessageOutputDTO{
		MessageID:   message.MessageIutput.MessageID,
		Sender:      message.MessageIutput.Sender,
		Receivers:   message.MessageIutput.Receivers,
		Content:     message.MessageIutput.Content,
		CreatedAt:   message.MessageIutput.CreatedAt.Format(time.RFC3339),
		MessageType: message.MessageIutput.MessageType,
	}
	switch message.MessageIutput.MessageType {
	case consts.Broadcast:
		// 将Federate 连接接收到的消息发送给所有客户端
		if len(manager.GetAllClientIDs()) > 0 {
			data, err := json.Marshal(messageOutputDTO)
			if err != nil {
				g.Log(consts.FederateLogger).Error(ctx, err)
				return
			}
			manager.GetClientLogic().SendBroadcastMessage(ctx, data)
		}
	}
	// TODO 目前中继消息仅支持广播消息的处理

	// 消息继续向其他的federate peer传递
	localIP := utils.GetFederateLocalIP(ctx)
	federateSource := append(message.FederateSource, localIP)
	broadcastMessageOutputDTO := &dto.BroadcastMessageOutputDTO{
		MessageOutput:  messageOutputDTO,
		FederateSource: federateSource,
	}

	data, err := json.Marshal(broadcastMessageOutputDTO)
	if err != nil {
		g.Log(consts.FederateLogger).Error(ctx, err)
		return
	}

	l.SendRelayMessage(ctx, data, message.FederateSource)
}

// SendRelayMessage 向所有级联对端发送消息，采用非阻塞写入
func (l *IFederateLogic) SendRelayMessage(ctx context.Context, message []byte, federateSource []string) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for addr, peer := range l.peers {
		// 不转发给来源 peer
		if utils.Contains(federateSource, addr) {
			continue
		}
		select {
		case peer.Send <- message:
			// 消息成功写入通道
		default:
			// 通道已满，丢弃消息或记录日志
			g.Log(consts.FederateLogger).Warningf(ctx, "Send channel is full, message dropped for peer: %s", peer.Addr)
		}
	}
}

// RemovePeer 根据 websocket 连接移除对端
func (l *IFederateLogic) RemovePeer(ctx context.Context, addr string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if peer, exists := l.peers[addr]; exists {
		close(peer.Send) // 关闭通道
		delete(l.peers, addr)
		g.Log(consts.FederateLogger).Warningf(ctx, "Peer disconnected: %s", addr)
	}
}

// GetAllPeerAddrs 返回所有在线Peer的 IP地址 列表
func (l *IFederateLogic) GetAllPeerAddrs() []string {
	l.mu.RLock()
	defer l.mu.RUnlock()

	peers := make([]string, 0, len(l.peers))
	for addr := range l.peers {
		peers = append(peers, addr)
	}
	return peers
}

func (l *IFederateLogic) connectToPeer(ctx context.Context, hostAddrDTO *dto.HostAddress) (*websocket.Conn, error) {
	// 实现连接逻辑，例如 WebSocket 连接
	url := fmt.Sprintf("ws://%s:%d%s", hostAddrDTO.IP, hostAddrDTO.Port, consts.FEDERATEMODULE)
	g.Log(consts.FederateLogger).Infof(ctx, "Start connecting: %s", url)
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, url, nil)
	if err != nil {
		g.Log(consts.FederateLogger).Error(ctx, err)
		return nil, err
	}
	return conn, nil
}

// RestAddPeer 从前端页面手动添加federate peer
func (l *IFederateLogic) RestAddPeer(ctx context.Context, hostAddrDTO *dto.HostAddress) error {
	err := l.checkAddPeer(ctx, hostAddrDTO.IP)
	if err != nil {
		return err
	}

	conn, err := l.connectToPeer(ctx, hostAddrDTO)
	if err != nil {
		return err
	}
	l.pushPeer(ctx, hostAddrDTO.IP, conn)
	return nil
}

// ConnectToPeers 配置文件中存在 federateRemoteAddress 配置时，需要添加federate peer
func (l *IFederateLogic) ConnectToPeers(ctx context.Context, hostAddrsDTO []*dto.HostAddress) {
	for _, addrDTO := range hostAddrsDTO {
		go func(hostAddrDTO *dto.HostAddress) {
			asyncCtx := context.WithoutCancel(ctx)
			count := 0
			for {
				err := l.RestAddPeer(asyncCtx, hostAddrDTO)
				if err != nil {
					count++
					if count >= consts.FederateRetryConnectCount {
						break
					}
					time.Sleep(consts.FederateRetryConnectinterval) // 重试间隔
					continue
				}
				break
			}
		}(addrDTO)
	}
}
