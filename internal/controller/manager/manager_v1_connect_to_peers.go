// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     manager_v1_connect_to_peers.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 23:07:18
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"GeoDigitalMap-messageRelayService/internal/common/utils"
	"context"
)

func (c *ControllerManagerFederateV1) ConnectToPeers(ctx context.Context) {
	hostAddrsDTO := utils.GetFederateRemoteAddress(ctx)
	c.federateLogic.ConnectToPeers(ctx, hostAddrsDTO)
}
