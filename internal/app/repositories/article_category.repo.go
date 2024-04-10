package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/models/model"
)

// rArticleCategory
//
// @Description: 文章分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-09 23:51:08
type rArticleCategory struct {
	DB *gorm.DB
}

// NewArticleCategoryRepository
//
// @Title NewArticleCategoryRepository
// @Description:  创建一个新的 rArticleCategory 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-09 23:51:43
// @return *rArticleCategory
func NewArticleCategoryRepository() *rArticleCategory {
	return &rArticleCategory{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 文章分类信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-09 23:54:22
// @receiver repo
// @param in
// @return error
func (repo *rArticleCategory) Insert(data *model.ArticleCategory) error {
	err := db.NewGormCore().Insert(data)

	if err != nil {
		return err
	}

	return nil
}

// GetByID
//
// @Title GetByID
// @Description:  按ID获取文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-09 23:54:47
// @receiver repo
// @param id
// @return out
// @return err
func (repo *rArticleCategory) GetByID(id uint32) (out *model.ArticleCategory, err error) {
	err = db.NewGormCore().GetByID(id, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (repo *rArticleCategory) GetByFID(fid uint32) (out *model.ArticleCategory, err error) {
	err = db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Field:     "id, fid, name",
		Condition: "status = 0 and fid = ?",
		Args:      []interface{}{fid},
		Order:     "sort desc, id desc",
		PageType:  2,
		Limit: db.Limit{
			Length: 1000,
		},
	}, out)
}
