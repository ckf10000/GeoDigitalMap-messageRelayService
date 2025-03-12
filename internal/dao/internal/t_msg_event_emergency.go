// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventEmergencyDao is the data access object for the table t_msg_event_emergency.
type TMsgEventEmergencyDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns TMsgEventEmergencyColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventEmergencyColumns defines and stores column names for the table t_msg_event_emergency.
type TMsgEventEmergencyColumns struct {
	EventId   string // 事件唯一标识符(主键)
	Alert     string // 应急警报等级或类型(如"critical"/"warning")，默认空字符串
	Cancel    string // 应急信标取消状态(true=已取消，false=有效)，默认false
	Type      string // 应急事件分类类型(如医疗/消防/警务)，默认"constructor"
	UpdateBy  string // 更新者id，默认0表示系统操作
	CreateBy  string // 创建者id，默认0表示系统创建
	CreatedAt string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted string // 软删标志，默认false表示有效数据
}

// tMsgEventEmergencyColumns holds the columns for the table t_msg_event_emergency.
var tMsgEventEmergencyColumns = TMsgEventEmergencyColumns{
	EventId:   "event_id",
	Alert:     "alert",
	Cancel:    "cancel",
	Type:      "type",
	UpdateBy:  "update_by",
	CreateBy:  "create_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDeleted: "is_deleted",
}

// NewTMsgEventEmergencyDao creates and returns a new DAO object for table data access.
func NewTMsgEventEmergencyDao() *TMsgEventEmergencyDao {
	return &TMsgEventEmergencyDao{
		group:   "default",
		table:   "t_msg_event_emergency",
		columns: tMsgEventEmergencyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventEmergencyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventEmergencyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventEmergencyDao) Columns() TMsgEventEmergencyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventEmergencyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventEmergencyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventEmergencyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
