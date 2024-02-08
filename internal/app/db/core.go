package db

import (
	"errors"
	"gorm.io/gorm"
)

// DatabaseCore
// @Description: 数据库核心类接口
type DatabaseCore interface {
	Insert(data interface{}) error
	InsertAndGetID(data interface{}) (uint, error)
	Update(condition string, data interface{}) error
	Query(condition string) ([]interface{}, error)
	GetByID(id int, out interface{}) error
}

// GormCore
// @Description: 基于GORM的数据库核心类，该结构体实现了 DatabaseCore 接口中定义的方法
type GormCore struct {
	db *gorm.DB
}

// NewGormCore
//
// @Title NewGormCore
// @Description: 创建并返回一个新的GormCore实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-05 17:50:14
// @return *GormCore
func NewGormCore() *GormCore {
	return &GormCore{
		db: DB, // 假设db.DB是你初始化后的全局*gorm.DB实例
	}
}

// SetDefaultTable
//
// @Title SetDefaultTable
// @Description: 设置默认表名
// @Author liuxingyu
// @Date 2024-02-06 17:47:57
// @receiver g
// @param tableName
// @return *GormCore
func (g *GormCore) SetDefaultTable(tableName string) *GormCore {
	// 设置默认表名并返回新的 GormCore 对象
	g.db = g.db.Table(tableName)
	return g
}

// Insert
//
// @Title Insert
// @Description: 新增数据
// @Author liuxingyu
// @Date 2024-02-03 02:20:47
// @receiver g
// @param data
// @return error
func (g *GormCore) Insert(data interface{}) error {
	return g.db.Create(data).Error
}

// InsertAndGetID
//
// @Title InsertAndGetID
// @Description: 插入数据并返回 ID
// @Author liuxingyu
// @Date 2024-02-03 23:11:48
// @receiver g
// @param data
// @return uint
// @return error
func (g *GormCore) InsertAndGetID(data interface{}) (uint, error) {
	result := g.db.Create(data)
	if result.Error != nil {
		return 0, result.Error
	}

	var id uint
	if err := g.db.Model(data).Pluck("id", &id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

// Update
//
// @Title Update
// @Description: 条件修改数据
// @Author liuxingyu
// @Date 2024-02-03 23:26:09
// @receiver g
// @param condition
// @param data
// @return error
func (g *GormCore) Update(condition string, data interface{}) error {
	return g.db.Where(condition).Updates(data).Error
}

// UpdateByID
//
// @Title UpdateByID
// @Description: 按ID修改数据
// @Author liuxingyu
// @Date 2024-02-03 23:27:38
// @receiver g
// @param id
// @param data
// @return error
func (g *GormCore) UpdateByID(id uint, data interface{}) error {
	return g.db.Model(data).Where("id = ?", id).Updates(data).Error
}

// Query
//
// @Title Query
// @Description:  条件查询数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-05 15:25:15
// @receiver g
// @param condition
// @param out
// @return error
func (g *GormCore) Query(condition string, out interface{}) error {
	err := g.db.Where(condition).Find(out).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByID
//
// @Title GetByID
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-05 16:38:25
// @receiver g
// @param id
// @param out
// @return error
func (g *GormCore) GetByID(id int, out interface{}) error {
	return g.db.Where("id = ?", id).First(out).Error
}

// QueryOne
//
// @Title QueryOne
// @Description: 条件查询一条数据
// @Author liuxingyu
// @Date 2024-02-08 01:30:02
// @receiver g
// @param out
// @param condition
// @param args
// @return error
func (g *GormCore) QueryOne(out interface{}, condition string, args ...interface{}) error {
	err := g.db.Where(condition, args...).First(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 查询未找到数据
			return nil
		}
		return err
	}
	return nil
}
