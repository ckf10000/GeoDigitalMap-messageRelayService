// Package manager
/*********************************************************************************************************************
* ProjectName:  GeoDigitalMap-messageRelayService
* FileName:     manager.go
* Description:  TODO
* Author:       zhouhanlin
* CreateDate:   2025-03-19 23:34:29
* Copyright ©2011-2025. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package manager

import (
	"context"

	"GeoDigitalMap-messageRelayService/api/manager/v1"
)

type IManagerV1 interface {
	AddPeer(ctx context.Context, req *v1.AddPeerReq) (res *v1.AddPeerRes, err error)
}
