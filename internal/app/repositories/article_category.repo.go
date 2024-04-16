package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
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

// GetListByFID
//
// @Title GetListByFID
// @Description: 按FID获取文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-11 23:55:53
// @receiver repo
// @param fid
// @return out
// @return err
func (repo *rArticleCategory) GetListByFID(fid uint32) (out []*model.ArticleCategory, err error) {
	err = db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Field:     "id, fid, name",
		Condition: "status = 1 and fid = ?",
		Args:      []interface{}{fid},
		Order:     "sort desc, id desc",
	}, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// GetList
//
// @Title GetList
// @Description: 文章分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 00:41:43
// @receiver repo
// @param data
// @return out
// @return err
func (repo *rArticleCategory) GetList(data dto.ListArticleCategoryInput) (out []*model.ArticleCategory, err error) {
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
// @Description: 文章分类信息列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 00:44:20
// @receiver repo
// @return int64
// @return error
func (repo *rArticleCategory) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameArticleCategory).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}

// UpdateByID
//
// @Title UpdateByID
// @Description: 按ID修改文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 01:00:24
// @receiver repo
// @param id
// @param in
// @return error
func (repo *rArticleCategory) UpdateByID(id uint32, in *model.ArticleCategory) error {
	err := db.NewGormCore().UpdateByID(id, in)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID
//
// @Title DeleteByID
// @Description: 按ID删除文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-16 23:51:44
// @receiver repo
// @param id
// @return error
func (repo *rArticleCategory) DeleteByID(id uint32) error {
	err := db.NewGormCore().DeleteByID(&model.ArticleCategory{ID: id})
	if err != nil {
		return err
	}

	return nil
}
