// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventMcvDao is the data access object for the table t_msg_event_mcv.
type TMsgEventMcvDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns TMsgEventMcvColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventMcvColumns defines and stores column names for the table t_msg_event_mcv.
type TMsgEventMcvColumns struct {
	EventId   string // 事件唯一标识符(主键)
	Device    string // 物理设备型号(长度限制100字符)，默认空字符串
	Os        string // 设备操作系统类型(长度限制100字符)，默认空字符串
	Platform  string // 消息事件客户端平台变体名称(长度限制100字符)，默认空字符串
	Version   string // 消息事件客户端软件版本号(长度限制100字符)，默认空字符串
	UpdateBy  string // 更新者id，默认0表示系统操作
	CreateBy  string // 创建者id，默认0表示系统创建
	CreatedAt string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted string // 软删标志，默认false表示有效数据
}

// tMsgEventMcvColumns holds the columns for the table t_msg_event_mcv.
var tMsgEventMcvColumns = TMsgEventMcvColumns{
	EventId:   "event_id",
	Device:    "device",
	Os:        "os",
	Platform:  "platform",
	Version:   "version",
	UpdateBy:  "update_by",
	CreateBy:  "create_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDeleted: "is_deleted",
}

// NewTMsgEventMcvDao creates and returns a new DAO object for table data access.
func NewTMsgEventMcvDao() *TMsgEventMcvDao {
	return &TMsgEventMcvDao{
		group:   "default",
		table:   "t_msg_event_mcv",
		columns: tMsgEventMcvColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventMcvDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventMcvDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventMcvDao) Columns() TMsgEventMcvColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventMcvDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventMcvDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventMcvDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
