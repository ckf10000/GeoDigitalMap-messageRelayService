// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventType is the golang structure of table t_msg_event_type for DAO operations like Where/Data.
type TMsgEventType struct {
	g.Meta      `orm:"table:t_msg_event_type, do:true"`
	Id          interface{} // 自增主键，唯一标识记录
	EventType   interface{} // 事件类型编码(长度限制50字符)，不可为空
	Alias       interface{} // 事件类型别名(长度限制100字符)，默认空字符串
	Description interface{} // 事件类型详细描述，默认空文本
	UpdateBy    interface{} // 更新者id，默认0表示系统操作
	CreateBy    interface{} // 创建者id，默认0表示系统创建
	CreatedAt   interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt   interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	DeletedAt   interface{} // 软删时间(unix时间戳，秒级)，默认0表示未删除
	IsDeleted   interface{} // 软删标志，默认false表示有效数据
}
