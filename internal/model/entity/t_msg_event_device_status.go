// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventDeviceStatus is the golang structure for table t_msg_event_device_status.
type TMsgEventDeviceStatus struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                // 事件唯一标识符(主键)
	Battery   string `json:"battery"   orm:"battery"    description:"设备电量百分比(长度限制100字符)，默认空字符串"`  // 设备电量百分比(长度限制100字符)，默认空字符串
	Readiness string `json:"readiness" orm:"readiness"  description:"设备就绪状态标识(长度限制100字符)，默认空字符串"` // 设备就绪状态标识(长度限制100字符)，默认空字符串
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`            // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`            // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`    // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`    // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`         // 软删标志，默认false表示有效数据
}
