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
	"github.com/gogf/gf/v2/frame/g"
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
