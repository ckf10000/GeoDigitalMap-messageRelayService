// Package client
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     client_v1_add_client.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-17 11:07:49
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package client

import (
	"context"
	"github.com/gorilla/websocket"
)

// AddClient 注册新的客户端连接
func (c *ControllerV1) AddClient(ctx context.Context, clientID string, ws *websocket.Conn) error {
	return c.clientLogic.AddClient(ctx, clientID, ws)
}
