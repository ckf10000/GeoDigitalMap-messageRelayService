// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventPoint is the golang structure of table t_msg_event_point for DAO operations like Where/Data.
type TMsgEventPoint struct {
	g.Meta    `orm:"table:t_msg_event_point, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	Lat       interface{} // 纬度坐标(wgs84坐标系)，默认0表示未记录，单位：度
	Lon       interface{} // 经度坐标(gs84坐标系)，默认0表示未记录，单位：度
	Hae       interface{} // 海拔高度(height above ellipsoid)，默认9999999.0表示无效值，单位：米
	Ce        interface{} // 圆误差概率(circular error)，默认9999999.0表示无效值，单位：米
	Le        interface{} // 线误差概率(linear error)，默认9999999.0表示无效值，单位：米
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
