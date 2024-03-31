package services

import (
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
	"ice_flame_gin/internal/app/repositories"
	"ice_flame_gin/internal/system"
)

// sUcSystemMasterRole
//
// @Description: 表示 UcSystemMasterRole 服务
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:37:57
type sUcSystemMasterRole struct {
}

// NewUcSystemMasterRoleService
//
// @Title NewUcSystemMasterRoleService
// @Description: 创建一个新的 UcSystemMasterRole 服务实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:38:38
// @return *sUcSystemMasterRole
func NewUcSystemMasterRoleService() *sUcSystemMasterRole {
	return &sUcSystemMasterRole{}
}

// CreateMasterRole
//
// @Title CreateMasterRole
// @Description: 用户角色信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:50:41
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) CreateMasterRole(in dto.SystemMasterRoleInput) *system.SysResponse {
	err := repositories.NewUcSystemMasterRoleRepository().Insert(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    nil,
	}
}

// ShowMasterRole
//
// @Title ShowMasterRole
// @Description: 用户角色信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-13 17:04:07
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) ShowMasterRole(in dto.ListSystemMasterRoleInput) *system.SysResponse {
	// 角色信息数据列表
	out, err := repositories.NewUcSystemMasterRoleRepository().GetList(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	// 角色信息记录的数量
	totalRecords, err := repositories.NewUcSystemMasterRoleRepository().CountRecords()
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data: dto.ListSystemMasterRoleOutput{
			List:  out,
			Total: totalRecords,
		},
	}
}

// GetMasterRoleById
//
// @Title GetMasterRoleById
// @Description: 按ID获取角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-15 15:52:59
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) GetMasterRoleById(id uint32) *system.SysResponse {
	out, err := repositories.NewUcSystemMasterRoleRepository().GetById(id)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    out,
	}
}

// ChangeMasterRoleById
//
// @Title ChangeMasterRoleById
// @Description: 按ID更改角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-17 02:06:42
// @receiver svc
// @param id
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) ChangeMasterRoleById(id uint32, in dto.SystemMasterRoleInput) *system.SysResponse {
	// 根据ID更新角色信息
	err := repositories.NewUcSystemMasterRoleRepository().UpdateByID(id, &model.UcSystemMasterRole{
		Name:   in.Name,
		Remark: in.Remark,
	})
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	}
}

// ShowMasterRoleAll
//
// @Title ShowMasterRoleAll
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 11:11:57
// @receiver svc
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) ShowMasterRoleAll() *system.SysResponse {
	output, err := repositories.NewUcSystemMasterRoleRepository().GetAll()
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	var data []*dto.SelectOptionOutput
	for _, v := range output {
		d := dto.SelectOptionOutput{
			Key:   v.ID,
			Value: v.Name,
		}

		data = append(data, &d)
	}

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    data,
	}
}

// CreateMasterRoleRelation
//
// @Title CreateMasterRoleRelation
// @Description: 新建管理员角色关联信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 17:30:41
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) CreateMasterRoleRelation(in dto.SystemMasterRoleRelationInput) *system.SysResponse {
	output, err := repositories.NewUcSystemMasterRoleRelationRepository().GetOneByRoleIdAndAccountId(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if output != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "关联信息已存在",
			Data:    nil,
		}
	}

	master, err := repositories.NewUcSystemMasterRepository().GetByAccountID(in.AccountId)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if master == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "管理员信息不存在",
			Data:    nil,
		}
	}

	role, err := repositories.NewUcSystemMasterRoleRepository().GetById(in.RoleId)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if role == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "角色信息不存在",
			Data:    nil,
		}
	}
	err = repositories.NewUcSystemMasterRoleRelationRepository().Insert(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	}
}

// ShowMasterRoleRelation
//
// @Title ShowMasterRoleRelation
// @Description: 管理员角色关联信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 23:20:23
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) ShowMasterRoleRelation(in dto.ListSystemMasterRoleRelationInput) *system.SysResponse {
	out, err := repositories.NewUcSystemMasterRoleRelationRepository().GetList(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	// 角色信息记录的数量
	totalRecords, err := repositories.NewUcSystemMasterRoleRelationRepository().CountRecords()
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data: dto.ListSystemMasterRoleRoleRelationOutput{
			List:  out,
			Total: totalRecords,
		},
	}
}

// GetMasterRoleRelationById
//
// @Title GetMasterRoleRelationById
// @Description: 按ID获取管理员角色绑定信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-22 17:01:52
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) GetMasterRoleRelationById(id uint32) *system.SysResponse {
	out, err := repositories.NewUcSystemMasterRoleRelationRepository().GetById(id)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    out,
	}
}

// ChangeMasterRoleRelationById
//
// @Title ChangeMasterRoleRelationById
// @Description: 按ID修改管理员角色绑定信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-22 17:47:05
// @receiver svc
// @param id
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterRole) ChangeMasterRoleRelationById(id uint32, in dto.SystemMasterRoleRelationInput) *system.SysResponse {
	master, err := repositories.NewUcSystemMasterRepository().GetByAccountID(in.AccountId)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if master == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "管理员信息不存在",
			Data:    nil,
		}
	}

	role, err := repositories.NewUcSystemMasterRoleRepository().GetById(in.RoleId)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if role == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "角色信息不存在",
			Data:    nil,
		}
	}

	output, err := repositories.NewUcSystemMasterRoleRelationRepository().GetOneByRoleIdAndAccountId(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if output != nil {
		if output.ID != id {
			return &system.SysResponse{
				Code:    1,
				Message: "关联信息已存在",
				Data:    nil,
			}
		}
	}

	err = repositories.NewUcSystemMasterRoleRelationRepository().UpdateByID(id, &model.UcSystemMasterRoleRelation{
		AccountID: in.AccountId,
		RoleID:    in.RoleId,
	})

	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	}
}
