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
// @Description: 访问类型分类写入
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

// GetList
//
// @Title GetList
// @Description: 访问类型分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 21:54:24
// @receiver repo
// @param data
// @return []*model.UcSystemMasterVisitCategory
// @return error
func (repo *rUcSystemMasterVisitorCategory) GetList(data dto.ListSystemMasterVisitorCategoryInput) ([]*model.UcSystemMasterVisitCategory, error) {
	var systemMasterVisitCategory []*model.UcSystemMasterVisitCategory
	err := db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Order:    data.Order,
		PageType: 2,
		Limit: db.Limit{
			Length: data.Length,
			Offset: data.Start,
		},
	}, &systemMasterVisitCategory)

	if err != nil {
		return nil, err
	}

	return systemMasterVisitCategory, nil
}

// CountRecords
//
// @Title CountRecords
// @Description: 访问类型分类列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 21:55:07
// @receiver repo
// @return int64
// @return error
func (repo *rUcSystemMasterVisitorCategory) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameUcSystemMasterVisitCategory).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}
