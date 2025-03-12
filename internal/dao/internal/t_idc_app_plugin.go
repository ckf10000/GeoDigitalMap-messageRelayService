// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TIdcAppPluginDao is the data access object for the table t_idc_app_plugin.
type TIdcAppPluginDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns TIdcAppPluginColumns // columns contains all the column names of Table for convenient usage.
}

// TIdcAppPluginColumns defines and stores column names for the table t_idc_app_plugin.
type TIdcAppPluginColumns struct {
	Id            string // 主键id
	PluginId      string // 插件id
	PluginName    string // 插件名
	TypeId        string // 文件资源类型id
	CategoryId    string // 文件资源类别id
	Version       string // 插件版本
	Description   string // 插件描述
	ReleaseStatus string // 插件发布状态：0.待发布，1.已发布
	UpdateBy      string // 更新者id
	CreateBy      string // 创建者id
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	IsDeleted     string // 是否已删除
}

// tIdcAppPluginColumns holds the columns for the table t_idc_app_plugin.
var tIdcAppPluginColumns = TIdcAppPluginColumns{
	Id:            "id",
	PluginId:      "plugin_id",
	PluginName:    "plugin_name",
	TypeId:        "type_id",
	CategoryId:    "category_id",
	Version:       "version",
	Description:   "description",
	ReleaseStatus: "release_status",
	UpdateBy:      "update_by",
	CreateBy:      "create_by",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	IsDeleted:     "is_deleted",
}

// NewTIdcAppPluginDao creates and returns a new DAO object for table data access.
func NewTIdcAppPluginDao() *TIdcAppPluginDao {
	return &TIdcAppPluginDao{
		group:   "default",
		table:   "t_idc_app_plugin",
		columns: tIdcAppPluginColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TIdcAppPluginDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TIdcAppPluginDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TIdcAppPluginDao) Columns() TIdcAppPluginColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TIdcAppPluginDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TIdcAppPluginDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TIdcAppPluginDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
