package db

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
)

// DatabaseCore
// @Description: 数据库核心类接口
type DatabaseCore interface {
	Select(field string) *GormCore
	Omit(field string) *GormCore
	SetDefaultTable(tableName string) *GormCore
	SetModel(model interface{}) *GormCore
	Insert(data interface{}) error
	InsertAndGetID(data interface{}) (uint32, error)
	Update(condition string, data interface{}) error
	UpdateByID(id uint32, data interface{}) error
	UpdateColumnsByID(id uint32, data interface{}) error
	Query(condition string, out interface{}) error
	GetByID(id uint32, out interface{}) error
	QueryOne(out interface{}, condition string, args ...interface{}) error
	QueryListWithCondition(opts QueryOptions, out interface{}) error
	Count() (int64, error)
	CountWhere(condition string, args ...interface{}) (int64, error)
	QueryWithHasOne(out interface{}, associations []string, condition string, args ...interface{}) error
	GetAll(out interface{}) error
	DeleteByID(model interface{}) error
}

// GormCore
// @Description: 基于GORM的数据库核心类，该结构体实现了 DatabaseCore 接口中定义的方法
type GormCore struct {
	db  *gorm.DB
	dbs map[string]*gorm.DB
}

// NewGormCore
//
// @Title NewGormCore
// @Description: 创建并返回一个新的GormCore实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-05 17:50:14
// @return *GormCore
func NewGormCore() DatabaseCore {
	// 默认链接default数据库
	firstDB := DB["default"]
	firstDB = firstDB.Debug() // 打开调试模式
	return &struct {
		*GormCore
	}{
		GormCore: &GormCore{
			db:  firstDB, // db.DB是初始化后的全局*gorm.DB实例
			dbs: DB,
		},
	}
}

// Select
//
// @Title Select
// @Description: Select选择某些字段
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-08 22:37:55
// @receiver g
// @param field
// @return *GormCore
func (g *GormCore) Select(field string) *GormCore {
	g.db = g.db.Select(field)
	return g
}

// Omit
//
// @Title Omit
// @Description: Omit忽略某些字段
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-08 22:39:56
// @receiver g
// @param field
// @return *GormCore
func (g *GormCore) Omit(field string) *GormCore {
	g.db = g.db.Omit(field)
	return g
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

// SetModel
//
// @Title SetModel
// @Description: 设置默认Model
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-06 14:06:48
// @receiver g
// @param model
func (g *GormCore) SetModel(model interface{}) *GormCore {
	g.db = g.db.Model(model)
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
func (g *GormCore) InsertAndGetID(data interface{}) (uint32, error) {
	result := g.db.Create(data)
	if result.Error != nil {
		return 0, result.Error
	}

	// 检查插入的记录是否包含自增主键 ID
	if result.RowsAffected == 0 {
		return 0, errors.New("no record inserted")
	}

	// 使用反射获取插入记录的 ID
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	idField := value.FieldByName("ID")
	if !idField.IsValid() {
		return 0, errors.New("failed to get ID from inserted record")
	}
	id := uint32(idField.Uint())

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
// @Description: 按ID修改数据，会忽略那些零值字段。
// @Author liuxingyu
// @Date 2024-02-03 23:27:38
// @receiver g
// @param id
// @param data
// @return error
func (g *GormCore) UpdateByID(id uint32, data interface{}) error {
	return g.db.Model(data).Where("id = ?", id).Updates(data).Error
}

// UpdateColumnsByID
//
// @Title UpdateColumnsByID
// @Description: 按ID修改数据，包括零值字段。
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-08 15:30:09
// @receiver g
// @param id
// @param data
// @return error
func (g *GormCore) UpdateColumnsByID(id uint32, data interface{}) error {
	return g.db.Model(data).Where("id = ?", id).UpdateColumns(data).Error
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
func (g *GormCore) GetByID(id uint32, out interface{}) error {
	err := g.db.Where("id = ?", id).First(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			v := reflect.ValueOf(out)
			if v.Kind() == reflect.Ptr && !v.IsNil() {
				v.Elem().Set(reflect.Zero(v.Elem().Type())) // 设置结构体指针的内容为零值
			}
			return nil
		}
		return err
	}
	return nil
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
			v := reflect.ValueOf(out)
			if v.Kind() == reflect.Ptr && !v.IsNil() {
				v.Elem().Set(reflect.Zero(v.Elem().Type())) // 设置结构体指针的内容为零值
			}
			return nil
		}
		return err
	}
	return nil
}

// QueryListWithCondition
//
// @Title QueryListWithCondition
// @Description: 方法用于根据条件查询列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-13 16:30:34
// @receiver g
// @param opts
// @param out
// @return error
func (g *GormCore) QueryListWithCondition(opts QueryOptions, out interface{}) error {
	db := g.db

	// 查询字段
	if opts.Field != "" {
		db = db.Select(opts.Field)
	}

	// 条件查询
	if opts.Condition != "" {
		if len(opts.Args) > 0 {
			db = db.Where(opts.Condition, opts.Args...)
		} else {
			db = db.Where(opts.Condition)
		}
	}

	if opts.Omit != "" {
		db = db.Omit(opts.Omit)
	}

	// has one 关联
	if len(opts.Preload) > 0 {
		for _, association := range opts.Preload {
			db = db.Preload(association)
		}
	}

	// PreloadFunc 处理预加载函数映射
	if len(opts.PreloadFunc) > 0 {
		for association, preloadFunc := range opts.PreloadFunc {
			if preloadFunc != nil {
				db = db.Preload(association, preloadFunc)
			} else {
				db = db.Preload(association)
			}
		}
	}

	// 分页类型判断
	switch opts.PageType {
	case 1:
		// 判断是否需要分页
		if opts.Page > 0 {
			if opts.Size == 0 {
				// 分页默认长度
				opts.Size = 10
			}
			// 计算偏移量
			offset := (opts.Page - 1) * opts.Size
			db = db.Offset(offset).Limit(opts.Size)
		}
	case 2:
		// 判断分页
		if opts.Length > 0 {
			db = db.Offset(opts.Offset).Limit(opts.Length)
		} else if opts.Offset > 0 {
			db = db.Limit(opts.Offset)
		}
	}

	db = db.Order(opts.Order)

	// 查询数据
	err := db.Debug().Find(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			v := reflect.ValueOf(out)
			if v.Kind() == reflect.Ptr && !v.IsNil() {
				if v.Elem().Kind() == reflect.Slice && v.Elem().Len() == 0 {
					return nil
				}
				v.Elem().Set(reflect.Zero(v.Elem().Type())) // 设置切片的内容为零值
			}
			return nil
		}
		return err
	}

	return nil
}

