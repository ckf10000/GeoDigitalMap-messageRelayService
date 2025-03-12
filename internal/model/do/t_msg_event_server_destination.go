// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventServerDestination is the golang structure of table t_msg_event_server_destination for DAO operations like Where/Data.
type TMsgEventServerDestination struct {
	g.Meta       `orm:"table:t_msg_event_server_destination, do:true"`
	EventId      interface{} // 事件唯一标识符(主键)
	Destinations interface{} // 目标地址组合(格式：ip:端口:协议:设备id)，不可为空
	UpdateBy     interface{} // 更新者id，默认0表示系统操作
	CreateBy     interface{} // 创建者id，默认0表示系统创建
	CreatedAt    interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt    interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted    interface{} // 软删标志，默认false表示有效数据
}
