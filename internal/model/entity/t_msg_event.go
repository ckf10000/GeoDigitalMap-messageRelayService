// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEvent is the golang structure for table t_msg_event.
type TMsgEvent struct {
	Id             string `json:"id"             orm:"id"              description:"事件唯一标识符(主键)"`                  // 事件唯一标识符(主键)
	XmlString      string `json:"xmlString"      orm:"xml_string"      description:"事件xml原始数据(长度限制1000字符)，默认空字符串"` // 事件xml原始数据(长度限制1000字符)，默认空字符串
	Sender         string `json:"sender"         orm:"sender"          description:"事件发送方标识符(长度限制100字符)，不可为空"`     // 事件发送方标识符(长度限制100字符)，不可为空
	Receiver       string `json:"receiver"       orm:"receiver"        description:"事件接收方标识符(长度限制100字符)，默认空字符串"`   // 事件接收方标识符(长度限制100字符)，默认空字符串
	SendResult     int    `json:"sendResult"     orm:"send_result"     description:"发送结果状态码(1-成功，0-失败)，默认1"`       // 发送结果状态码(1-成功，0-失败)，默认1
	OccurrenceTime int64  `json:"occurrenceTime" orm:"occurrence_time" description:"事件发生时间(unix时间戳，秒级)，默认0表示未记录"`  // 事件发生时间(unix时间戳，秒级)，默认0表示未记录
	Checksum       string `json:"checksum"       orm:"checksum"        description:"数据校验和(md5哈希值)，默认空字符串"`         // 数据校验和(md5哈希值)，默认空字符串
	UpdateBy       int    `json:"updateBy"       orm:"update_by"       description:"更新者id，默认0表示系统操作"`              // 更新者id，默认0表示系统操作
	CreateBy       int    `json:"createBy"       orm:"create_by"       description:"创建者id，默认0表示系统创建"`              // 创建者id，默认0表示系统创建
	CreatedAt      int64  `json:"createdAt"      orm:"created_at"      description:"创建时间(unix时间戳，秒级)，默认当前时间"`      // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt      int64  `json:"updatedAt"      orm:"updated_at"      description:"更新时间(unix时间戳，秒级)，默认当前时间"`      // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted      bool   `json:"isDeleted"      orm:"is_deleted"      description:"软删标志，默认false表示有效数据"`           // 软删标志，默认false表示有效数据
}
