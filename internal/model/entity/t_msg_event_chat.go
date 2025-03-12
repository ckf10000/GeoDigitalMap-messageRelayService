// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventChat is the golang structure for table t_msg_event_chat.
type TMsgEventChat struct {
	EventId        string `json:"eventId"        orm:"event_id"         description:"事件唯一标识符(主键)"`                         // 事件唯一标识符(主键)
	ChatRoom       string `json:"chatRoom"       orm:"chat_room"        description:"聊天室唯一标识符(如\"cmd-center-01\")，默认空字符串"` // 聊天室唯一标识符(如"cmd-center-01")，默认空字符串
	GroupOwner     string `json:"groupOwner"     orm:"group_owner"      description:"群组管理员标识符(uid格式)，默认空字符串"`              // 群组管理员标识符(uid格式)，默认空字符串
	Id             string `json:"id"             orm:"id"               description:"聊天记录唯一id(uuid格式)，默认空字符串"`             // 聊天记录唯一id(uuid格式)，默认空字符串
	Parent         string `json:"parent"         orm:"parent"           description:"父级聊天会话标识符，默认空字符串"`                    // 父级聊天会话标识符，默认空字符串
	SenderCallSign string `json:"senderCallSign" orm:"sender_call_sign" description:"发送方呼号标识(如\"alpha-01\")，默认空字符串"`       // 发送方呼号标识(如"alpha-01")，默认空字符串
	UpdateBy       int    `json:"updateBy"       orm:"update_by"        description:"更新者id，默认0表示系统操作"`                     // 更新者id，默认0表示系统操作
	CreateBy       int    `json:"createBy"       orm:"create_by"        description:"创建者id，默认0表示系统创建"`                     // 创建者id，默认0表示系统创建
	CreatedAt      int64  `json:"createdAt"      orm:"created_at"       description:"创建时间(unix时间戳，秒级)，默认当前时间"`             // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt      int64  `json:"updatedAt"      orm:"updated_at"       description:"更新时间(unix时间戳，秒级)，默认当前时间"`             // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted      bool   `json:"isDeleted"      orm:"is_deleted"       description:"软删标志，默认false表示有效数据"`                  // 软删标志，默认false表示有效数据
}
