package models

import (
	"ice_flame_gin/internal/pkg/utils"
)

// UcAccount
// @Description: 用户中心主表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 15:33:18
type UcAccount struct {
	ID             uint    `gorm:"primaryKey"`
	IdentityCardID *string `gorm:"index;type:char(18)"`
	Username       string  `gorm:"type:char(40)"`
	PasswordHash   string  `gorm:"type:char(60)"`
	Tel            string  `gorm:"type:char(20)"`
	Status         int8    `gorm:"default:0"`
	RealNameType   int8    `gorm:"default:1"`
	CreatedAt      utils.CustomTime
	UpdatedAt      utils.CustomTime
}

// TableName
//
// @Title TableName
// @Description 关联表名
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 15:33:36
// @receiver UcAccount
// @return string
func (UcAccount) TableName() string {
	return "uc_account"
}