// Count
//
// @Title Count
// @Description: 计算记录数量
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-14 11:00:34
// @receiver g
// @param condition
// @param args
// @return int64
// @return error
func (g *GormCore) Count() (int64, error) {
	var count int64
	err := g.db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CountWhere
//
// @Title CountWhere
// @Description: 根据条件计算记录的数量
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-14 11:03:13
// @receiver g
// @param condition
// @param args
// @return int64
// @return error
func (g *GormCore) CountWhere(condition string, args ...interface{}) (int64, error) {
	var count int64
	err := g.db.Where(condition, args...).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryWithHasOne
//
// @Title QueryWithHasOne
// @Description: 一对一关联的查询一条数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 00:39:07
// @receiver g
// @param out
// @param associations
// @param condition
// @param args
// @return error
func (g *GormCore) QueryWithHasOne(out interface{}, associations []string, condition string, args ...interface{}) error {
	// 将条件添加到查询中
	g.db = g.db.Where(condition, args...)

	// 遍历关联表，对每个关联表进行预加载
	for _, association := range associations {
		g.db = g.db.Preload(association)
	}

	// 执行查询并获取结果
	err := g.db.First(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			v := reflect.ValueOf(out)
			if v.Kind() == reflect.Ptr && !v.IsNil() {
				v.Elem().Set(reflect.Zero(v.Elem().Type())) // 设置结构体指针的内容为零值
			}
			return nil
		}
		return err
	}
	return nil
}

// GetAll
//
// @Title GetAll
// @Description: 获取所有数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 10:50:17
// @receiver g
// @param out
// @return error
func (g *GormCore) GetAll(out interface{}) error {
	err := g.db.Find(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			v := reflect.ValueOf(out)
			if v.Kind() == reflect.Ptr && !v.IsNil() {
				if v.Elem().Kind() == reflect.Slice && v.Elem().Len() == 0 {
					return nil
				}
				v.Elem().Set(reflect.Zero(v.Elem().Type())) // 设置切片的内容为零值
			}
			return nil
		}
		return err
	}
	return nil
}

// DeleteByID
//
// @Title DeleteByID
// @Description: 按ID删除数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-06 14:11:58
// @receiver g
// @param model
// @return error
func (g *GormCore) DeleteByID(model interface{}) error {
	err := g.db.Delete(model).Error
	if err != nil {
		return err
	}
	return nil
}
