package dto

import (
	"ice_flame_gin/internal/app/models/association"
	"ice_flame_gin/internal/app/models/model"
)

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
	Order  string
	Start  int
	Length int
}

// ListSystemMasterRoleOutput
// @Description: 角色信息列表输出信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-14 14:01:54
type ListSystemMasterRoleOutput struct {
	List  []*model.UcSystemMasterRole // 角色信息列表
	Total int64                       // 角色信息记录数量
}

// SystemMasterRoleRelationInput
//
// @Description: 管理员关联角色信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 00:24:53
type SystemMasterRoleRelationInput struct {
	AccountId uint32
	RoleId    uint32
}

// ListSystemMasterRoleRelationInput
//
// @Description: 管理员关联角色信息列表输入信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 23:19:43
type ListSystemMasterRoleRelationInput struct {
	Order  string
	Start  int
	Length int
}

// ListSystemMasterRoleRoleRelationOutput
//
// @Description: 管理员关联角色信息列表输出信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-19 23:44:05
type ListSystemMasterRoleRoleRelationOutput struct {
	List  []*association.UcSystemMasterRoleRelation // 角色信息列表
	Total int64                                     // 角色信息记录数量
}
