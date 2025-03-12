// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventPrecisionLocation is the golang structure for table t_msg_event_precision_location.
type TMsgEventPrecisionLocation struct {
	EventId     string `json:"eventId"     orm:"event_id"      description:"事件唯一标识符(主键)"`                 // 事件唯一标识符(主键)
	AltSrc      string `json:"altSrc"      orm:"alt_src"       description:"高度数据源标识(如dted0)，默认空字符串表示未指定"` // 高度数据源标识(如dted0)，默认空字符串表示未指定
	GeoPointSrc string `json:"geoPointSrc" orm:"geo_point_src" description:"地理坐标数据源标识(如gps/北斗)，默认空字符串"`   // 地理坐标数据源标识(如gps/北斗)，默认空字符串
	UpdateBy    int    `json:"updateBy"    orm:"update_by"     description:"更新者id，默认0表示系统操作"`             // 更新者id，默认0表示系统操作
	CreateBy    int    `json:"createBy"    orm:"create_by"     description:"创建者id，默认0表示系统创建"`             // 创建者id，默认0表示系统创建
	CreatedAt   int64  `json:"createdAt"   orm:"created_at"    description:"创建时间(unix时间戳，秒级)，默认当前时间"`     // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt   int64  `json:"updatedAt"   orm:"updated_at"    description:"更新时间(unix时间戳，秒级)，默认当前时间"`     // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted   bool   `json:"isDeleted"   orm:"is_deleted"    description:"软删标志，默认false表示有效数据"`          // 软删标志，默认false表示有效数据
}
