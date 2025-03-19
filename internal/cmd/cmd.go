// Package cmd
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     cmd.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-14 23:24:22
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package cmd

import (
	"GeoDigitalMap-messageRelayService/internal/common/utils"
	"GeoDigitalMap-messageRelayService/internal/consts"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start message relay server",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {

			// 创建服务器实例池
			servers := make([]*ghttp.Server, 0, consts.ServiceInstancePoolCapacity)
			if isServiceEnable := utils.GetServiceEnableStatus(ctx, consts.RestAPIService); isServiceEnable {
				ser := CreateRestAPIServer(ctx)
				if ser == nil {
					return nil
				}
				servers = append(servers, ser)
			}
			if isServiceEnable := utils.GetServiceEnableStatus(ctx, consts.TCPSocketService); isServiceEnable {
				ser := CreateTCPSocketServer()
				if ser == nil {
					return nil
				}
				servers = append(servers, ser)
			}
			if isServiceEnable := utils.GetServiceEnableStatus(ctx, consts.SSLSocketService); isServiceEnable {
				ser := CreateSSLSocketServer(ctx)
				if ser == nil {
					return nil
				}
				servers = append(servers, ser)
			}
			if isServiceEnable := utils.GetServiceEnableStatus(ctx, consts.FederateService); isServiceEnable {
				ser := CreateFederationServer(ctx)
				if ser == nil {
					return nil
				}
				servers = append(servers, ser)
			}

			// 使用 goroutine 启动所有服务器
			for _, s := range servers {
				server := s // 避免闭包捕获问题
				go func() {
					server.Run()
				}()

			}

			// 主进程阻塞等待
			g.Wait() // 或使用 select{} 保持主进程存活
			return nil
		},
	}
)
