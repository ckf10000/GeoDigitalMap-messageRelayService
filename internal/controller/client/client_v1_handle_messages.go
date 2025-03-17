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
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/model/dto"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
)

// HandleMessages 循环读取并解析客户端消息，交由底层逻辑进行路由
func (c *ControllerV1) HandleMessages(ctx context.Context, ws *websocket.Conn) {
	asyncCtx := context.WithoutCancel(ctx)
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			g.Log().Error(asyncCtx, err)
		}
	}(ws)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			g.Log().Error(ctx, err)
			c.clientLogic.RemoveClient(ws)
			break
		}
		g.Log(consts.SocketLogger).Infof(ctx, "received handle: %s", message)

		// 将收到的消息原样回传
		//if err = ws.WriteMessage(msgType, message); err != nil {
		//	break
		//}

		var msg dto.MessageOutputDTO
		if err = json.Unmarshal(message, &msg); err != nil {
			// 解析出错，记录日志或返回错误响应
			g.Log(consts.SocketLogger).Error(ctx, err)
			continue
		}
		c.clientLogic.RouteMessage(ctx, msg)
	}
}
