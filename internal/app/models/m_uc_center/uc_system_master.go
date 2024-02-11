package models

import (
	"ice_flame_gin/internal/pkg/utils"
)

//	UcSystemMaster
//	@Description: 后台管理用户
//
// @Author liuxingyu
// @Date 2024-02-11 21:56:55
type UcSystemMaster struct {
	ID          uint `gorm:"primaryKey"`
	AccountID   uint
	Name        string
	Tel         string
	SuperMaster uint `gorm:"default:0"`
	CreatedAt   utils.CustomTime
	UpdatedAt   utils.CustomTime
}

// TableName
//
// @Title TableName
// @Description: 关联表名
// @Author liuxingyu
// @Date 2024-02-11 21:57:14
// @receiver UcSystemMaster
// @return string
func (UcSystemMaster) TableName() string {
	return "uc_system_master"
}
