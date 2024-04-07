package association

import "ice_flame_gin/internal/app/models/model"

// UcSystemMasterRoleRelation
// @Description: 管理员角色关联相关表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-27 15:40:44
type UcSystemMasterRoleRelation struct {
	model.UcSystemMasterRoleRelation
	Master model.UcSystemMaster     `gorm:"foreignKey:AccountID;references:AccountID"`
	Role   model.UcSystemMasterRole `gorm:"foreignKey:RoleID"`
}

// UcSystemVisitorLogs
// @Description: 管理员访问记录相关信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-27 15:46:07
type UcSystemVisitorLogs struct {
	model.UcSystemMasterVisitorLog
	Master        model.UcSystemMaster              `gorm:"foreignKey:AccountID;references:AccountID"`
	VisitCategory model.UcSystemMasterVisitCategory `gorm:"foreignKey:VisitCategory;references:ID"`
}

// UcSystemMaster
//
// @Description: 管理员关联用户中心表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-06 23:25:58
type UcSystemMaster struct {
	model.UcSystemMaster
	Account model.UcAccount `gorm:"foreignKey:AccountID;references:ID"`
}
