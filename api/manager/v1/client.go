// Package v1
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-20 18:06:27
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package v1

import (
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"github.com/gogf/gf/v2/frame/g"
)

// GetOnlineClientsReq 获取所有的在线客户端的请求
type GetOnlineClientsReq struct {
	g.Meta `path:"getOnlineClients" method:"get" tags:"系统管理" summary:"获取所有的在线客户端"`
}

// GetOnlineClientsRes 获取所有的在线客户端的响应
type GetOnlineClientsRes struct {
	List []*dto.OnlineClientOutput `json:"list"` // 在线客户端列表
}
