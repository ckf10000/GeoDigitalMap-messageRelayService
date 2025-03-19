// Package federate
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate_new.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 00:25:38
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package federate

import (
	"GeoDigitalMap-messageRelayService/internal/logic/federate"
	"GeoDigitalMap-messageRelayService/internal/logic/manager"
)

type ControllerV1 struct {
	federateLogic *federate.IFederateLogic
}

func NewV1() *ControllerV1 {
	federateLogic := federate.NewIFederateLogic()
	manager.SetFederatePeerLogic(federateLogic)
	return &ControllerV1{
		federateLogic: federateLogic,
	}
}
