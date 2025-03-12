// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventHow is the golang structure for table t_msg_event_how.
type TMsgEventHow struct {
	Id          int    `json:"id"          orm:"id"          description:"自增主键，唯一标识记录"`               // 自增主键，唯一标识记录
	Code        string `json:"code"        orm:"code"        description:"事件类型编码(长度限制50字符)，不可为空"`     // 事件类型编码(长度限制50字符)，不可为空
	How         string `json:"how"         orm:"how"         description:"事件触发方式(长度限制50字符)，不可为空"`     // 事件触发方式(长度限制50字符)，不可为空
	Alias       string `json:"alias"       orm:"alias"       description:"事件别名(长度限制100字符)，默认空字符串"`    // 事件别名(长度限制100字符)，默认空字符串
	Description string `json:"description" orm:"description" description:"事件详细描述，默认空文本"`              // 事件详细描述，默认空文本
	UpdateBy    int    `json:"updateBy"    orm:"update_by"   description:"更新者id，默认0表示系统操作"`           // 更新者id，默认0表示系统操作
	CreateBy    int    `json:"createBy"    orm:"create_by"   description:"创建者id，默认0表示系统创建"`           // 创建者id，默认0表示系统创建
	CreatedAt   int64  `json:"createdAt"   orm:"created_at"  description:"创建时间(unix时间戳，秒级)，默认当前时间"`   // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt   int64  `json:"updatedAt"   orm:"updated_at"  description:"更新时间(unix时间戳，秒级)，默认当前时间"`   // 更新时间(unix时间戳，秒级)，默认当前时间
	DeletedAt   int64  `json:"deletedAt"   orm:"deleted_at"  description:"软删时间(unix时间戳，秒级)，默认0表示未删除"` // 软删时间(unix时间戳，秒级)，默认0表示未删除
	IsDeleted   bool   `json:"isDeleted"   orm:"is_deleted"  description:"软删标志，默认false表示有效数据"`        // 软删标志，默认false表示有效数据
}
