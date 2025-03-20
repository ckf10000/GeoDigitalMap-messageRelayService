// Package federate
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     peer_logic.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-18 17:45:26
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package federate

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"time"
)

// WritePump 将待发送的消息写入到 websocket 连接，带有心跳检测
func (p *Peer) WritePump(ctx context.Context) {
	asyncCtx := context.WithoutCancel(ctx)
	ticker := time.NewTicker(consts.ClientHeartbeatDuration)
	defer func() {
		ticker.Stop()
		// 关闭连接
		_ = p.Conn.Close()
		return
	}()

	for {
		select {
		case message, ok := <-p.Send:
			if !ok {
				// 通道关闭，退出循环
				return
			}

			// 发送消息给客户端
			if err := p.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				g.Log(consts.FederateLogger).Errorf(asyncCtx, "Write message error: %+v", err)
				return
			}
		}
	}
}
