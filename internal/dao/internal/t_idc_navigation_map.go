// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TIdcNavigationMapDao is the data access object for the table t_idc_navigation_map.
type TIdcNavigationMapDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of the current DAO.
	columns TIdcNavigationMapColumns // columns contains all the column names of Table for convenient usage.
}

// TIdcNavigationMapColumns defines and stores column names for the table t_idc_navigation_map.
type TIdcNavigationMapColumns struct {
	Id            string // 主键id
	PackageId     string // 导航地图包id
	PackageName   string // 导航地图包名
	TypeId        string // 文件资源类型id
	CategoryId    string // 文件资源类别id
	Version       string // 导航地图包版本
	Description   string // 导航地图包描述
	ReleaseStatus string // 导航地图包发布状态：0.待发布，1.已发布
	UpdateBy      string // 更新者id
	CreateBy      string // 创建者id
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	IsDeleted     string // 是否已删除
}

// tIdcNavigationMapColumns holds the columns for the table t_idc_navigation_map.
var tIdcNavigationMapColumns = TIdcNavigationMapColumns{
	Id:            "id",
	PackageId:     "package_id",
	PackageName:   "package_name",
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

// NewTIdcNavigationMapDao creates and returns a new DAO object for table data access.
func NewTIdcNavigationMapDao() *TIdcNavigationMapDao {
	return &TIdcNavigationMapDao{
		group:   "default",
		table:   "t_idc_navigation_map",
		columns: tIdcNavigationMapColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TIdcNavigationMapDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TIdcNavigationMapDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TIdcNavigationMapDao) Columns() TIdcNavigationMapColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TIdcNavigationMapDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TIdcNavigationMapDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TIdcNavigationMapDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
