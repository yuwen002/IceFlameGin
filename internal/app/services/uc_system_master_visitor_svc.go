package services

import (
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/repositories"
	"ice_flame_gin/internal/system"
)

// sUcSystemMasterVisitor
// @Description: 访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:28:49
type sUcSystemMasterVisitor struct {
}

// NewUcSystemMasterVisitor
//
// @Title NewUcSystemMasterVisitor
// @Description: 创建一个新的 NewUcSystemMasterVisitor 服务实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:29:01
// @return *sUcSystemMasterVisitor
func NewUcSystemMasterVisitor() *sUcSystemMasterVisitor {
	return &sUcSystemMasterVisitor{}
}

// CreateVisitorCategory
//
// @Title CreateVisitorCategory
// @Description: 新建访问类型分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:57:28
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterVisitor) CreateVisitorCategory(in dto.SystemMasterVisitorCategoryInput) *system.SysResponse {
	err := repositories.NewUcSystemMasterVisitor().Insert(in)
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
