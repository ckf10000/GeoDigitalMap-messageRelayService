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

type MessageOutputDTO struct {
	ID          string               `json:"id" orm:"id,primary"`
	Sender      string               `json:"sender" orm:"sender"`
	Receivers   []string             `json:"receivers" orm:"receivers"` // 使用PostgreSQL数组类型
	Content     []byte               `json:"content" orm:"content"`
	Status      consts.MessageStatus `json:"status" orm:"status"`
	CreatedAt   time.Time            `json:"created_at" orm:"created_at"`
	RetryCount  int                  `json:"retry_count" orm:"retry_count"`
	MessageType consts.MessageType   `json:"message_type" orm:"message_type"`
}
