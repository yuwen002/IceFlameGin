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
	err = db.NewGormCore().QueryListWithCondition(db.QueryOptions{
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

// GetByID
//
// @Title GetByID
// @Description: 按ID获取文章标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-01 23:59:05
// @receiver repo
// @param id
// @return out
// @return err
func (repo *rArticleTag) GetByID(id uint32) (out *model.ArticleTag, err error) {
	err = db.NewGormCore().GetByID(id, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateByID
//
// @Title UpdateByID
// @Description: 按ID更新文章标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-02 00:15:21
// @receiver repo
// @param id
// @param in
// @return error
func (repo *rArticleTag) UpdateByID(id uint32, in *model.ArticleTag) error {
	err := db.NewGormCore().UpdateByID(id, in)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID
//
// @Title DeleteByID
// @Description: 按ID删除文章标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-10 13:58:24
// @receiver repo
// @param id
// @return error
func (repo *rArticleTag) DeleteByID(id uint32) error {
	err := db.NewGormCore().DeleteByID(&model.ArticleTag{ID: id})
	if err != nil {
		return err
	}

	return nil
}

// GetAll
//
// @Title GetAll
// @Description: 获取全部文章标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-23 14:47:05
// @receiver repo
// @return *model.ArticleTag
// @return error
func (repo *rArticleTag) GetAll() ([]*model.ArticleTag, error) {
	var out []*model.ArticleTag
	err := db.NewGormCore().GetAll(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
