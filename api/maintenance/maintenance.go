// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package maintenance

import (
	"context"

	"GeoDigitalMap-messageRelayService/api/maintenance/v1"
)

type IMaintenanceV1 interface {
	ServiceHealthCheck(ctx context.Context, req *v1.ServiceHealthCheckReq) (res *v1.ServiceHealthCheckRes, err error)
}
