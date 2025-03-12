// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMsgEventLinkDao is the data access object for the table t_msg_event_link.
type TMsgEventLinkDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns TMsgEventLinkColumns // columns contains all the column names of Table for convenient usage.
}

// TMsgEventLinkColumns defines and stores column names for the table t_msg_event_link.
type TMsgEventLinkColumns struct {
	EventId        string // 事件唯一标识符(主键)
	Mime           string // 资源mime类型(rfc 3023标准，如 application/xml/text/xml)，默认空字符串
	ParentCallsign string // 上级节点呼号标识(用于层级关系追踪)，默认空字符串
	ProductionTime string // 资源生成时间(unix时间戳，秒级)，默认0表示未记录
	Relation       string // 关联关系类型(cot标准：subject/object/indirect-object)，默认空字符串
	Remarks        string // 链接附加说明(自由文本，最大长度100字符)，默认空字符串
	Type           string // 链接资源分类(如影像/文档/传感器数据)，默认空字符串
	Uid            string // 关联资源全局唯一标识符(遵循cot uid规范)，默认空字符串
	Url            string // 资源访问地址(url格式，支持短链接)，默认空字符串
	Version        string // 链接模式版本标识(格式：主版本.次版本)，默认空字符串
	UpdateBy       string // 更新者id，默认0表示系统操作
	CreateBy       string // 创建者id，默认0表示系统创建
	CreatedAt      string // 创建时间(unix时间戳，秒级)，默认当前时间
	UpdatedAt      string // 更新时间(unix时间戳，秒级)，默认当前时间
	IsDeleted      string // 软删标志，默认false表示有效数据
}

// tMsgEventLinkColumns holds the columns for the table t_msg_event_link.
var tMsgEventLinkColumns = TMsgEventLinkColumns{
	EventId:        "event_id",
	Mime:           "mime",
	ParentCallsign: "parent_callsign",
	ProductionTime: "production_time",
	Relation:       "relation",
	Remarks:        "remarks",
	Type:           "type",
	Uid:            "uid",
	Url:            "url",
	Version:        "version",
	UpdateBy:       "update_by",
	CreateBy:       "create_by",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	IsDeleted:      "is_deleted",
}

// NewTMsgEventLinkDao creates and returns a new DAO object for table data access.
func NewTMsgEventLinkDao() *TMsgEventLinkDao {
	return &TMsgEventLinkDao{
		group:   "default",
		table:   "t_msg_event_link",
		columns: tMsgEventLinkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMsgEventLinkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMsgEventLinkDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMsgEventLinkDao) Columns() TMsgEventLinkColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMsgEventLinkDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMsgEventLinkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TMsgEventLinkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
