// Package client
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     manager_logic.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 15:00:37
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package client

// GlobalClientLogic 全局单例，便于其它模块调用管理在线客户端
var GlobalClientLogic = NewIClientLogic()

// GetClientLogic 返回全局客户端逻辑管理实例
func GetClientLogic() *IClientLogic {
	return GlobalClientLogic
}

// GetAllClients 是辅助接口，用于管理接口返回在线客户端列表
func GetAllClients() []string {
	return GlobalClientLogic.GetAllClients()
}
