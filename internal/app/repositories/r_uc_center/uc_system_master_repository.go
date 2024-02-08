package repositories

import "gorm.io/gorm"

//	UcSystemMasterRepository
//	@Description: 后台管理员
//
// @Author liuxingyu
// @Date 2024-02-08 23:31:58
type UcSystemMasterRepository struct {
	DB *gorm.DB
}
