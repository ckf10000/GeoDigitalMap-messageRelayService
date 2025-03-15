// Package main
package main

import (
	"GeoDigitalMap-messageRelayService/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
