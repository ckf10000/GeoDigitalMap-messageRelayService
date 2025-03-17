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
	federateLogic "GeoDigitalMap-messageRelayService/internal/logic/federate"
)

type ControllerV1 struct {
	federate *federateLogic.IFederateLogic
}

func NewV1() *ControllerV1 {
	return &ControllerV1{
		federate: federateLogic.NewFederateLogic(),
	}
}
