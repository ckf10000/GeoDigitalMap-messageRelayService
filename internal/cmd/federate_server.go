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
	"GeoDigitalMap-messageRelayService/internal/middleware/forward"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func CreateFederationServer(ctx context.Context) *ghttp.Server {
	ser := g.Server(consts.FederateService)
	ser.SetLogger(g.Log(consts.FederateLogger))

	if !utils.CheckFederateLocalIP(ctx) {
		str := "The configured federated address does not match the IP address of the deployment server"
		g.Log(consts.FederateLogger).Error(ctx, str)
		return nil
	}

	// Bind WebSocket handler to / endpoint
	ser.BindHandler(consts.FEDERATEROOT, func(r *ghttp.Request) {
		subCtx := r.Context()
		federateCtl := federate.NewV1()
		ws, err := forward.WSUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
		if err != nil {
			g.Log(consts.FederateLogger).Errorf(subCtx, "WS upgrade failed: %+v", err)
			return
		}
		err = federateCtl.AddPeer(subCtx, r.RemoteAddr, ws)
		if err != nil {
			g.Log(consts.FederateLogger).Errorf(subCtx, "Adding peer failed, %+v", err)
			return
		}
		// 处理级联连接的消息和心跳处理
		go federateCtl.HandleMessages(subCtx, ws)

		// 连接所有的federate peer
		federateCtl.ConnectToPeers(subCtx)
	})
	//ser.SetGraceful(true)
	//ser.EnableAdmin()
	// Configure static file serving
	//ser.SetServerRoot("static")
	// Set server port
	//ser.SetPort(28090)
	return ser
}
