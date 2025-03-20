// Package auth
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     socket.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-15 19:46:19
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package auth

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ClientAuth(r *ghttp.Request) (string, string, error) {
	ctx := r.Context()
	userId := r.GetQuery("userId")
	token := r.GetQuery("token")
	if userId.IsEmpty() {
		err := "userId is required"
		g.Log(consts.SocketLogger).Error(ctx, err)
		return "", "", gerror.New(err)
	}
	if token.IsEmpty() {
		err := "token is required"
		g.Log(consts.SocketLogger).Error(ctx, err)
		return "", "", gerror.New(err)
	}
	// TODO 可扩展添加JWT验证逻辑，目前假设都有效
	r.SetCtxVar(consts.ContextUserIdKey, userId.String())
	r.SetCtxVar(consts.ContextTokenKey, token.String())
	return userId.String(), token.String(), nil
}

func ClientAuthMiddleware(r *ghttp.Request) {
	_, _, err := ClientAuth(r)
	if err != nil {
		//r.Response.WriteHeader(400)
		// 终止后续执行
		r.Response.WriteExit(err.Error())
		//r.Exit()
		return
	}
	// 认证成功，继续执行下一个中间件或路由
	r.Middleware.Next()
}

func PeerAuthMiddleware(r *ghttp.Request) {
	// TODO 认证成功，继续执行下一个中间件或路由
	r.Middleware.Next()
}
