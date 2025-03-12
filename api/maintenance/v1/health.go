// Package v1
package v1

import "github.com/gogf/gf/v2/frame/g"

// ServiceHealthCheckReq 服务健康检查的请求
type ServiceHealthCheckReq struct {
	g.Meta `path:"healthCheck" method:"get" tags:"系统维护" summary:"服务健康检查"`
}

// ServiceHealthCheckRes 服务健康检查的响应
type ServiceHealthCheckRes struct {
	Description string `json:"description"` // 描述
}
