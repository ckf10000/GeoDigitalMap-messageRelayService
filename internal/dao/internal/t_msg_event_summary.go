// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventSummaryDao is the data access object for the table t_msg_event_summary.
type TMsgEventSummaryDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns TMsgEventSummaryColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventSummaryColumns defines and stores column names for the table t_msg_event_summary.
type TMsgEventSummaryColumns struct {
	EventId   string // 事件唯一标识符(主键)
	How       string // 事件触发方式标识(长度限制100字符)，默认空字符串
	Time      string // 事件发生时间(unix时间戳，秒级)，默认0表示未记录
	Start     string // 事件有效开始时间(unix时间戳，秒级)，默认0表示立即生效
	Stale     string // 事件失效时间(unix时间戳，秒级)，默认0表示永不过期
	Type      string // 事件分类类型(长度限制100字符)，默认空字符串
	Version   string // 事件版本标识符(长度限制100字符)，默认空字符串
	UpdateBy  string // 更新者id，默认0表示系统操作
	CreateBy  string // 创建者id，默认0表示系统创建
	CreatedAt string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted string // 软删标志，默认false表示有效数据
}

// tMsgEventSummaryColumns holds the columns for the table t_msg_event_summary.
var tMsgEventSummaryColumns = TMsgEventSummaryColumns{
	EventId:   "event_id",
	How:       "how",
	Time:      "time",
	Start:     "start",
	Stale:     "stale",
	Type:      "type",
	Version:   "version",
	UpdateBy:  "update_by",
	CreateBy:  "create_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDeleted: "is_deleted",
}

// NewTMsgEventSummaryDao creates and returns a new DAO object for table data access.
func NewTMsgEventSummaryDao() *TMsgEventSummaryDao {
	return &TMsgEventSummaryDao{
		group:   "default",
		table:   "t_msg_event_summary",
		columns: tMsgEventSummaryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventSummaryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventSummaryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventSummaryDao) Columns() TMsgEventSummaryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventSummaryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventSummaryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventSummaryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
