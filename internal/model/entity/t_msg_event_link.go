// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventLink is the golang structure for table t_msg_event_link.
type TMsgEventLink struct {
	EventId        string `json:"eventId"        orm:"event_id"        description:"事件唯一标识符(主键)"`                                            // 事件唯一标识符(主键)
	Mime           string `json:"mime"           orm:"mime"            description:"资源mime类型(rfc 3023标准，如 application/xml/text/xml)，默认空字符串"` // 资源mime类型(rfc 3023标准，如 application/xml/text/xml)，默认空字符串
	ParentCallsign string `json:"parentCallsign" orm:"parent_callsign" description:"上级节点呼号标识(用于层级关系追踪)，默认空字符串"`                              // 上级节点呼号标识(用于层级关系追踪)，默认空字符串
	ProductionTime int64  `json:"productionTime" orm:"production_time" description:"资源生成时间(unix时间戳，秒级)，默认0表示未记录"`                            // 资源生成时间(unix时间戳，秒级)，默认0表示未记录
	Relation       string `json:"relation"       orm:"relation"        description:"关联关系类型(cot标准：subject/object/indirect-object)，默认空字符串"`    // 关联关系类型(cot标准：subject/object/indirect-object)，默认空字符串
	Remarks        string `json:"remarks"        orm:"remarks"         description:"链接附加说明(自由文本，最大长度100字符)，默认空字符串"`                          // 链接附加说明(自由文本，最大长度100字符)，默认空字符串
	Type           string `json:"type"           orm:"type"            description:"链接资源分类(如影像/文档/传感器数据)，默认空字符串"`                            // 链接资源分类(如影像/文档/传感器数据)，默认空字符串
	Uid            string `json:"uid"            orm:"uid"             description:"关联资源全局唯一标识符(遵循cot uid规范)，默认空字符串"`                        // 关联资源全局唯一标识符(遵循cot uid规范)，默认空字符串
	Url            string `json:"url"            orm:"url"             description:"资源访问地址(url格式，支持短链接)，默认空字符串"`                             // 资源访问地址(url格式，支持短链接)，默认空字符串
	Version        string `json:"version"        orm:"version"         description:"链接模式版本标识(格式：主版本.次版本)，默认空字符串"`                            // 链接模式版本标识(格式：主版本.次版本)，默认空字符串
	UpdateBy       int    `json:"updateBy"       orm:"update_by"       description:"更新者id，默认0表示系统操作"`                                        // 更新者id，默认0表示系统操作
	CreateBy       int    `json:"createBy"       orm:"create_by"       description:"创建者id，默认0表示系统创建"`                                        // 创建者id，默认0表示系统创建
	CreatedAt      int64  `json:"createdAt"      orm:"created_at"      description:"创建时间(unix时间戳，秒级)，默认当前时间"`                                // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt      int64  `json:"updatedAt"      orm:"updated_at"      description:"更新时间(unix时间戳，秒级)，默认当前时间"`                                // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted      bool   `json:"isDeleted"      orm:"is_deleted"      description:"软删标志，默认false表示有效数据"`                                     // 软删标志，默认false表示有效数据
}
