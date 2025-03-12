// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventSensor is the golang structure for table t_msg_event_sensor.
type TMsgEventSensor struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                // 事件唯一标识符(主键)
	Elevation string `json:"elevation" orm:"elevation"  description:"传感器海拔高度(单位：米)，默认空字符串"`       // 传感器海拔高度(单位：米)，默认空字符串
	Vfov      string `json:"vfov"      orm:"vfov"       description:"垂直视场角(vertical fov)，默认空字符串"` // 垂直视场角(vertical fov)，默认空字符串
	North     string `json:"north"     orm:"north"      description:"北向校准参数，默认空字符串"`              // 北向校准参数，默认空字符串
	Roll      string `json:"roll"      orm:"roll"       description:"传感器滚动角参数，默认空字符串"`            // 传感器滚动角参数，默认空字符串
	Range     string `json:"range"     orm:"range"      description:"传感器监测范围，默认空字符串"`             // 传感器监测范围，默认空字符串
	Azimuth   string `json:"azimuth"   orm:"azimuth"    description:"传感器方位角参数，默认空字符串"`            // 传感器方位角参数，默认空字符串
	Model     string `json:"model"     orm:"model"      description:"传感器型号，默认空字符串"`               // 传感器型号，默认空字符串
	Fov       string `json:"fov"       orm:"fov"        description:"传感器视场角，默认空字符串"`              // 传感器视场角，默认空字符串
	Type      string `json:"type"      orm:"type"       description:"传感器类型，默认空字符串"`               // 传感器类型，默认空字符串
	Version   string `json:"version"   orm:"version"    description:"传感器版本号，默认空字符串"`              // 传感器版本号，默认空字符串
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`            // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`            // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`    // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`    // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`         // 软删标志，默认false表示有效数据
}
