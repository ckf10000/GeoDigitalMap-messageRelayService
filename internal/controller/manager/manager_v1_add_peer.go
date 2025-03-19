// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     manager_v1_add_peer.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-20 00:09:24
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"GeoDigitalMap-messageRelayService/api/manager/v1"
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"context"
)

func (c *ControllerV1) AddPeer(ctx context.Context, req *v1.AddPeerReq) (res *v1.AddPeerRes, err error) {
	hostAddrsDTO := &dto.HostAddress{
		IP:   req.IP,
		Port: req.Port,
	}
	return nil, c.federateLogic.RestAddPeer(ctx, hostAddrsDTO)
}
