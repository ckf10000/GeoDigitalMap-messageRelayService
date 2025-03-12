// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventPrecisionLocation is the golang structure of table t_msg_event_precision_location for DAO operations like Where/Data.
type TMsgEventPrecisionLocation struct {
	g.Meta      `orm:"table:t_msg_event_precision_location, do:true"`
	EventId     interface{} // 事件唯一标识符(主键)
	AltSrc      interface{} // 高度数据源标识(如dted0)，默认空字符串表示未指定
	GeoPointSrc interface{} // 地理坐标数据源标识(如gps/北斗)，默认空字符串
	UpdateBy    interface{} // 更新者id，默认0表示系统操作
	CreateBy    interface{} // 创建者id，默认0表示系统创建
	CreatedAt   interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt   interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted   interface{} // 软删标志，默认false表示有效数据
}
