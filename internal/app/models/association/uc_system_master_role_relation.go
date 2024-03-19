package association

import "ice_flame_gin/internal/app/models/model"

type UcSystemMasterRoleRelation struct {
	model.UcSystemMasterRoleRelation
	Master model.UcSystemMaster     `gorm:"references:AccountID"`
	Role   model.UcSystemMasterRole `gorm:"foreignKey:RoleID"`
}
