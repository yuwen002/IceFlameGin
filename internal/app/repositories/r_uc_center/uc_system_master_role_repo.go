package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	"ice_flame_gin/internal/app/models/model"
)

// rUcSystemMasterRoleRepository
//
// @Description: 管理员角色
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:10:01
type rUcSystemMasterRoleRepository struct {
	DBs map[string]*gorm.DB
	DB  *gorm.DB
}

// NewUcSystemMasterRoleRepository
//
// @Title NewUcSystemMasterRoleRepository
// @Description: 创建一个新的 rUcSystemMasterRoleRepository 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:10:55
// @return *rUcSystemMasterRoleRepository
func NewUcSystemMasterRoleRepository() *rUcSystemMasterRoleRepository {
	return &rUcSystemMasterRoleRepository{
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
func (repo *rUcSystemMasterRoleRepository) Insert(data dto.SystemMasterRoleInput) error {
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
func (repo *rUcSystemMasterRoleRepository) GetList(data dto.ListSystemMasterRoleInput) ([]*model.UcSystemMasterRole, error) {
	var systemMasterRoles []*model.UcSystemMasterRole
	err := db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Order:    data.Order,
		Page:     data.Page,
		PageSize: data.PageSize,
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
func (repo *rUcSystemMasterRoleRepository) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameUcSystemMasterRole).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}
