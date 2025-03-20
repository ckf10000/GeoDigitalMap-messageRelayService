// Package v1
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 23:34:41
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package v1

import (
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"github.com/gogf/gf/v2/frame/g"
)

// AddPeerReq 添加级联对端设备的请求
type AddPeerReq struct {
	g.Meta `path:"add-peer" method:"post" tags:"系统管理" summary:"添加级联对端设备"`
	IP     string `json:"ip" v:"required|ip" dc:"ip地址"`
	Port   uint32 `json:"port" v:"required|min:1025|max:65535" dc:"端口号"`
}

// AddPeerRes 添加级联对端设备的响应
type AddPeerRes struct{}

// GetAllFederatePeersReq 获取所有已级联的对端列表的请求
type GetAllFederatePeersReq struct {
	g.Meta `path:"getFederatePeers" method:"get" tags:"系统管理" summary:"获取所有已级联的对端列表"`
}

// GetAllFederatePeersRes 获取所有已级联的对端列表的响应
type GetAllFederatePeersRes struct {
	List []*dto.FederatePeerOutput `json:"list"` // 所有已级联的对端列表
}
