// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventContact is the golang structure of table t_msg_event_contact for DAO operations like Where/Data.
type TMsgEventContact struct {
	g.Meta      `orm:"table:t_msg_event_contact, do:true"`
	EventId     interface{} // 事件唯一标识符(主键)
	Callsign    interface{} // 通信呼号标识(如"alpha-01")，默认空字符串
	Dsn         interface{} // 国防交换网络编号(defense switched network)，默认空字符串
	Email       interface{} // 电子邮箱地址(需符合name@domain格式)，默认空字符串
	Endpoint    interface{} // api终端地址(ip:port格式)，默认空字符串
	Freq        interface{} // 通信频率(单位：mhz，如"121.500")，默认空字符串
	Hostname    interface{} // 网络主机名(需dns可解析)，默认空字符串
	Iconsetpath interface{} // 图标资源存储路径(相对路径)，默认空字符串
	Modulation  interface{} // 无线电调制方式(am=调幅/fm=调频)
	Phone       interface{} // 联系电话号码(国际格式如+86-13912345678)，默认空字符串
	Version     interface{} // 联系人模式版本号(格式：v1.2.3)，默认空字符串
	UpdateBy    interface{} // 更新者id，默认0表示系统操作
	CreateBy    interface{} // 创建者id，默认0表示系统创建
	CreatedAt   interface{} // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt   interface{} // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted   interface{} // 软删标志，默认false表示有效数据
}
