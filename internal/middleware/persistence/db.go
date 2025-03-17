// Package persistence
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     db.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-16 23:54:28
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package persistence

import (
	"GeoDigitalMap-messageRelayService/internal/consts"
	"context"
	"errors"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

// ConnDb 检查数据库连接是否正常
func ConnDb(ctx context.Context) error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("connect to the database failed")
	}
	g.Log(consts.RestAPILogger).Info(ctx, "database connect successful")
	return nil
}
