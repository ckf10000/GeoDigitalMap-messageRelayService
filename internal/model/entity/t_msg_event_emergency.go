// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventEmergency is the golang structure for table t_msg_event_emergency.
type TMsgEventEmergency struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                                 // 事件唯一标识符(主键)
	Alert     string `json:"alert"     orm:"alert"      description:"应急警报等级或类型(如\"critical\"/\"warning\")，默认空字符串"` // 应急警报等级或类型(如"critical"/"warning")，默认空字符串
	Cancel    string `json:"cancel"    orm:"cancel"     description:"应急信标取消状态(true=已取消，false=有效)，默认false"`         // 应急信标取消状态(true=已取消，false=有效)，默认false
	Type      string `json:"type"      orm:"type"       description:"应急事件分类类型(如医疗/消防/警务)，默认\"constructor\""`       // 应急事件分类类型(如医疗/消防/警务)，默认"constructor"
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`                             // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`                             // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`                     // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`                     // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`                          // 软删标志，默认false表示有效数据
}
