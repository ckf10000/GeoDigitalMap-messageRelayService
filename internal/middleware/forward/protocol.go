// Package forward
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     protocol.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-15 22:48:19
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package forward

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"net/http"
)

// WSUpGrader HTTP升级到WebSocket
var WSUpGrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		g.Log(consts.SocketLogger).Infof(r.Context(), "HTTP upgrade protocol, 'ua': %s, 'referer': %s", r.Header["User-Agent"], r.Header["Referer"])
		return true // 允许所有来源连接(允许跨域)
	},
	// Error handler for upgrade failures
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		// Implement error handling logic here
	},
}

func WebSocketUpgradeMiddleware(r *ghttp.Request) {
	ctx := r.Context()

	// 执行 WS 升级
	ws, err := WSUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		str := fmt.Sprintf("WS upgrade failed: %+v", err)
		g.Log(consts.SocketLogger).Error(ctx, str)
		//r.Response.WriteHeader(500)
		// 终止后续执行
		r.Response.WriteExit(err.Error())
		//r.Exit()
		return
	}

	// 将 WebSocket 连接存入上下文，供后续处理
	r.SetCtxVar("ws_conn", ws)

	// 继续执行后续处理
	r.Middleware.Next()
}
