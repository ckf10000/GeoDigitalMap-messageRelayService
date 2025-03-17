// Package client
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client_v1_handle_messages.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 11:09:43
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package client

import (
	"context"
	"github.com/gorilla/websocket"
)

func (c *ControllerV1) HandleMessages(ctx context.Context, ws *websocket.Conn) {
	c.clientLogic.HandleMessages(ctx, ws)
}
