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
	"GeoDigitalMap-messageRelayService/api/manager"
	"GeoDigitalMap-messageRelayService/internal/logic/client"
	"GeoDigitalMap-messageRelayService/internal/logic/federate"
	manageLogic "GeoDigitalMap-messageRelayService/internal/logic/manager"
)

type ControllerV1 struct {
	clientLogic   manageLogic.IClientLogic
	federateLogic manageLogic.IFederateLogic
}

func NewV1() manager.IManagerV1 {
	return &ControllerV1{
		clientLogic:   manageLogic.GetClientLogic(),
		federateLogic: manageLogic.GetFederatePeerLogic(),
	}
}

type ControllerManagerClientV1 struct {
	clientLogic *client.IClientLogic
}

type ControllerManagerFederateV1 struct {
	federateLogic *federate.IFederateLogic
}

func NewControllerManagerClientV1() *ControllerManagerClientV1 {
	clientLogic := client.NewIClientLogic()
	manageLogic.SetClientLogic(clientLogic)
	return &ControllerManagerClientV1{
		clientLogic: clientLogic,
	}
}

func NewControllerManagerFederateV1() *ControllerManagerFederateV1 {
	federateLogic := federate.NewIFederateLogic()
	manageLogic.SetFederatePeerLogic(federateLogic)
	return &ControllerManagerFederateV1{
		federateLogic: federateLogic,
	}
}
