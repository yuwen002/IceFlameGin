package db

import "gorm.io/gorm"

// DatabaseCore
// @Description: 数据库核心类接口
type DatabaseCore interface {
	Insert(data interface{}) error
	InsertAndGetID(data interface{}) (uint, error)
	Update(condition string, data interface{}) error
	Query(condition string) ([]interface{}, error)
}

// GormCore
// @Description: 基于GORM的数据库核心类，该结构体实现了 DatabaseCore 接口中定义的方法
type GormCore struct {
	db *gorm.DB
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

func (g *GormCore) Query(condition string) ([]interface{}, error) {

	var users []User
	err := g.db.Where(condition).Find(&users).Error
	if err != nil {
		return nil, err
	}
	result := make([]interface{}, len(users))
	for i := range users {
		result[i] = &users[i]
	}
	return result, nil
}
