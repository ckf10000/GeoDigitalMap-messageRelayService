// Package cmd
package cmd

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"GeoDigitalMap-messageRelayService/internal/controller/maintenance"
	"context"
	"errors"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// connDb 检查数据库连接是否正常
func connDb(ctx context.Context) error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("connection to the database failed")
	}
	g.Log(consts.RestAPIService).Info(ctx, "database connection successful")
	return nil
}

func RestAPIServer(ctx context.Context) *ghttp.Server {
	ser := g.Server(consts.RestAPIService)
	ser.SetLogger(g.Log(consts.RestAPIService))
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

	// 检查数据库是否能连接
	err := connDb(ctx)
	if err != nil {
		panic(err)
	}
	ser.SetGraceful(true)
	ser.EnableAdmin()
	return ser
}

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
