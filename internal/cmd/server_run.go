// Package cmd
package cmd

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/controller/maintenance"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			ser := g.Server()
			// 主路由定义
			ser.Group("/api", func(group *ghttp.RouterGroup) {
				// 添加了一个全局中间件，处理响应数据
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// 注册 V1 子路由
				registerV1Routes(group)
				// TODO 未来可以添加 V2... 等这样的设计模块子路由
			})
			//ser.SetIndexFolder(true)
			//ser.SetServerRoot("/resource/public/resource")
			ser.AddStaticPath("/swagger/js", "/resource/public/resource/js")
			ser.AddStaticPath("/swagger/css", "/resource/public/resource/css")
			ser.SetSwaggerUITemplate(consts.SwaggerUITemplate)
			ser.Run()
			return nil
		},
	}
)

// registerV1Routes 将 v1 路由注册逻辑单独封装
func registerV1Routes(group *ghttp.RouterGroup) {
	group.Group("/v1", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.Auth)
		// TODO 添加自定义的中间件钩子

		// 注册子模块的路由
		registerMaintenanceRoutes(group)
	})
}

// registerMaintenanceRoutes 系统运维模块路由的封装
func registerMaintenanceRoutes(group *ghttp.RouterGroup) {
	group.Group("/maintenance", func(group *ghttp.RouterGroup) {
		group.Bind(maintenance.NewV1())
	})
}
