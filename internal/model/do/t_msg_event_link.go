// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventLink is the golang structure of table t_msg_event_link for DAO operations like Where/Data.
type TMsgEventLink struct {
	g.Meta         `orm:"table:t_msg_event_link, do:true"`
	EventId        interface{} // 事件唯一标识符(主键)
	Mime           interface{} // 资源mime类型(rfc 3023标准，如 application/xml/text/xml)，默认空字符串
	ParentCallsign interface{} // 上级节点呼号标识(用于层级关系追踪)，默认空字符串
	ProductionTime interface{} // 资源生成时间(unix时间戳，秒级)，默认0表示未记录
	Relation       interface{} // 关联关系类型(cot标准：subject/object/indirect-object)，默认空字符串
	Remarks        interface{} // 链接附加说明(自由文本，最大长度100字符)，默认空字符串
	Type           interface{} // 链接资源分类(如影像/文档/传感器数据)，默认空字符串
	Uid            interface{} // 关联资源全局唯一标识符(遵循cot uid规范)，默认空字符串
	Url            interface{} // 资源访问地址(url格式，支持短链接)，默认空字符串
	Version        interface{} // 链接模式版本标识(格式：主版本.次版本)，默认空字符串
	UpdateBy       interface{} // 更新者id，默认0表示系统操作
	CreateBy       interface{} // 创建者id，默认0表示系统创建
	CreatedAt      interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt      interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted      interface{} // 软删标志，默认false表示有效数据
}
