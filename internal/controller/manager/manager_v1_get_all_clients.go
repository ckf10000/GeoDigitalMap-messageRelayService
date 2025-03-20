// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     manager_v1_get_all_clients.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-20 21:05:33
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"GeoDigitalMap-messageRelayService/api/manager/v1"
	"GeoDigitalMap-messageRelayService/internal/logic/manager"
	"context"
)

func (c *ControllerV1) GetOnlineClients(ctx context.Context, req *v1.GetOnlineClientsReq) (res *v1.GetOnlineClientsRes, err error) {
	// **确保 res 被初始化**
	res = &v1.GetOnlineClientsRes{}
	res.List = manager.GetAllClients(ctx)
	return res, nil
}
