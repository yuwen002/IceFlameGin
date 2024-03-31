package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
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
func (repo *rUcSystemMasterRoleRelation) Insert(data dto.SystemMasterRoleRelationInput) error {
	err := db.NewGormCore().Insert(&model.UcSystemMasterRoleRelation{
		AccountID: data.AccountId,
		RoleID:    data.RoleId,
	})

	if err != nil {
		return err
	}

	return nil
}

// GetById
//
// @Title GetById
// @Description: 按ID获取管理员关联角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-22 17:05:31
// @receiver r
// @param id
// @return *model.UcSystemMasterRoleRelation
// @return error
func (repo *rUcSystemMasterRoleRelation) GetById(id uint32) (*model.UcSystemMasterRoleRelation, error) {
	var relation *model.UcSystemMasterRoleRelation
	condition := "id = ?"
	err := db.NewGormCore().QueryOne(&relation, condition, id)
	if err != nil {
		return nil, nil
	}

	return relation, nil
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
func (repo *rUcSystemMasterRoleRelation) GetByIdHasOne(id uint32) (*association.UcSystemMasterRoleRelation, error) {
	var relation *association.UcSystemMasterRoleRelation
	condition := "id = ?"
	associations := []string{"Role", "Master"}
	err := db.NewGormCore().QueryWithHasOne(&relation, associations, condition, id)
	if err != nil {
		return nil, nil
	}

	return relation, nil
}

// GetList
//
// @Title GetList
// @Description: 管理员关联角色列表信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 23:30:42
// @receiver r
// @param data
// @return []*association.UcSystemMasterRoleRelation
// @return error
func (repo *rUcSystemMasterRoleRelation) GetList(data dto.ListSystemMasterRoleRelationInput) ([]*association.UcSystemMasterRoleRelation, error) {
	var relations []*association.UcSystemMasterRoleRelation
	err := db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Order:    "id desc",
		Preload:  []string{"Master", "Role"},
		PageType: 2,
		Limit: db.Limit{
			Length: data.Length,
			Offset: data.Start,
		},
	}, &relations)

	if err != nil {
		return nil, err
	}

	return relations, nil
}

// CountRecords
//
// @Title CountRecords
// @Description: 管理员关联角色列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 17:17:10
// @receiver r
// @return int64
// @return error
func (repo *rUcSystemMasterRoleRelation) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameUcSystemMasterRoleRelation).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}

// GetOneByRoleIdAndAccountId
//
// @Title GetOneByRoleIdAndAccountId
// @Description: 条件查询管理员角色关联
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 17:26:59
// @receiver r
// @param in
// @return *model.UcSystemMasterRoleRelation
// @return error
func (repo *rUcSystemMasterRoleRelation) GetOneByRoleIdAndAccountId(in dto.SystemMasterRoleRelationInput) (*model.UcSystemMasterRoleRelation, error) {
	var out *model.UcSystemMasterRoleRelation
	condition := "role_id =  ? and account_id = ?"
	err := db.NewGormCore().QueryOne(&out, condition, in.RoleId, in.AccountId)
	if err != nil {
		return nil, nil
	}

	return out, nil
}

// UpdateByID
//
// @Title UpdateByID
// @Description: 按ID更改管理员角色关联
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-22 17:44:49
// @receiver repo
// @param id
// @param in
// @return error
func (repo *rUcSystemMasterRoleRelation) UpdateByID(id uint32, in *model.UcSystemMasterRoleRelation) error {
	err := db.NewGormCore().UpdateByID(id, in)
	if err != nil {
		return err
	}

	return nil
}
