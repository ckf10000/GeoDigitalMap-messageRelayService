// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventTrackDao is the data access object for the table t_msg_event_track.
type TMsgEventTrackDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns TMsgEventTrackColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventTrackColumns defines and stores column names for the table t_msg_event_track.
type TMsgEventTrackColumns struct {
	EventId   string // 事件唯一标识符(主键)
	Course    string // 移动方向角度(单位：度，长度限制100字符)，默认空字符串
	Ecourse   string // 方向角测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Eslope    string // 坡度测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Espeed    string // 速度测量误差(1σ标准差，长度限制100字符)，默认空字符串
	Speed     string // 移动速度值(单位：米/秒，长度限制100字符)，默认空字符串
	Version   string //
	UpdateBy  string // 更新者id，默认0表示系统操作
	CreateBy  string // 创建者id，默认0表示系统创建
	CreatedAt string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted string // 软删标志，默认false表示有效数据
}

// tMsgEventTrackColumns holds the columns for the table t_msg_event_track.
var tMsgEventTrackColumns = TMsgEventTrackColumns{
	EventId:   "event_id",
	Course:    "course",
	Ecourse:   "ecourse",
	Eslope:    "eslope",
	Espeed:    "espeed",
	Speed:     "speed",
	Version:   "version",
	UpdateBy:  "update_by",
	CreateBy:  "create_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDeleted: "is_deleted",
}

// NewTMsgEventTrackDao creates and returns a new DAO object for table data access.
func NewTMsgEventTrackDao() *TMsgEventTrackDao {
	return &TMsgEventTrackDao{
		group:   "default",
		table:   "t_msg_event_track",
		columns: tMsgEventTrackColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventTrackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventTrackDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventTrackDao) Columns() TMsgEventTrackColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventTrackDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventTrackDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventTrackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
