// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventSummary is the golang structure of table t_msg_event_summary for DAO operations like Where/Data.
type TMsgEventSummary struct {
	g.Meta    `orm:"table:t_msg_event_summary, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	How       interface{} // 事件触发方式标识(长度限制100字符)，默认空字符串
	Time      interface{} // 事件发生时间(unix时间戳，秒级)，默认0表示未记录
	Start     interface{} // 事件有效开始时间(unix时间戳，秒级)，默认0表示立即生效
	Stale     interface{} // 事件失效时间(unix时间戳，秒级)，默认0表示永不过期
	Type      interface{} // 事件分类类型(长度限制100字符)，默认空字符串
	Version   interface{} // 事件版本标识符(长度限制100字符)，默认空字符串
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
