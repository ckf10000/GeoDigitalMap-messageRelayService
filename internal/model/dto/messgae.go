// Package dto
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     messgae.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-16 21:51:02
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package dto

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"time"
)

type MessageIutputDTO struct {
	MessageID   string             `json:"messageID"`   // 消息的messageID
	UID         string             `json:"uid"`         // 消息在服务端分配的uid
	Sender      string             `json:"sender"`      // 消息的发送者
	Receivers   []string           `json:"receivers"`   // 消息的接收者，使用PostgreSQL数组类型
	Content     interface{}        `json:"content"`     // 消息的内容
	CreatedAt   time.Time          `json:"createdAt"`   // 消息的创建时间，time.Time 默认序列化支持的是RFC3339格式：yyyy-MM-ddTHH:mm:ssZ
	MessageType consts.MessageType `json:"messageType"` // 消息类型
}

type RelayMessageIutputDTO struct {
	MessageIutput  *MessageIutputDTO
	FederateSource []string `json:"federateSource"` // 标记消息来源
}

type MessageOutputDTO struct {
	MessageID   string             `json:"messageID"`   // 消息的messageID
	Sender      string             `json:"sender"`      // 消息的发送者
	Receivers   []string           `json:"receivers"`   // 消息的接收者，使用PostgreSQL数组类型
	Content     interface{}        `json:"content"`     // 消息的内容
	CreatedAt   string             `json:"createdAt"`   // 消息的创建时间, messageOutputDTO.CreatedAt.Format(time.RFC3339)
	MessageType consts.MessageType `json:"messageType"` // 消息类型
}

type BroadcastMessageOutputDTO struct {
	MessageOutput  *MessageOutputDTO
	FederateSource []string `json:"federateSource"` // 标记消息来源
}
