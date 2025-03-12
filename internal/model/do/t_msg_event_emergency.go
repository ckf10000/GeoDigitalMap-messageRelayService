// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventEmergency is the golang structure of table t_msg_event_emergency for DAO operations like Where/Data.
type TMsgEventEmergency struct {
	g.Meta    `orm:"table:t_msg_event_emergency, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	Alert     interface{} // 应急警报等级或类型(如"critical"/"warning")，默认空字符串
	Cancel    interface{} // 应急信标取消状态(true=已取消，false=有效)，默认false
	Type      interface{} // 应急事件分类类型(如医疗/消防/警务)，默认"constructor"
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
