// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventContactDao is the data access object for the table t_msg_event_contact.
type TMsgEventContactDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns TMsgEventContactColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventContactColumns defines and stores column names for the table t_msg_event_contact.
type TMsgEventContactColumns struct {
	EventId     string // 事件唯一标识符(主键)
	Callsign    string // 通信呼号标识(如"alpha-01")，默认空字符串
	Dsn         string // 国防交换网络编号(defense switched network)，默认空字符串
	Email       string // 电子邮箱地址(需符合name@domain格式)，默认空字符串
	Endpoint    string // api终端地址(ip:port格式)，默认空字符串
	Freq        string // 通信频率(单位：mhz，如"121.500")，默认空字符串
	Hostname    string // 网络主机名(需dns可解析)，默认空字符串
	Iconsetpath string // 图标资源存储路径(相对路径)，默认空字符串
	Modulation  string // 无线电调制方式(am=调幅/fm=调频)
	Phone       string // 联系电话号码(国际格式如+86-13912345678)，默认空字符串
	Version     string // 联系人模式版本号(格式：v1.2.3)，默认空字符串
	UpdateBy    string // 更新者id，默认0表示系统操作
	CreateBy    string // 创建者id，默认0表示系统创建
	CreatedAt   string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt   string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted   string // 软删标志，默认false表示有效数据
}

// tMsgEventContactColumns holds the columns for the table t_msg_event_contact.
var tMsgEventContactColumns = TMsgEventContactColumns{
	EventId:     "event_id",
	Callsign:    "callsign",
	Dsn:         "dsn",
	Email:       "email",
	Endpoint:    "endpoint",
	Freq:        "freq",
	Hostname:    "hostname",
	Iconsetpath: "iconsetpath",
	Modulation:  "modulation",
	Phone:       "phone",
	Version:     "version",
	UpdateBy:    "update_by",
	CreateBy:    "create_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	IsDeleted:   "is_deleted",
}

// NewTMsgEventContactDao creates and returns a new DAO object for table data access.
func NewTMsgEventContactDao() *TMsgEventContactDao {
	return &TMsgEventContactDao{
		group:   "default",
		table:   "t_msg_event_contact",
		columns: tMsgEventContactColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventContactDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventContactDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventContactDao) Columns() TMsgEventContactColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventContactDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventContactDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventContactDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
