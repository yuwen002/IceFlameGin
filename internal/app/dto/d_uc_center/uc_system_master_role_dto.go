package dto

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

type SystemMasterRoleOutput struct {
	Order    string
	Page     int
	PageSize int
}
