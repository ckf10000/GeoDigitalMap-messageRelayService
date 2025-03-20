// Package cmd
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     federate_server.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-16 22:51:26
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package cmd

import (
	"GeoDigitalMap-messageRelayService/internal/common/utils"
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/controller/federate"
	"GeoDigitalMap-messageRelayService/internal/middleware/auth"
	"GeoDigitalMap-messageRelayService/internal/middleware/forward"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
)

func CreateFederationServer(ctx context.Context) *ghttp.Server {
	ser := g.Server(consts.FederateService)
	ser.SetLogger(g.Log(consts.FederateLogger))

	if !utils.CheckFederateLocalIP(ctx) {
		str := "The configured federated address does not match the IP address of the deployment server"
		g.Log(consts.FederateLogger).Error(ctx, str)
		return nil
	}

	// 主路由定义
	ser.Group(consts.ROUTEROOT, func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 添加认证中间件
		group.Middleware(auth.PeerAuthMiddleware)
		// 添加 WebSocket 连接升级中间件
		group.Middleware(forward.WebSocketUpgradeMiddleware)

		// WebSocket 连接处理
		group.GET(consts.FEDERATEMODULE, func(r *ghttp.Request) {
			subCtx := r.Context()
			federateCtl := federate.NewV1()

			// 从上下文获取 WebSocket 连接
			wsValue := r.GetCtxVar(consts.ContextWSConnKey).Interface() // 先转换为 interface{}
			ws, ok := wsValue.(*websocket.Conn)                         // 再进行类型断言
			if !ok || ws == nil {
				str := fmt.Sprint("Failed to get WebSocket connection from context")
				g.Log(consts.SocketLogger).Error(subCtx, str)
				r.Response.WriteExit(str)
				return
			}

			err := federateCtl.AddPeer(subCtx, r.RemoteAddr, ws)
			if err != nil {
				return
			}
			// 处理级联连接的消息和心跳处理
			go federateCtl.HandleMessages(subCtx, ws)
		})
	})

	//ser.SetGraceful(true)
	//ser.EnableAdmin()
	// Configure static file serving
	//ser.SetServerRoot("static")
	// Set server port
	//ser.SetPort(28090)
	return ser
}
