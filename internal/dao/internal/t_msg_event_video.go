// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventVideoDao is the data access object for the table t_msg_event_video.
type TMsgEventVideoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns TMsgEventVideoColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventVideoColumns defines and stores column names for the table t_msg_event_video.
type TMsgEventVideoColumns struct {
	EventId   string // 事件唯一标识符(主键)
	Sensor    string // 视频传感器标识符(长度限制100字符)，默认空字符串
	Spi       string // 安全参数索引(security parameters index，长度限制100字符)，默认空字符串
	Url       string // 视频资源访问地址(长度限制100字符)，默认空字符串
	UpdateBy  string // 更新者id，默认0表示系统操作
	CreateBy  string // 创建者id，默认0表示系统创建
	CreatedAt string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted string // 软删标志，默认false表示有效数据
}

// tMsgEventVideoColumns holds the columns for the table t_msg_event_video.
var tMsgEventVideoColumns = TMsgEventVideoColumns{
	EventId:   "event_id",
	Sensor:    "sensor",
	Spi:       "spi",
	Url:       "url",
	UpdateBy:  "update_by",
	CreateBy:  "create_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDeleted: "is_deleted",
}

// NewTMsgEventVideoDao creates and returns a new DAO object for table data access.
func NewTMsgEventVideoDao() *TMsgEventVideoDao {
	return &TMsgEventVideoDao{
		group:   "default",
		table:   "t_msg_event_video",
		columns: tMsgEventVideoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventVideoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventVideoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventVideoDao) Columns() TMsgEventVideoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventVideoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventVideoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventVideoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
