// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventColor is the golang structure of table t_msg_event_color for DAO operations like Where/Data.
type TMsgEventColor struct {
	g.Meta    `orm:"table:t_msg_event_color, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	Argb      interface{} // argb颜色编码(格式：#aarrggbb 或 #rrggbb)，如#ff00ff00表示不透明品红色
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
