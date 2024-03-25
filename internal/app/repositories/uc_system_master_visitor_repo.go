package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
)

// rUcSystemMasterVisitor
// @Description: 访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:11:38
type rUcSystemMasterVisitorCategory struct {
	DBs map[string]*gorm.DB
	DB  *gorm.DB
}

// NewUcSystemMasterVisitor
//
// @Title NewUcSystemMasterVisitor
// @Description: 创建一个新的 rUcSystemMasterVisitor 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:11:56
// @return *rUcSystemMasterVisitor
func NewUcSystemMasterVisitor() *rUcSystemMasterVisitorCategory {
	return &rUcSystemMasterVisitorCategory{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 访问类型写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:18:05
// @receiver repo
// @param data
// @return error
func (repo *rUcSystemMasterVisitorCategory) Insert(data dto.SystemMasterVisitorCategoryInput) error {
	err := db.NewGormCore().Insert(&model.UcSystemMasterVisitCategory{
		Title: data.Title,
	})

	if err != nil {
		return err
	}

	return nil
}
