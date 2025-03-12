// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventRemarks is the golang structure for table t_msg_event_remarks.
type TMsgEventRemarks struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                    // 事件唯一标识符(主键)
	Keywords  string `json:"keywords"  orm:"keywords"   description:"逗号分隔的备注关键词(如\"任务a,紧急\")，默认空字符串"` // 逗号分隔的备注关键词(如"任务a,紧急")，默认空字符串
	Source    string `json:"source"    orm:"source"     description:"备注来源标识符(发送方uid)，默认空字符串"`         // 备注来源标识符(发送方uid)，默认空字符串
	Time      int64  `json:"time"      orm:"time"       description:"备注生成时间(unix时间戳，秒级)，默认0表示未记录"`    // 备注生成时间(unix时间戳，秒级)，默认0表示未记录
	ToField   string `json:"toField"   orm:"to_field"   description:"目标接收方uid，默认空字符串表示广播"`            // 目标接收方uid，默认空字符串表示广播
	Version   string `json:"version"   orm:"version"    description:"备注格式版本标识，默认空字符串"`                // 备注格式版本标识，默认空字符串
	Intag     string `json:"intag"     orm:"intag"      description:"扩展标签信息(最大长度1000字符)，默认空字符串"`      // 扩展标签信息(最大长度1000字符)，默认空字符串
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`                // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`                // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`        // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`        // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`             // 软删标志，默认false表示有效数据
}
