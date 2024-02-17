package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ice_flame_gin/config"
	"log"
)

var DB map[string]*gorm.DB

// init
//
// @Title init
// @Description: 自动初始化数据库
// @Author liuxingyu
// @Date 2024-02-03 00:54:07
func init() {
	// 当包被导入时，自动调用InitDB()函数来初始化数据库连接
	err := InitDB()
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
}

// InitDB
//
// @Title InitDB
// @Description 连接数据库
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 16:21:35
// @return error
func InitDB() error {
	//  初始化数据库连接
	DB = make(map[string]*gorm.DB)

	dbConfigs := config.GlobalConfig.Database
	for dbName, dbConfig := range dbConfigs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Name,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}

		DB[dbName] = db
	}

	return nil
}

// GetDB
//
// @Title GetDB
// @Description:
// @Author liuxingyu
// @Date 2024-02-17 11:42:13
// @param dbName
// @return *gorm.DB
func GetDB(dbName string) *gorm.DB {
	return DB[dbName]
}

// CloseDB
//
// @Title CloseDB
// @Description 关闭数据库
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 16:22:03
func CloseDB() {
	for _, db := range DB {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
}
