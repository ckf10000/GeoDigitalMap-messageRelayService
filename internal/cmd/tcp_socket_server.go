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
	"GeoDigitalMap-messageRelayService/internal/middleware/auth"
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

	err := ser.SetConfig(ghttp.ServerConfig{
		MaxHeaderBytes: 1024, // 减少内存占用
		IdleTimeout:    300,  // 秒
		ReadTimeout:    10,
		WriteTimeout:   10,
		KeepAlive:      true, // 启用长连接
	})
	if err != nil {
		return nil
	}

	// 主路由分组
	ser.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(auth.ClientAuth) // 身份验证中间件
		registerTCPV1Routes(group)        // 注册子路由
	})

	// Bind WebSocket handler to / endpoint
	ser.BindHandler("/", func(r *ghttp.Request) {
		// Upgrade HTTP connect to WebSocket
		ws, err1 := WSUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err1 != nil {
			r.Response.Write(err1.Error())
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
			msgType, msg, err2 := ws.ReadMessage()
			if err2 != nil {
				break // Connection closed or error occurred
			}
			// Log received message
			g.Log(consts.TCPSocketService).Infof(ctx, "received message: %s", msg)
			// Echo the message back to client
			if err = ws.WriteMessage(msgType, msg); err != nil {
				break // Error writing message
			}
		}
		// Log connect closure
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

func registerTCPV1Routes(group *ghttp.RouterGroup) {
	group.Group("ws/v1", func(v1 *ghttp.RouterGroup) {
		v1.Bind()
	})
}
