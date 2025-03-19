// Package dto
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     network.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 21:36:40
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package dto

type HostAddress struct {
	IP   string `json:"ip" v:"ip"`          // ip地址
	Port uint32 `json:"port" v:"max:65535"` // 端口号
}
