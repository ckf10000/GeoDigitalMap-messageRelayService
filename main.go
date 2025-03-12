// Package main
package main

import (
	"GeoDigitalMap-messageRelayService/internal/cmd"
	"context"
	"errors"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// connDb 检查数据库连接是否正常
func connDb(ctx context.Context) error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("connection to the database failed")
	}
	g.Log().Info(ctx, "database connection successful")
	return nil
}
func main() {
	var err error
	ctx := context.TODO()
	// 检查数据库是否能连接
	err = connDb(ctx)
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(gctx.GetInitCtx())
}
