// Package utils
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     config.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 16:10:46
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func GetServiceEnableStatus(ctx context.Context, serviceName string) bool {
	isEnableKey := fmt.Sprintf("server.%s.isEnable", serviceName)
	gvar, err := g.Cfg().Get(ctx, isEnableKey)
	if err != nil {
		return true
	}
	return gvar.Bool()
}
