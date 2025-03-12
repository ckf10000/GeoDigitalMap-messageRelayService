// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventRemarks is the golang structure of table t_msg_event_remarks for DAO operations like Where/Data.
type TMsgEventRemarks struct {
	g.Meta    `orm:"table:t_msg_event_remarks, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	Keywords  interface{} // 逗号分隔的备注关键词(如"任务a,紧急")，默认空字符串
	Source    interface{} // 备注来源标识符(发送方uid)，默认空字符串
	Time      interface{} // 备注生成时间(unix时间戳，秒级)，默认0表示未记录
	ToField   interface{} // 目标接收方uid，默认空字符串表示广播
	Version   interface{} // 备注格式版本标识，默认空字符串
	Intag     interface{} // 扩展标签信息(最大长度1000字符)，默认空字符串
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
