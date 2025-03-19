// Package federate
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate_v1_handle_message.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 00:48:18
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package federate

import (
	"context"
	"github.com/gorilla/websocket"
)

func (c *ControllerV1) HandleMessages(ctx context.Context, ws *websocket.Conn) {
	c.federateLogic.HandleMessages(ctx, ws)
}
