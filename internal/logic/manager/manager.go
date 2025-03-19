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
	"context"
)

// IClientLogic 接口定义
type IClientLogic interface {
	GetAllClientIDs() []string
	SendBroadcastMessage(ctx context.Context, message []byte)
}

// IFederateLogic 接口定义
type IFederateLogic interface {
	GetAllPeerAddrs() []string
	SendBroadcastMessage(ctx context.Context, message []byte, federateSource []string)
}

// globalClientLogic 全局单例，便于其它模块调用管理在线客户端
var globalClientLogic IClientLogic

// globalFederatePeerLogic 全局单例，便于其它模块调用管理Federate在线Peer
var globalFederatePeerLogic IFederateLogic
