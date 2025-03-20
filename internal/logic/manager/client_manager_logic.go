// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client_manager_logic.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 15:00:37
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"context"
)

// SetClientLogic 用于设置 Client 逻辑
func SetClientLogic(logic IClientLogic) {
	globalClientLogic = logic
}

// GetClientLogic 返回全局客户端逻辑管理实例
func GetClientLogic() IClientLogic {
	return globalClientLogic
}

// GetAllClientIDs 是辅助接口，用于管理接口返回在线客户端ID列表
func GetAllClientIDs() []*string {
	return globalClientLogic.GetAllClientIDs()
}

// GetAllClients 是辅助接口，用于管理接口返回在线客户端列表
func GetAllClients(ctx context.Context) []*dto.OnlineClientOutput {
	return globalClientLogic.GetAllClients(ctx)
}
