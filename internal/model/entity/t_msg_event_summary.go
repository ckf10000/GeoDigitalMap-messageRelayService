// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventSummary is the golang structure for table t_msg_event_summary.
type TMsgEventSummary struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                    // 事件唯一标识符(主键)
	How       string `json:"how"       orm:"how"        description:"事件触发方式标识(长度限制100字符)，默认空字符串"`     // 事件触发方式标识(长度限制100字符)，默认空字符串
	Time      int64  `json:"time"      orm:"time"       description:"事件发生时间(unix时间戳，秒级)，默认0表示未记录"`    // 事件发生时间(unix时间戳，秒级)，默认0表示未记录
	Start     int64  `json:"start"     orm:"start"      description:"事件有效开始时间(unix时间戳，秒级)，默认0表示立即生效"` // 事件有效开始时间(unix时间戳，秒级)，默认0表示立即生效
	Stale     int64  `json:"stale"     orm:"stale"      description:"事件失效时间(unix时间戳，秒级)，默认0表示永不过期"`   // 事件失效时间(unix时间戳，秒级)，默认0表示永不过期
	Type      string `json:"type"      orm:"type"       description:"事件分类类型(长度限制100字符)，默认空字符串"`       // 事件分类类型(长度限制100字符)，默认空字符串
	Version   string `json:"version"   orm:"version"    description:"事件版本标识符(长度限制100字符)，默认空字符串"`      // 事件版本标识符(长度限制100字符)，默认空字符串
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`                // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`                // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`        // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`        // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`             // 软删标志，默认false表示有效数据
}
