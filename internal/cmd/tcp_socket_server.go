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
		// Upgrade HTTP connection to WebSocket
		ws, err := WSUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err != nil {
			r.Response.Write(err.Error())
			return
		}

		defer func(ws *websocket.Conn) {
			err = ws.Close()
			if err != nil {
				g.Log(consts.TCPSocketService).Error(ctx, err)
			}
		}(ws)

		// Message handling loop
		for {
			// Read incoming WebSocket message
			msgType, msg, err1 := ws.ReadMessage()
			if err1 != nil {
				break // Connection closed or error occurred
			}
			// Log received message
			g.Log(consts.TCPSocketService).Infof(ctx, "received message: %s", msg)
			// Echo the message back to client
			if err = ws.WriteMessage(msgType, msg); err != nil {
				break // Error writing message
			}
		}
		// Log connection closure
		g.Log(consts.TCPSocketService).Info(ctx, "websocket connection closed")
	})
	ser.SetGraceful(true)
	ser.EnableAdmin()
	// Configure static file serving
	//ser.SetServerRoot("static")
	// Set server port
	//ser.SetPort(28080)
	return ser
}
