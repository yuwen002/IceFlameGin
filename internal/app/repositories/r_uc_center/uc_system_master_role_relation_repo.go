package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	"ice_flame_gin/internal/app/models/association"
	"ice_flame_gin/internal/app/models/model"
)

// rUcSystemMasterRoleRelation
//
// @Description: 角色信息关联信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 00:18:53
type rUcSystemMasterRoleRelation struct {
	DBs map[string]*gorm.DB
	DB  *gorm.DB
}

// NewUcSystemMasterRoleRelationRepository
//
// @Title NewUcSystemMasterRoleRelationRepository
// @Description: 创建一个新的 rUcSystemMasterRoleRelation 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 00:20:25
// @return *rUcSystemMasterRoleRelation
func NewUcSystemMasterRoleRelationRepository() *rUcSystemMasterRoleRelation {
	return &rUcSystemMasterRoleRelation{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 管理员关联角色信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 00:27:44
// @receiver r
// @param data
// @return error
func (r *rUcSystemMasterRoleRelation) Insert(data dto.SystemMasterRoleRelationInput) error {
	err := db.NewGormCore().Insert(&model.UcSystemMasterRoleRelation{
		AccountID: data.AccountId,
		RoleID:    data.RoleId,
	})

	if err != nil {
		return err
	}

	return nil
}

// GetByIdHasOne
//
// @Title GetByIdHasOne
// @Description: 按ID获取管理员关联角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 00:47:33
// @receiver r
// @param id
// @return *association.UcSystemMasterRoleRelation
// @return error
func (r *rUcSystemMasterRoleRelation) GetByIdHasOne(id uint32) (*association.UcSystemMasterRoleRelation, error) {
	var relation *association.UcSystemMasterRoleRelation
	condition := "id = ?"
	associations := []string{"Role", "Master"}
	err := db.NewGormCore().QueryWithHasOne(&relation, associations, condition, id)
	if err != nil {
		return nil, err
	}

	return relation, nil
}
