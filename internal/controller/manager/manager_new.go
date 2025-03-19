// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     manager_new.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 22:35:58
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"GeoDigitalMap-messageRelayService/internal/logic/client"
	"GeoDigitalMap-messageRelayService/internal/logic/federate"
	"GeoDigitalMap-messageRelayService/internal/logic/manager"
)

type ControllerManagerClientV1 struct {
	clientLogic *client.IClientLogic
}

type ControllerManagerFederateV1 struct {
	federateLogic *federate.IFederateLogic
}

func NewControllerManagerClientV1() *ControllerManagerClientV1 {
	clientLogic := client.NewIClientLogic()
	manager.SetClientLogic(clientLogic)
	return &ControllerManagerClientV1{
		clientLogic: clientLogic,
	}
}

func NewControllerManagerFederateV1() *ControllerManagerFederateV1 {
	federateLogic := federate.NewIFederateLogic()
	manager.SetFederatePeerLogic(federateLogic)
	return &ControllerManagerFederateV1{
		federateLogic: federateLogic,
	}
}
