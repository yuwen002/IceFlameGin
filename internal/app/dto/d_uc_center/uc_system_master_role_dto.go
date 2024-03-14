package dto

import "ice_flame_gin/internal/app/models/model"

// SystemMasterRoleInput
//
// @Description: 角色信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-08 00:17:30
type SystemMasterRoleInput struct {
	Name         string
	Remark       string
	SupperMaster bool
}

// ListSystemMasterRoleInput
// @Description: 角色信息列表输入信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-14 14:01:34
type ListSystemMasterRoleInput struct {
	Order    string
	Page     int
	PageSize int
}

// ListSystemMasterRoleOutput
// @Description: 角色信息列表输出信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-14 14:01:54
type ListSystemMasterRoleOutput struct {
	List  []*model.UcSystemMasterRole // 角色信息列表
	Total int64                       // 角色信息记录数量
}
