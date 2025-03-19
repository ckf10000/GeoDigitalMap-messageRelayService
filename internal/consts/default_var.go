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
	FederatePeerCount            = 1024             // 最多可级联的对端设备数量
	SocketClientCount            = 100000           // 最多可接入客户端的数量
	MessageChannelSize           = 1024             // 消息通道缓冲区大小
	ClientHeartbeatDuration      = 60 * time.Second // server 检测 socket client 的心跳周期
	ServiceInstancePoolCapacity  = 16               // 创建服务实例池的最大容量
	FederateRetryConnectCount    = 3                // 连接federate peer 失败重试的次数
	FederateRetryConnectinterval = 10 * time.Second // 连接federate peer 失败重试的间隔
)
