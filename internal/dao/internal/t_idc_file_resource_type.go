// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TIdcFileResourceTypeDao is the data access object for the table t_idc_file_resource_type.
type TIdcFileResourceTypeDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns TIdcFileResourceTypeColumns // columns contains all the column names of Table for convenient usage.
}

// TIdcFileResourceTypeColumns defines and stores column names for the table t_idc_file_resource_type.
type TIdcFileResourceTypeColumns struct {
	Id          string // 文件资源类型id
	FileType    string // 文件资源类型
	Alias       string // 文件资源类型别名
	Description string // 类型描述
	UpdateBy    string // 更新者id
	CreateBy    string // 创建者id
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
	IsDeleted   string // 是否已删除
}

// tIdcFileResourceTypeColumns holds the columns for the table t_idc_file_resource_type.
var tIdcFileResourceTypeColumns = TIdcFileResourceTypeColumns{
	Id:          "id",
	FileType:    "file_type",
	Alias:       "alias",
	Description: "description",
	UpdateBy:    "update_by",
	CreateBy:    "create_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	IsDeleted:   "is_deleted",
}

// NewTIdcFileResourceTypeDao creates and returns a new DAO object for table data access.
func NewTIdcFileResourceTypeDao() *TIdcFileResourceTypeDao {
	return &TIdcFileResourceTypeDao{
		group:   "default",
		table:   "t_idc_file_resource_type",
		columns: tIdcFileResourceTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TIdcFileResourceTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TIdcFileResourceTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TIdcFileResourceTypeDao) Columns() TIdcFileResourceTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TIdcFileResourceTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TIdcFileResourceTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TIdcFileResourceTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
