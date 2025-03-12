// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TIdcFileResourceDao is the data access object for the table t_idc_file_resource.
type TIdcFileResourceDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns TIdcFileResourceColumns // columns contains all the column names of Table for convenient usage.
}

// TIdcFileResourceColumns defines and stores column names for the table t_idc_file_resource.
type TIdcFileResourceColumns struct {
	Id         string // 文件id
	ResourceId string // 所属业务模块的资源ID
	FileName   string // 文件名
	TypeId     string // 文件资源类型id
	CategoryId string // 文件资源类别id
	OssPath    string // 存储对象路径
	Status     string // 文件状态：0.失败，1.成功，2.处理中，3.待上传
	Size       string // 文件大小
	Checksum   string // 文件md5值
	Reason     string // 失败原因
	UpdateBy   string // 更新者id
	CreateBy   string // 创建者id
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	IsDeleted  string // 是否已删除
}

// tIdcFileResourceColumns holds the columns for the table t_idc_file_resource.
var tIdcFileResourceColumns = TIdcFileResourceColumns{
	Id:         "id",
	ResourceId: "resource_id",
	FileName:   "file_name",
	TypeId:     "type_id",
	CategoryId: "category_id",
	OssPath:    "oss_path",
	Status:     "status",
	Size:       "size",
	Checksum:   "checksum",
	Reason:     "reason",
	UpdateBy:   "update_by",
	CreateBy:   "create_by",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	IsDeleted:  "is_deleted",
}

// NewTIdcFileResourceDao creates and returns a new DAO object for table data access.
func NewTIdcFileResourceDao() *TIdcFileResourceDao {
	return &TIdcFileResourceDao{
		group:   "default",
		table:   "t_idc_file_resource",
		columns: tIdcFileResourceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TIdcFileResourceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TIdcFileResourceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TIdcFileResourceDao) Columns() TIdcFileResourceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TIdcFileResourceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TIdcFileResourceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TIdcFileResourceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
