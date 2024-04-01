package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
)

// rSinglePage
//
// @Description: 单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 23:14:31
type rSinglePage struct {
	DB *gorm.DB
}

// NewSinglePageRepository
//
// @Title NewSinglePageRepository
// @Description: 创建一个新的 rSinglePage 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 23:15:19
// @return *rSinglePage
func NewSinglePageRepository() *rSinglePage {
	return &rSinglePage{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 单页信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 23:15:35
// @receiver r
// @param data
// @return error
func (repo *rSinglePage) Insert(data dto.SinglePageInput) error {
	err := db.NewGormCore().Insert(&model.SinglePage{
		Title:       data.Title,
		Description: data.Description,
		Keyword:     data.Keyword,
		Content:     data.Content,
		Thumbnail:   data.Thumbnail,
		Click:       data.Click,
		Status:      data.Status,
	})

	if err != nil {
		return err
	}

	return nil
}

// GetList
//
// @Title GetList
// @Description:  单页信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-02 00:19:12
// @receiver repo
// @param data
// @return []*model.SinglePage
// @return error
func (repo *rSinglePage) GetList(data dto.ListSinglePageInput) ([]*model.SinglePage, error) {
	var out []*model.SinglePage
	err := db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Field:    "title, thumbnail, click, status, created_at, updated_at",
		Order:    data.Order,
		PageType: 2,
		Limit: db.Limit{
			Length: data.Length,
			Offset: data.Start,
		},
	}, &out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

// CountRecords
//
// @Title CountRecords
// @Description: 单页信息列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-02 00:28:29
// @receiver repo
// @return int64
// @return error
func (repo *rSinglePage) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameSinglePage).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}
