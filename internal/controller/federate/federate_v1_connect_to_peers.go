// Package federate
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate_v1_connect_to_peers.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 20:48:25
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package federate

import (
	"GeoDigitalMap-messageRelayService/internal/common/utils"
	"context"
)

func (c *ControllerV1) ConnectToPeers(ctx context.Context) {
	hostAddrsDTO := utils.GetFederateRemoteAddress(ctx)
	c.federateLogic.ConnectToPeers(ctx, hostAddrsDTO)
}
