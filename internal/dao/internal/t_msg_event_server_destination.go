// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventServerDestinationDao is the data access object for the table t_msg_event_server_destination.
type TMsgEventServerDestinationDao struct {
	table   string                            // table is the underlying table name of the DAO.
	group   string                            // group is the database configuration group name of the current DAO.
	columns TMsgEventServerDestinationColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventServerDestinationColumns defines and stores column names for the table t_msg_event_server_destination.
type TMsgEventServerDestinationColumns struct {
	EventId      string // 事件唯一标识符(主键)
	Destinations string // 目标地址组合(格式：ip:端口:协议:设备id)，不可为空
	UpdateBy     string // 更新者id，默认0表示系统操作
	CreateBy     string // 创建者id，默认0表示系统创建
	CreatedAt    string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt    string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted    string // 软删标志，默认false表示有效数据
}

// tMsgEventServerDestinationColumns holds the columns for the table t_msg_event_server_destination.
var tMsgEventServerDestinationColumns = TMsgEventServerDestinationColumns{
	EventId:      "event_id",
	Destinations: "destinations",
	UpdateBy:     "update_by",
	CreateBy:     "create_by",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	IsDeleted:    "is_deleted",
}

// NewTMsgEventServerDestinationDao creates and returns a new DAO object for table data access.
func NewTMsgEventServerDestinationDao() *TMsgEventServerDestinationDao {
	return &TMsgEventServerDestinationDao{
		group:   "default",
		table:   "t_msg_event_server_destination",
		columns: tMsgEventServerDestinationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventServerDestinationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventServerDestinationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventServerDestinationDao) Columns() TMsgEventServerDestinationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventServerDestinationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventServerDestinationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventServerDestinationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
