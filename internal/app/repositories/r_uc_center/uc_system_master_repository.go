package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
)

//	UcSystemMasterRepository
//	@Description: 后台管理员
//
// @Author liuxingyu
// @Date 2024-02-08 23:31:58
type UcSystemMasterRepository struct {
	DB *gorm.DB
}

// NewUcSystemMasterRepository
//
// @Title NewUcSystemMasterRepository
// @Description: 创建一个新的 UcSystemMasterRepository 仓库实例
// @Author liuxingyu
// @Date 2024-02-08 23:36:52
// @return *UcSystemMasterRepository 创建一个新的 UcSystemMasterRepository 仓库实例
func NewUcSystemMasterRepository() *UcSystemMasterRepository {
	return &UcSystemMasterRepository{DB: db.DB}
}
