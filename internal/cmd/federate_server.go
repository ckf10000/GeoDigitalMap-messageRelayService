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
	"GeoDigitalMap-messageRelayService/internal/consts"
	federateCTL "GeoDigitalMap-messageRelayService/internal/controller/federate"
	"GeoDigitalMap-messageRelayService/internal/middleware/forward"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func CreateFederationServer(ctx context.Context) *ghttp.Server {
	s := g.Server(consts.FederateService)
	s.SetLogger(g.Log(consts.FederateLogger))

	// Bind WebSocket handler to / endpoint
	s.BindHandler(consts.FEDERATEROOT, func(r *ghttp.Request) {
		federate := federateCTL.NewV1()
		ws, err := forward.WSUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err != nil {
			g.Log(consts.FederateService).Errorf(ctx, "WS upgrade failed: %+v", err)
			return
		}
		err = federate.AddPeer(ctx, r.RemoteAddr, ws)
		if err != nil {
			g.Log(consts.FederateService).Errorf(ctx, "Adding peer failed, %+v", err)
			return
		}
		// 启动级联连接的消息和心跳处理
		go federate.HandleFederateConnection(ctx, ws)
	})
	s.SetGraceful(true)
	s.EnableAdmin()
	// Configure static file serving
	//ser.SetServerRoot("static")
	// Set server port
	//ser.SetPort(28090)
	return s
}
