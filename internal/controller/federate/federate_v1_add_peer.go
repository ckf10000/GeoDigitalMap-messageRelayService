// Package federate
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate_v1_add_peer.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 00:40:53
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package federate

import (
	"context"
	"github.com/gorilla/websocket"
)

func (c *ControllerV1) AddPeer(ctx context.Context, remoteAddr string, ws *websocket.Conn) error {
	return c.federate.AddPeer(ctx, remoteAddr, ws)
}
