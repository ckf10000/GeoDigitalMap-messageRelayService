// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventVideo is the golang structure for table t_msg_event_video.
type TMsgEventVideo struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                                        // 事件唯一标识符(主键)
	Sensor    string `json:"sensor"    orm:"sensor"     description:"视频传感器标识符(长度限制100字符)，默认空字符串"`                         // 视频传感器标识符(长度限制100字符)，默认空字符串
	Spi       string `json:"spi"       orm:"spi"        description:"安全参数索引(security parameters index，长度限制100字符)，默认空字符串"` // 安全参数索引(security parameters index，长度限制100字符)，默认空字符串
	Url       string `json:"url"       orm:"url"        description:"视频资源访问地址(长度限制100字符)，默认空字符串"`                         // 视频资源访问地址(长度限制100字符)，默认空字符串
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`                                    // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`                                    // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`                            // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`                            // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`                                 // 软删标志，默认false表示有效数据
}
