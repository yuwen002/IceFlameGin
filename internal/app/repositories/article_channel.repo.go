package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/models/model"
)

// rArticleChannel
//
// @Description: 文章频道
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:20:54
type rArticleChannel struct {
	DB *gorm.DB
}

// NewArticleChannelRepository
//
// @Title NewArticleChannelRepository
// @Description: 创建一个新的 rArticleChannel 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:21:04
// @return *rArticleCategory
func NewArticleChannelRepository() *rArticleChannel {
	return &rArticleChannel{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 文章频道信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:24:01
// @receiver repo
// @param data
// @return error
func (repo *rArticleChannel) Insert(data *model.ArticleChannel) error {
	err := db.NewGormCore().Insert(data)
	if err != nil {
		return err
	}

	return nil
}
