// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUcSystemMaster = "uc_system_master"

// UcSystemMaster 管理员用户表
type UcSystemMaster struct {
	ID          uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true;comment:自增ID" json:"id"`                       // 自增ID
	AccountID   uint32    `gorm:"column:account_id;type:int(10) unsigned;index:account_id,priority:1;comment:关联account id" json:"account_id"` // 关联account id
	Name        string    `gorm:"column:name;type:varchar(32);comment:用户姓名" json:"name"`                                                      // 用户姓名
	Email       string    `gorm:"column:email;type:varchar(60);comment:用户邮箱" json:"email"`                                                    // 用户邮箱
	Tel         string    `gorm:"column:tel;type:varchar(15);comment:用户电话" json:"tel"`                                                        // 用户电话
	SuperMaster bool      `gorm:"column:super_master;type:tinyint(1);comment:超级管理员(=1为超级管理员)" json:"super_master"`                            // 超级管理员(=1为超级管理员)
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

// TableName UcSystemMaster's table name
func (*UcSystemMaster) TableName() string {
	return TableNameUcSystemMaster
}
