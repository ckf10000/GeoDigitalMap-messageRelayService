// Package cmd
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     tcp_socket_server.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-14 22:22:15
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package cmd

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"net/http"
)

// WSUpGrader HTTP升级到WebSocket
var WSUpGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源连接(允许跨域)
	},
	// Error handler for upgrade failures
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		// Implement error handling logic here
	},
}

func TCPSocketServer(ctx context.Context) *ghttp.Server {
	ser := g.Server(consts.TCPSocketService)
	ser.SetLogger(g.Log(consts.TCPSocketService))

	// Bind WebSocket handler to / endpoint
	ser.BindHandler("/", func(r *ghttp.Request) {
		ws, err1 := WSUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err1 != nil {
			r.Response.Write(err1.Error())
			return
		}

		defer func(ws *websocket.Conn) {
			err2 := ws.Close()
			if err2 != nil {
				g.Log(consts.TCPSocketService).Error(ctx, err2)
			}
		}(ws)

		for {
			msgType, msg, err3 := ws.ReadMessage()
			if err3 != nil {
				break
			}
			g.Log(consts.TCPSocketService).Infof(ctx, "received message: %s", msg)
			if err3 = ws.WriteMessage(msgType, msg); err3 != nil {
				break
			}
		}
		g.Log(consts.TCPSocketService).Info(ctx, "websocket connect closed")
	})
	ser.SetGraceful(true)
	ser.EnableAdmin()
	// Configure static file serving
	//ser.SetServerRoot("static")
	// Set server port
	//ser.SetPort(28080)
	return ser
}
