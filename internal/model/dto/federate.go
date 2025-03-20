// Package dto
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-20 22:43:28
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package dto

type FederatePeerOutput struct {
	PeerAddr string `json:"peerAddress"` // 用IP地址作为对端的唯一标识
	JoinAt   string `json:"joinAt"`      // peer接入时间`
}
