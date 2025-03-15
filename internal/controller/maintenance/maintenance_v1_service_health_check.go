package maintenance

import (
	"GeoDigitalMap-messageRelayService/api/maintenance/v1"
	"context"
)

func (c *ControllerV1) ServiceHealthCheck(ctx context.Context, req *v1.ServiceHealthCheckReq) (res *v1.ServiceHealthCheckRes, err error) {
	str := "The current service status is very healthy"
	//g.Log(consts.RestAPIService).Info(ctx, str)  // 健康检查的时候，频率太高，容易打印很多日志，此处不建议打印日志
	return &v1.ServiceHealthCheckRes{Description: str}, err
}
