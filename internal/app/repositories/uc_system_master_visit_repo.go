package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
)

// rUcSystemMasterVisitCategory
// @Description: 访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:11:38
type rUcSystemMasterVisitCategory struct {
	DBs map[string]*gorm.DB
	DB  *gorm.DB
}

// NewUcSystemMasterVisit
//
// @Title NewUcSystemMasterVisit
// @Description: 创建一个新的 rUcSystemMasterVisit 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:11:56
// @return *rUcSystemMasterVisit
func NewUcSystemMasterVisit() *rUcSystemMasterVisitCategory {
	return &rUcSystemMasterVisitCategory{
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
func (repo *rUcSystemMasterVisitCategory) Insert(data dto.SystemMasterVisitCategoryInput) error {
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
func (repo *rUcSystemMasterVisitCategory) GetList(data dto.ListSystemMasterVisitCategoryInput) ([]*model.UcSystemMasterVisitCategory, error) {
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
func (repo *rUcSystemMasterVisitCategory) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameUcSystemMasterVisitCategory).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}

// GetById
//
// @Title GetById
// @Description: 按ID获取访问类型分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 14:32:46
// @receiver repo
// @param id
// @return out
// @return err
func (repo *rUcSystemMasterVisitCategory) GetById(id uint32) (out *model.UcSystemMasterVisitCategory, err error) {
	err = db.NewGormCore().GetByID(id, &out)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, nil
	}

	return out, nil
}

// UpdateByID
//
// @Title UpdateByID
// @Description: 按ID更新访问类型分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 15:05:48
// @receiver repo
// @param id
// @param in
// @return error
func (repo *rUcSystemMasterVisitCategory) UpdateByID(id uint32, in *model.UcSystemMasterVisitCategory) error {
	err := db.NewGormCore().UpdateByID(id, in)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
//
// @Title GetAll
// @Description: 获取全部访问类型分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 15:06:26
// @receiver repo
// @return out
// @return err
func (repo *rUcSystemMasterVisitCategory) GetAll() (out []*model.UcSystemMasterVisitCategory, err error) {
	err = db.NewGormCore().GetAll(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
