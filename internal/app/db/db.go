package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"ice_flame_gin/config"
	"log"
)

var DB *gorm.DB

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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBConfig.Username,
		config.DBConfig.Password,
		config.DBConfig.Host,
		config.DBConfig.Port,
		config.DBConfig.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

// CloseDB
//
// @Title CloseDB
// @Description 关闭数据库
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 16:22:03
func CloseDB() {
	if DB != nil {
		db, _ := DB.DB()
		db.Close()
	}
}
