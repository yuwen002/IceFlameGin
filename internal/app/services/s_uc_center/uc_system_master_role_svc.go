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