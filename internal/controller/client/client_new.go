// Package client
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client_new.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 11:03:08
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package client

import (
	"GeoDigitalMap-messageRelayService/internal/logic/client"
)

type ControllerV1 struct {
	clientLogic *client.IClientLogic
}

func NewV1() *ControllerV1 {
	return &ControllerV1{
		clientLogic: client.GetClientLogic(),
	}
}
