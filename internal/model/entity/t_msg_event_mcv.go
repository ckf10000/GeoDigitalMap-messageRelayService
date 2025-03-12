// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventMcv is the golang structure for table t_msg_event_mcv.
type TMsgEventMcv struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                     // 事件唯一标识符(主键)
	Device    string `json:"device"    orm:"device"     description:"物理设备型号(长度限制100字符)，默认空字符串"`        // 物理设备型号(长度限制100字符)，默认空字符串
	Os        string `json:"os"        orm:"os"         description:"设备操作系统类型(长度限制100字符)，默认空字符串"`      // 设备操作系统类型(长度限制100字符)，默认空字符串
	Platform  string `json:"platform"  orm:"platform"   description:"消息事件客户端平台变体名称(长度限制100字符)，默认空字符串"` // 消息事件客户端平台变体名称(长度限制100字符)，默认空字符串
	Version   string `json:"version"   orm:"version"    description:"消息事件客户端软件版本号(长度限制100字符)，默认空字符串"`  // 消息事件客户端软件版本号(长度限制100字符)，默认空字符串
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`                 // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`                 // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`         // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`         // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`              // 软删标志，默认false表示有效数据
}
