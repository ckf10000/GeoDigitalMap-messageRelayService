// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventSensor is the golang structure of table t_msg_event_sensor for DAO operations like Where/Data.
type TMsgEventSensor struct {
	g.Meta    `orm:"table:t_msg_event_sensor, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	Elevation interface{} // 传感器海拔高度(单位：米)，默认空字符串
	Vfov      interface{} // 垂直视场角(vertical fov)，默认空字符串
	North     interface{} // 北向校准参数，默认空字符串
	Roll      interface{} // 传感器滚动角参数，默认空字符串
	Range     interface{} // 传感器监测范围，默认空字符串
	Azimuth   interface{} // 传感器方位角参数，默认空字符串
	Model     interface{} // 传感器型号，默认空字符串
	Fov       interface{} // 传感器视场角，默认空字符串
	Type      interface{} // 传感器类型，默认空字符串
	Version   interface{} // 传感器版本号，默认空字符串
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
