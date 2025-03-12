// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventGroup is the golang structure of table t_msg_event_group for DAO operations like Where/Data.
type TMsgEventGroup struct {
	g.Meta    `orm:"table:t_msg_event_group, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	Name      interface{} // 群组名称(唯一标识，如"指挥中心-01")，不可为空
	Role      interface{} // 成员角色(admin=管理员/member=普通成员)，默认member
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
