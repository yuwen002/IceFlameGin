package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
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

// NewArticleTagRepository
//
// @Title NewArticleTag
// @Description: 创建一个新的 rArticleTag 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-22 23:45:27
// @return *rArticleTag
func NewArticleTagRepository() *rArticleTag {
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

// GetList
//
// @Title GetList
// @Description: 文章标签列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-30 23:49:48
// @receiver repo
// @param data
// @return out
// @return err
func (repo *rArticleTag) GetList(data dto.ListArticleTagInput) (out []*model.ArticleTag, err error) {
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
// @Description: 文章标签信息列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-30 23:50:14
// @receiver repo
// @return int64
// @return error
func (repo *rArticleTag) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameArticleTag).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}
