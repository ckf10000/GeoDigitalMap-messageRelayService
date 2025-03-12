// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventTrack is the golang structure of table t_msg_event_track for DAO operations like Where/Data.
type TMsgEventTrack struct {
	g.Meta    `orm:"table:t_msg_event_track, do:true"`
	EventId   interface{} // 事件唯一标识符(主键)
	Course    interface{} // 移动方向角度(单位：度，长度限制100字符)，默认空字符串
	Ecourse   interface{} // 方向角测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Eslope    interface{} // 坡度测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Espeed    interface{} // 速度测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Speed     interface{} // 移动速度值(单位：米/秒，长度限制100字符)，默认空字符串
	Version   interface{} //
	UpdateBy  interface{} // 更新者id，默认0表示系统操作
	CreateBy  interface{} // 创建者id，默认0表示系统创建
	CreatedAt interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted interface{} // 软删标志，默认false表示有效数据
}
