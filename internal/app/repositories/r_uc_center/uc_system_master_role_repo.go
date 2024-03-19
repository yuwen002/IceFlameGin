package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	"ice_flame_gin/internal/app/models/model"
)

// rUcSystemMasterRole
//
// @Description: 管理员角色
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:10:01
type rUcSystemMasterRole struct {
	DBs map[string]*gorm.DB
	DB  *gorm.DB
}

// NewUcSystemMasterRoleRepository
//
// @Title NewUcSystemMasterRoleRepository
// @Description: 创建一个新的 rUcSystemMasterRole 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:10:55
// @return *rUcSystemMasterRole
func NewUcSystemMasterRoleRepository() *rUcSystemMasterRole {
	return &rUcSystemMasterRole{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 管理员角色写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:27:21
// @receiver repo
// @param data
// @return error
func (repo *rUcSystemMasterRole) Insert(data dto.SystemMasterRoleInput) error {
	err := db.NewGormCore().Insert(&model.UcSystemMasterRole{
		Name:         data.Name,
		Remark:       data.Remark,
		SupperMaster: data.SupperMaster,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetList
//
// @Title GetList
// @Description: 管理员角色列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-13 16:45:08
// @receiver repo
// @param data
// @return *model.UcSystemMasterRole
// @return error
func (repo *rUcSystemMasterRole) GetList(data dto.ListSystemMasterRoleInput) ([]*model.UcSystemMasterRole, error) {
	var systemMasterRoles []*model.UcSystemMasterRole
	err := db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Order:    data.Order,
		PageType: 2,
		Limit: db.Limit{
			Length: data.Length,
			Offset: data.Start,
		},
	}, &systemMasterRoles)

	if err != nil {
		return nil, err
	}

	return systemMasterRoles, nil
}

// CountRecords
//
// @Title CountRecords
// @Description: 管理员角色列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-14 11:14:13
// @receiver repo
// @return int64
// @return error
func (repo *rUcSystemMasterRole) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameUcSystemMasterRole).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}

// GetById
//
// @Title GetById
// @Description: 按ID获取角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-15 15:48:01
// @receiver repo
// @param id
// @return *model.UcSystemMasterRole
// @return error
func (repo *rUcSystemMasterRole) GetById(id uint32) (*model.UcSystemMasterRole, error) {
	var role *model.UcSystemMasterRole
	err := db.NewGormCore().GetByID(id, &role)
	if err != nil {
		return nil, err
	}

	if role.ID == 0 {
		return nil, nil
	}

	return role, nil
}

// UpdateByID
//
// @Title UpdateByID
// @Description: 按ID修改角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-17 02:05:00
// @receiver repo
// @param id
// @param in
// @return error
func (repo *rUcSystemMasterRole) UpdateByID(id uint32, in *model.UcSystemMasterRole) error {
	err := db.NewGormCore().UpdateByID(id, in)
	if err != nil {
		return err
	}

	return nil
}
