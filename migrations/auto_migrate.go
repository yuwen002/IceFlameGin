package migrations

import (
	db2 "ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
	"ice_flame_gin/internal/app/services"
)

// CreateDatabase
//
// @Title CreateDatabase
// @Description: 新建数据库表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-09-04 18:14:26
func CreateDatabase() {
	db := db2.DB["default"]
	db.AutoMigrate(
		&model.UcAccount{},
		&model.UcSystemMaster{},
		&model.UcSystemMasterRole{},
		&model.UcSystemMasterRoleRelation{},
		&model.UcSystemMasterVisitCategory{},
		&model.UcSystemMasterVisitorLog{},
		&model.SinglePage{},
		&model.ArticleChannel{},
		&model.ArticleTag{},
		&model.ArticleTagOrm{},
		&model.Article{},
	)
}

func SyncData() {
	services.NewUcSystemMasterService().Register(dto.RegisterSystemMasterInput{
		Password: "123456",
		Tel:      "15566036902",
		Name:     "yuwen002",
		Email:    "yuwen002@163.com",
	})
}
