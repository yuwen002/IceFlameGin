package services

import (
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	repositories "ice_flame_gin/internal/app/repositories/r_uc_center"
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
