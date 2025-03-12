// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TIdcFileResourceCategoryDao is the data access object for the table t_idc_file_resource_category.
type TIdcFileResourceCategoryDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of the current DAO.
	columns TIdcFileResourceCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// TIdcFileResourceCategoryColumns defines and stores column names for the table t_idc_file_resource_category.
type TIdcFileResourceCategoryColumns struct {
	Id              string // 文件资源类别id
	Category        string // 文件资源类别
	Description     string // 类别描述
	OssType         string // 对象存储类型
	EndpointUrl     string // 对象存储端点url
	AccessKey       string // 访问对象存储的key
	SecretAccessKey string // 访问对象存储的secret key
	RegionName      string // 区域
	BucketName      string // 桶名
	UpdateBy        string // 更新者id
	CreateBy        string // 创建者id
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
	DeletedAt       string // 删除时间
	IsDeleted       string // 是否已删除
}

// tIdcFileResourceCategoryColumns holds the columns for the table t_idc_file_resource_category.
var tIdcFileResourceCategoryColumns = TIdcFileResourceCategoryColumns{
	Id:              "id",
	Category:        "category",
	Description:     "description",
	OssType:         "oss_type",
	EndpointUrl:     "endpoint_url",
	AccessKey:       "access_key",
	SecretAccessKey: "secret_access_key",
	RegionName:      "region_name",
	BucketName:      "bucket_name",
	UpdateBy:        "update_by",
	CreateBy:        "create_by",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
	IsDeleted:       "is_deleted",
}

// NewTIdcFileResourceCategoryDao creates and returns a new DAO object for table data access.
func NewTIdcFileResourceCategoryDao() *TIdcFileResourceCategoryDao {
	return &TIdcFileResourceCategoryDao{
		group:   "default",
		table:   "t_idc_file_resource_category",
		columns: tIdcFileResourceCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TIdcFileResourceCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TIdcFileResourceCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TIdcFileResourceCategoryDao) Columns() TIdcFileResourceCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TIdcFileResourceCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TIdcFileResourceCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TIdcFileResourceCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
