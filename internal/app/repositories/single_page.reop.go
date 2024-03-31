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
func (r *rSinglePage) Insert(data dto.SinglePageInput) error {
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
