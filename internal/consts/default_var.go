// Package consts
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     default_var.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 09:41:54
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package consts

import "time"

const (
	ShardCount              = 32               // 根据CPU核心数动态调整
	FederatePeerCount       = 1024             // 最多可级联的对端设备数量
	SocketClientCount       = 100000           // 最多可接入客户端的数量
	MessageChannelSize      = 1024             // 消息通道缓冲区大小
	ClientHeartbeatDuration = 60 * time.Second // server 检测 socket client 的心跳周期
)
