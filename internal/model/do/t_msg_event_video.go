// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventVideo is the golang structure of table t_msg_event_video for DAO operations like Where/Data.
type TMsgEventVideo struct {
	g.Meta    `orm:"table:t_msg_event_video, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	Sensor    interface{} // 视频传感器标识符(长度限制100字符)，默认空字符串
	Spi       interface{} // 安全参数索引(security parameters index，长度限制100字符)，默认空字符串
	Url       interface{} // 视频资源访问地址(长度限制100字符)，默认空字符串
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
