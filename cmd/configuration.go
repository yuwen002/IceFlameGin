package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// main
//
// @Title main
// @Description  @todo  测试代码，待完善
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-19 15:45:12
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "../test",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	gormDb, _ := gorm.Open(mysql.Open(""))
	g.UseDB(gormDb) // reuse your gorm db

	// 按照约定为结构体 `model.User` 生成类型安全的 DAO API
	g.ApplyBasic(
		// 根据 `user` 表生成结构 `User`
		g.GenerateModel("single_page"),

		// 根据 `user` 表生成结构 `Employee`
		//g.GenerateModelAs("users", "Employee"),

		// 根据 `user` 表和生成时选项生成结构 `User`
		//g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),
	)

	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateAllTable()...,
	)

	// 生成代码
	g.Execute()
}
