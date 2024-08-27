package migrations

import (
	db2 "ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/models/model"
)

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
