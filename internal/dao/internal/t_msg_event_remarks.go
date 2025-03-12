// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventRemarksDao is the data access object for the table t_msg_event_remarks.
type TMsgEventRemarksDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns TMsgEventRemarksColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventRemarksColumns defines and stores column names for the table t_msg_event_remarks.
type TMsgEventRemarksColumns struct {
	EventId   string // 事件唯一标识符(主键)
	Keywords  string // 逗号分隔的备注关键词(如"任务a,紧急")，默认空字符串
	Source    string // 备注来源标识符(发送方uid)，默认空字符串
	Time      string // 备注生成时间(unix时间戳，秒级)，默认0表示未记录
	ToField   string // 目标接收方uid，默认空字符串表示广播
	Version   string // 备注格式版本标识，默认空字符串
	Intag     string // 扩展标签信息(最大长度1000字符)，默认空字符串
	UpdateBy  string // 更新者id，默认0表示系统操作
	CreateBy  string // 创建者id，默认0表示系统创建
	CreatedAt string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted string // 软删标志，默认false表示有效数据
}

// tMsgEventRemarksColumns holds the columns for the table t_msg_event_remarks.
var tMsgEventRemarksColumns = TMsgEventRemarksColumns{
	EventId:   "event_id",
	Keywords:  "keywords",
	Source:    "source",
	Time:      "time",
	ToField:   "to_field",
	Version:   "version",
	Intag:     "intag",
	UpdateBy:  "update_by",
	CreateBy:  "create_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDeleted: "is_deleted",
}

// NewTMsgEventRemarksDao creates and returns a new DAO object for table data access.
func NewTMsgEventRemarksDao() *TMsgEventRemarksDao {
	return &TMsgEventRemarksDao{
		group:   "default",
		table:   "t_msg_event_remarks",
		columns: tMsgEventRemarksColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventRemarksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventRemarksDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventRemarksDao) Columns() TMsgEventRemarksColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventRemarksDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventRemarksDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventRemarksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
