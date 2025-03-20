// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     manager_v1_get_federate_peers.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-20 22:48:23
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"GeoDigitalMap-messageRelayService/api/manager/v1"
	"GeoDigitalMap-messageRelayService/internal/logic/manager"
	"context"
)

func (c *ControllerV1) GetAllFederatePeers(ctx context.Context, req *v1.GetAllFederatePeersReq) (res *v1.GetAllFederatePeersRes, err error) {
	// **确保 res 被初始化**
	res = &v1.GetAllFederatePeersRes{}
	res.List = manager.GetAllFederatePeers(ctx)
	return res, nil
}
