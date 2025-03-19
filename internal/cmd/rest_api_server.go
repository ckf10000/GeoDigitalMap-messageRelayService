// Package cmd
package cmd

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/controller/maintenance"
	"GeoDigitalMap-messageRelayService/internal/middleware/persistence"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func CreateRestAPIServer(ctx context.Context) *ghttp.Server {
	ser := g.Server(consts.RestAPIService)
	ser.SetLogger(g.Log(consts.RestAPILogger))
	// 主路由定义
	ser.Group(consts.APIROOT, func(group *ghttp.RouterGroup) {
		// 添加了一个全局中间件，处理响应数据
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 注册 V1 子路由
		registerRestV1Routes(group)
		// TODO 未来可以添加 V2... 等这样的设计模块子路由
	})
	//ser.SetIndexFolder(true)
	//ser.SetServerRoot("/resource/public/resource")
	ser.AddStaticPath("/swagger/js", "/resource/public/resource/js")
	ser.AddStaticPath("/swagger/css", "/resource/public/resource/css")
	ser.SetSwaggerUITemplate(consts.SwaggerUITemplate)

	// 检查数据库是否能连接
	err := persistence.ConnDb(ctx)
	if err != nil {
		panic(err)
	}
	//ser.SetGraceful(true)
	//ser.EnableAdmin()
	return ser
}

// registerV1Routes 将 v1 路由注册逻辑单独封装
func registerRestV1Routes(group *ghttp.RouterGroup) {
	group.Group(consts.V1, func(v1 *ghttp.RouterGroup) {
		//v1.Middleware(middleware.Auth)
		// TODO 添加自定义的中间件钩子

		// 注册子模块的路由
		registerMaintenanceRoutes(v1)
	})
}

// registerMaintenanceRoutes 系统运维模块路由的封装
func registerMaintenanceRoutes(group *ghttp.RouterGroup) {
	group.Group(consts.MAINTENANCEMODULE, func(group *ghttp.RouterGroup) {
		group.Bind(maintenance.NewV1())
	})
}
