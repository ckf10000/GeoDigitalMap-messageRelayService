// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client_manager_logic.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-18 22:31:21
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import "GeoDigitalMap-messageRelayService/internal/model/dto"

// SetFederatePeerLogic 用于设置 Federate 逻辑
func SetFederatePeerLogic(logic IFederateLogic) {
	globalFederatePeerLogic = logic
}

// GetFederatePeerLogic 返回全局Federate在线Peer逻辑管理实例
func GetFederatePeerLogic() IFederateLogic {
	return globalFederatePeerLogic
}

// GetAllPeerAddrs 是辅助接口，用于返回所有在线Peer的 IP地址列表
func GetAllPeerAddrs() []*string {
	return globalFederatePeerLogic.GetAllPeerAddrs()
}

// GetAllFederatePeers 是辅助接口，用于返回所有Federate在线Peer列表
func GetAllFederatePeers() []*dto.FederatePeerOutput {
	return globalFederatePeerLogic.GetAllFederatePeers()
}
