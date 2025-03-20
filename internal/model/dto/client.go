// Package dto
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-20 17:17:18
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package dto

type OnlineClientOutput struct {
	UserID   string `json:"userId"`   // 客户端用户唯一标识
	ClientID string `json:"clientID"` // 客户端唯一标识
	Username string `json:"username"` // 用户名称
	JoinAt   string `json:"joinAt"`   // 客户端接入时间`
}
