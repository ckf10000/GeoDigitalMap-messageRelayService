// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventTrack is the golang structure for table t_msg_event_track.
type TMsgEventTrack struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                     // 事件唯一标识符(主键)
	Course    string `json:"course"    orm:"course"     description:"移动方向角度(单位：度，长度限制100字符)，默认空字符串"`   // 移动方向角度(单位：度，长度限制100字符)，默认空字符串
	Ecourse   string `json:"ecourse"   orm:"ecourse"    description:"方向角测量误差(1σ标准差，长度限制100字符)，默认空字符串"` // 方向角测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Eslope    string `json:"eslope"    orm:"eslope"     description:"坡度测量误差(1σ标准差，长度限制100字符)，默认空字符串"`  // 坡度测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Espeed    string `json:"espeed"    orm:"espeed"     description:"速度测量误差(1σ标准差，长度限制100字符)，默认空字符串"`  // 速度测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Speed     string `json:"speed"     orm:"speed"      description:"移动速度值(单位：米/秒，长度限制100字符)，默认空字符串"`  // 移动速度值(单位：米/秒，长度限制100字符)，默认空字符串
	Version   string `json:"version"   orm:"version"    description:""`                                //
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`                 // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`                 // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`         // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`         // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`              // 软删标志，默认false表示有效数据
}
