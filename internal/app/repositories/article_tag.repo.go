package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/models/model"
)

// rArticleTag
//
// @Description: 文章tag
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-22 23:45:04
type rArticleTag struct {
	DB *gorm.DB
}

// NewArticleTag
//
// @Title NewArticleTag
// @Description: 创建一个新的 rArticleTag 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-22 23:45:27
// @return *rArticleTag
func NewArticleTag() *rArticleTag {
	return &rArticleTag{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 文章tag信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-22 23:48:12
// @receiver repo
// @param data
// @return error
func (repo *rArticleTag) Insert(data *model.ArticleTag) error {
	err := db.NewGormCore().Insert(data)
	if err != nil {
		return err
	}

	return nil
}
