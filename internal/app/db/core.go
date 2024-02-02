package db

import "gorm.io/gorm"

// DatabaseCore
// @Description: 数据库核心类接口
type DatabaseCore interface {
	Insert(data interface{}) error
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

func (g *GormCore) Update(condition string, data interface{}) error {
	return g.db.Where(condition).Updates(data).Error
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
