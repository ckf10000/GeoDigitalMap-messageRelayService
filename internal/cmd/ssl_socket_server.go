// Package cmd
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     ssl_socket_server.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-14 22:23:03
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package cmd

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/controller/client"
	"GeoDigitalMap-messageRelayService/internal/middleware/auth"
	"GeoDigitalMap-messageRelayService/internal/middleware/cert"
	"GeoDigitalMap-messageRelayService/internal/middleware/forward"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
)

func CreateSSLSocketServer(ctx context.Context) *ghttp.Server {
	crtPath, keyPath := cert.GenerateCert(ctx, consts.SSLSocketService)
	ser := g.Server(consts.SSLSocketService)
	ser.SetLogger(g.Log(consts.SocketLogger))

	// 主路由定义
	ser.Group(consts.ROUTEROOT, func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 添加认证中间件
		group.Middleware(auth.ClientAuthMiddleware)
		// 添加 WebSocket 连接升级中间件
		group.Middleware(forward.WebSocketUpgradeMiddleware)

		// WebSocket 连接处理
		group.GET(consts.ROUTEROOT, func(r *ghttp.Request) {
			subCtx := r.Context()
			clientCtl := client.NewV1()

			// 从上下文获取 WebSocket 连接
			wsValue := r.GetCtxVar("ws_conn").Interface() // 先转换为 interface{}
			ws, ok := wsValue.(*websocket.Conn)           // 再进行类型断言
			if !ok || ws == nil {
				str := fmt.Sprint("Failed to get WebSocket connection from context")
				g.Log(consts.SocketLogger).Error(subCtx, str)
				r.Response.WriteExit(str)
				return
			}

			// 注册新的客户端连接
			err := clientCtl.AddClient(subCtx, r.RemoteAddr, ws)
			if err != nil {
				str := fmt.Sprintf("Remote host connection failed: %+v", err)
				g.Log(consts.SocketLogger).Error(subCtx, str)
				r.Response.WriteExit(str)
				return
			}

			// 异步启动消息处理逻辑
			go clientCtl.HandleMessages(subCtx, ws)
		})
	})

	ser.EnableHTTPS(crtPath, keyPath)
	//ser.SetGraceful(true)
	//ser.EnableAdmin()
	// Configure static file serving
	//ser.SetServerRoot("static")
	// Set server port
	//ser.SetPort(28443)
	return ser
}
