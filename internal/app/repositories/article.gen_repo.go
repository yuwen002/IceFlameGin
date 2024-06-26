// Code generated repository by gen.
// Date 2024-05-21 23:59:31

package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/models/model"
)

// rArticle
//
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-21 23:59:31
type rArticle struct {
	DB *gorm.DB
}

// NewArticle
//
// @Title NewArticle
// @Description:  创建一个新的 rArticle 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-21 23:59:31
// @return *rArticle
func NewArticle() interface{} {
	return &rArticle{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 数据信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-21 23:59:31
// @receiver repo
// @param in
// @return error
func (repo *rArticle) Insert(data *model.Article) error {
	err := db.NewGormCore().Insert(data)

	if err != nil {
		return err
	}

	return nil
}

// InsertAndGetID
//
// @Title InsertAndGetID
// @Description: 数据信息写入，并返回写入ID
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-21 23:59:31
// @receiver repo
// @param data
// @return uint32
// @return error
func (repo *rArticle) InsertAndGetID(data *model.Article) (uint32, error) {
	id, err := db.NewGormCore().InsertAndGetID(data)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateByID
//
// @Title UpdateByID
// @Description: 按ID修改数据信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-21 23:59:31
// @receiver repo
// @param id
// @param in
// @return error
func (repo *rArticle) UpdateByID(id uint32, in *model.Article) error {
	err := db.NewGormCore().UpdateByID(id, in)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID
//
// @Title DeleteByID
// @Description: 按ID删除数据信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-21 23:59:31
// @receiver repo
// @param id
// @return error
func (repo *rArticle) DeleteByID(id uint32) error {
	err := db.NewGormCore().DeleteByID(&model.Article{ID: id})
	if err != nil {
		return err
	}

	return nil
}

// CountRecords
//
// @Title CountRecords
// @Description: 数据信息列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-21 23:59:31
// @receiver repo
// @return int64
// @return error
func (repo *rArticle) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetModel(&model.Article{}).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}

// GetList
//
// @Title GetList
// @Description: 数据信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-21 23:59:31
// @receiver repo
// @param data
// @return out
// @return err
func (repo *rArticle) GetList(data db.GetListOptions) ([]*model.Article, error) {
	var out []*model.Article
	err := db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Order:       data.Order,
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
