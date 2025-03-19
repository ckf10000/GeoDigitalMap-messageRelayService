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
	"GeoDigitalMap-messageRelayService/internal/controller/client"
	"GeoDigitalMap-messageRelayService/internal/middleware/forward"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func CreateTCPSocketServer() *ghttp.Server {
	ser := g.Server(consts.TCPSocketService)
	ser.SetLogger(g.Log(consts.SocketLogger))

	// Bind WebSocket handler to / endpoint
	ser.BindHandler(consts.WSSROOT, func(r *ghttp.Request) {
		subCtx := r.Context()
		clientCtl := client.NewV1()
		ws, err := forward.WSUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err != nil {
			g.Log(consts.SocketLogger).Errorf(subCtx, "WS upgrade failed: %+v", err)
			r.Response.Write(err.Error())
			return
		}

		// 注册新的客户端连接，使用 RemoteAddr 作为示例标识（实际项目中建议使用唯一ID）
		err = clientCtl.AddClient(subCtx, r.RemoteAddr, ws)
		if err != nil {
			g.Log(consts.SocketLogger).Errorf(subCtx, "Remote host connection failed: %+v", err)
			r.Response.Write(err.Error())
			return
		}

		// 异步启动消息处理逻辑
		go clientCtl.HandleMessages(subCtx, ws)
	})
	//ser.SetGraceful(true)
	//ser.EnableAdmin()
	// Configure static file serving
	//ser.SetServerRoot("static")
	// Set server port
	//ser.SetPort(28080)
	return ser
}
