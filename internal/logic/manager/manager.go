// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     manager.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 08:40:53
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"context"
	"github.com/gorilla/websocket"
)

// IClientLogic 接口定义
type IClientLogic interface {
	GetAllClientIDs() []*string
	GetAllClients() []*dto.OnlineClientOutput
	SendBroadcastMessage(ctx context.Context, message []byte)
	AddClient(ctx context.Context, id string, conn *websocket.Conn) error
	HandleMessages(ctx context.Context, conn *websocket.Conn)
}

// IFederateLogic 接口定义
type IFederateLogic interface {
	GetAllPeerAddrs() []string
	SendRelayMessage(ctx context.Context, message []byte, federateSource []string)
	ConnectToPeers(ctx context.Context, hostAddrsDTO []*dto.HostAddress)
	AddPeer(ctx context.Context, addr string, conn *websocket.Conn) error
	HandleMessages(ctx context.Context, conn *websocket.Conn)
	RestAddPeer(ctx context.Context, hostAddrDTO *dto.HostAddress) error
}

// globalClientLogic 全局单例，便于其它模块调用管理在线客户端
var globalClientLogic IClientLogic

// globalFederatePeerLogic 全局单例，便于其它模块调用管理Federate在线Peer
var globalFederatePeerLogic IFederateLogic
