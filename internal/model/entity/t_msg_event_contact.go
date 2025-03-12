// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TMsgEventContact is the golang structure for table t_msg_event_contact.
type TMsgEventContact struct {
	EventId     string `json:"eventId"     orm:"event_id"    description:"事件唯一标识符(主键)"`                               // 事件唯一标识符(主键)
	Callsign    string `json:"callsign"    orm:"callsign"    description:"通信呼号标识(如\"alpha-01\")，默认空字符串"`              // 通信呼号标识(如"alpha-01")，默认空字符串
	Dsn         string `json:"dsn"         orm:"dsn"         description:"国防交换网络编号(defense switched network)，默认空字符串"` // 国防交换网络编号(defense switched network)，默认空字符串
	Email       string `json:"email"       orm:"email"       description:"电子邮箱地址(需符合name@domain格式)，默认空字符串"`           // 电子邮箱地址(需符合name@domain格式)，默认空字符串
	Endpoint    string `json:"endpoint"    orm:"endpoint"    description:"api终端地址(ip:port格式)，默认空字符串"`                 // api终端地址(ip:port格式)，默认空字符串
	Freq        string `json:"freq"        orm:"freq"        description:"通信频率(单位：mhz，如\"121.500\")，默认空字符串"`          // 通信频率(单位：mhz，如"121.500")，默认空字符串
	Hostname    string `json:"hostname"    orm:"hostname"    description:"网络主机名(需dns可解析)，默认空字符串"`                     // 网络主机名(需dns可解析)，默认空字符串
	Iconsetpath string `json:"iconsetpath" orm:"iconsetpath" description:"图标资源存储路径(相对路径)，默认空字符串"`                     // 图标资源存储路径(相对路径)，默认空字符串
	Modulation  string `json:"modulation"  orm:"modulation"  description:"无线电调制方式(am=调幅/fm=调频)"`                      // 无线电调制方式(am=调幅/fm=调频)
	Phone       string `json:"phone"       orm:"phone"       description:"联系电话号码(国际格式如+86-13912345678)，默认空字符串"`       // 联系电话号码(国际格式如+86-13912345678)，默认空字符串
	Version     string `json:"version"     orm:"version"     description:"联系人模式版本号(格式：v1.2.3)，默认空字符串"`                // 联系人模式版本号(格式：v1.2.3)，默认空字符串
	UpdateBy    int    `json:"updateBy"    orm:"update_by"   description:"更新者id，默认0表示系统操作"`                           // 更新者id，默认0表示系统操作
	CreateBy    int    `json:"createBy"    orm:"create_by"   description:"创建者id，默认0表示系统创建"`                           // 创建者id，默认0表示系统创建
	CreatedAt   int64  `json:"createdAt"   orm:"created_at"  description:"创建时间(unix时间戳，秒级)，默认当前时间"`                   // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt   int64  `json:"updatedAt"   orm:"updated_at"  description:"更新时间(unix时间戳，秒级)，默认当前时间"`                   // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted   bool   `json:"isDeleted"   orm:"is_deleted"  description:"软删标志，默认false表示有效数据"`                        // 软删标志，默认false表示有效数据
}
