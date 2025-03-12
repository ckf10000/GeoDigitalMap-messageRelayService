// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventTypeDao is the data access object for the table t_msg_event_type.
type TMsgEventTypeDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns TMsgEventTypeColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventTypeColumns defines and stores column names for the table t_msg_event_type.
type TMsgEventTypeColumns struct {
	Id          string // 自增主键，唯一标识记录
	EventType   string // 事件类型编码(长度限制50字符)，不可为空
	Alias       string // 事件类型别名(长度限制100字符)，默认空字符串
	Description string // 事件类型详细描述，默认空文本
	UpdateBy    string // 更新者id，默认0表示系统操作
	CreateBy    string // 创建者id，默认0表示系统创建
	CreatedAt   string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt   string // 更新时间(unix时间戳，秒级)，默认当前时间
	DeletedAt   string // 软删时间(unix时间戳，秒级)，默认0表示未删除
	IsDeleted   string // 软删标志，默认false表示有效数据
}

// tMsgEventTypeColumns holds the columns for the table t_msg_event_type.
var tMsgEventTypeColumns = TMsgEventTypeColumns{
	Id:          "id",
	EventType:   "event_type",
	Alias:       "alias",
	Description: "description",
	UpdateBy:    "update_by",
	CreateBy:    "create_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	IsDeleted:   "is_deleted",
}

// NewTMsgEventTypeDao creates and returns a new DAO object for table data access.
func NewTMsgEventTypeDao() *TMsgEventTypeDao {
	return &TMsgEventTypeDao{
		group:   "default",
		table:   "t_msg_event_type",
		columns: tMsgEventTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventTypeDao) Columns() TMsgEventTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
