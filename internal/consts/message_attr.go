// Package consts
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     message_attr.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 11:15:48
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package consts

// MessageType 定义消息类型
type MessageType string

const (
	PointToPoint MessageType = "P2P"
	PointToGroup MessageType = "P2G"
	Broadcast    MessageType = "BROADCAST"
)
