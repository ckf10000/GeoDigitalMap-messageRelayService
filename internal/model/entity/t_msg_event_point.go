// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventPoint is the golang structure for table t_msg_event_point.
type TMsgEventPoint struct {
	EventId   string `json:"eventId"   orm:"event_id"   description:"事件唯一标识符(主键)"`                                        // 事件唯一标识符(主键)
	Lat       string `json:"lat"       orm:"lat"        description:"纬度坐标(wgs84坐标系)，默认0表示未记录，单位：度"`                       // 纬度坐标(wgs84坐标系)，默认0表示未记录，单位：度
	Lon       string `json:"lon"       orm:"lon"        description:"经度坐标(gs84坐标系)，默认0表示未记录，单位：度"`                        // 经度坐标(gs84坐标系)，默认0表示未记录，单位：度
	Hae       string `json:"hae"       orm:"hae"        description:"海拔高度(height above ellipsoid)，默认9999999.0表示无效值，单位：米"` // 海拔高度(height above ellipsoid)，默认9999999.0表示无效值，单位：米
	Ce        string `json:"ce"        orm:"ce"         description:"圆误差概率(circular error)，默认9999999.0表示无效值，单位：米"`        // 圆误差概率(circular error)，默认9999999.0表示无效值，单位：米
	Le        string `json:"le"        orm:"le"         description:"线误差概率(linear error)，默认9999999.0表示无效值，单位：米"`          // 线误差概率(linear error)，默认9999999.0表示无效值，单位：米
	UpdateBy  int    `json:"updateBy"  orm:"update_by"  description:"更新者id，默认0表示系统操作"`                                    // 更新者id，默认0表示系统操作
	CreateBy  int    `json:"createBy"  orm:"create_by"  description:"创建者id，默认0表示系统创建"`                                    // 创建者id，默认0表示系统创建
	CreatedAt int64  `json:"createdAt" orm:"created_at" description:"创建时间(unix时间戳，秒级)，默认当前时间"`                            // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" description:"更新时间(unix时间戳，秒级)，默认当前时间"`                            // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted bool   `json:"isDeleted" orm:"is_deleted" description:"软删标志，默认false表示有效数据"`                                 // 软删标志，默认false表示有效数据
}
