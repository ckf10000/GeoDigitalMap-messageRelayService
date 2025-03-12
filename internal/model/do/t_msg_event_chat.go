// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventChat is the golang structure of table t_msg_event_chat for DAO operations like Where/Data.
type TMsgEventChat struct {
	g.Meta         `orm:"table:t_msg_event_chat, do:true"`
	EventId        interface{} // 事件唯一标识符(主键)
	ChatRoom       interface{} // 聊天室唯一标识符(如"cmd-center-01")，默认空字符串
	GroupOwner     interface{} // 群组管理员标识符(uid格式)，默认空字符串
	Id             interface{} // 聊天记录唯一id(uuid格式)，默认空字符串
	Parent         interface{} // 父级聊天会话标识符，默认空字符串
	SenderCallSign interface{} // 发送方呼号标识(如"alpha-01")，默认空字符串
	UpdateBy       interface{} // 更新者id，默认0表示系统操作
	CreateBy       interface{} // 创建者id，默认0表示系统创建
	CreatedAt      interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt      interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted      interface{} // 软删标志，默认false表示有效数据
}
