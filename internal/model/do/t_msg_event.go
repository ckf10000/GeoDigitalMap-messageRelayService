// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEvent is the golang structure of table t_msg_event for DAO operations like Where/Data.
type TMsgEvent struct {
	g.Meta         `orm:"table:t_msg_event, do:true"`
	Id             interface{} // 事件唯一标识符(主键)
	XmlString      interface{} // 事件xml原始数据(长度限制1000字符)，默认空字符串
	Sender         interface{} // 事件发送方标识符(长度限制100字符)，不可为空
	Receiver       interface{} // 事件接收方标识符(长度限制100字符)，默认空字符串
	SendResult     interface{} // 发送结果状态码(1-成功，0-失败)，默认1
	OccurrenceTime interface{} // 事件发生时间(unix时间戳，秒级)，默认0表示未记录
	Checksum       interface{} // 数据校验和(md5哈希值)，默认空字符串
	UpdateBy       interface{} // 更新者id，默认0表示系统操作
	CreateBy       interface{} // 创建者id，默认0表示系统创建
	CreatedAt      interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt      interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted      interface{} // 软删标志，默认false表示有效数据
}
