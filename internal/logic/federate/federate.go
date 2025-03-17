// Package federate
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-16 21:47:08
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package federate

import (
	"sync"
)

// IFederateLogic 负责级联对端的管理与消息转发
type IFederateLogic struct {
	peers sync.Map // map[string]*FederatedPeer
}

// NewFederateLogic 构造新的级联逻辑实例
func NewFederateLogic() *IFederateLogic {
	return &IFederateLogic{}
}
