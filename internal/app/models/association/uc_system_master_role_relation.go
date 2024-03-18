package association

import "ice_flame_gin/internal/app/models/model"

type UcSystemMasterRoleRelation struct {
	model.UcSystemMasterRoleRelation
	UcSystemMaster     model.UcSystemMaster     `gorm:"foreignKey:AccountID"`
	UcSystemMasterRole model.UcSystemMasterRole `gorm:"foreignKey:AccountID"`
}
