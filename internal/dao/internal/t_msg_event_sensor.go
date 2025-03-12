// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventSensorDao is the data access object for the table t_msg_event_sensor.
type TMsgEventSensorDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns TMsgEventSensorColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventSensorColumns defines and stores column names for the table t_msg_event_sensor.
type TMsgEventSensorColumns struct {
	EventId   string // 事件唯一标识符(主键)
	Elevation string // 传感器海拔高度(单位：米)，默认空字符串
	Vfov      string // 垂直视场角(vertical fov)，默认空字符串
	North     string // 北向校准参数，默认空字符串
	Roll      string // 传感器滚动角参数，默认空字符串
	Range     string // 传感器监测范围，默认空字符串
	Azimuth   string // 传感器方位角参数，默认空字符串
	Model     string // 传感器型号，默认空字符串
	Fov       string // 传感器视场角，默认空字符串
	Type      string // 传感器类型，默认空字符串
	Version   string // 传感器版本号，默认空字符串
	UpdateBy  string // 更新者id，默认0表示系统操作
	CreateBy  string // 创建者id，默认0表示系统创建
	CreatedAt string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted string // 软删标志，默认false表示有效数据
}

// tMsgEventSensorColumns holds the columns for the table t_msg_event_sensor.
var tMsgEventSensorColumns = TMsgEventSensorColumns{
	EventId:   "event_id",
	Elevation: "elevation",
	Vfov:      "vfov",
	North:     "north",
	Roll:      "roll",
	Range:     "range",
	Azimuth:   "azimuth",
	Model:     "model",
	Fov:       "fov",
	Type:      "type",
	Version:   "version",
	UpdateBy:  "update_by",
	CreateBy:  "create_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDeleted: "is_deleted",
}

// NewTMsgEventSensorDao creates and returns a new DAO object for table data access.
func NewTMsgEventSensorDao() *TMsgEventSensorDao {
	return &TMsgEventSensorDao{
		group:   "default",
		table:   "t_msg_event_sensor",
		columns: tMsgEventSensorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventSensorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventSensorDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventSensorDao) Columns() TMsgEventSensorColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventSensorDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventSensorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventSensorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
