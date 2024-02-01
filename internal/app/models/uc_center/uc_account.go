package uc_center

import (
	"time"
)

// UcAccount
// @Description: 中户中心主表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 15:33:18
type UcAccount struct {
	ID             uint   `gorm:"primaryKey"`
	IdentityCardID uint   `gorm:"index"`
	Username       string `gorm:"size:40"`
	PasswordHash   string `gorm:"size:60"`
	Tel            string `gorm:"size:20"`
	Status         int8   `gorm:"default:0"`
	RealNameType   int8   `gorm:"default:1"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
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