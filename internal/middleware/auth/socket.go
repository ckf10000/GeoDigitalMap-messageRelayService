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
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

func ClientAuth(r *ghttp.Request) {
	clientID := r.GetQuery("clientID")
	if clientID.IsEmpty() {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "clientID required")
	}

	// 可扩展添加JWT验证逻辑
	r.SetCtxVar("clientID", clientID.String())
	r.Middleware.Next()
}
