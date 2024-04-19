package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
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

// GetList
//
// @Title GetList
// @Description: 文章频道信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:15:30
// @receiver repo
// @param data
// @return out
// @return err
func (repo *rArticleChannel) GetList(data dto.ListArticleChannelInput) (out []*model.ArticleChannel, err error) {
	db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Order:       "sort desc, id desc",
		Preload:     nil,
		PreloadFunc: nil,
		PageType:    2,
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
// @Description: 文章频道信息列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:17:11
// @receiver repo
// @return int64
// @return error
func (repo *rArticleChannel) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameArticleChannel).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}
