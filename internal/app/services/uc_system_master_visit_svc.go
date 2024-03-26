package services

import (
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
	"ice_flame_gin/internal/app/repositories"
	"ice_flame_gin/internal/system"
)

// sUcSystemMasterVisit
// @Description: 访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:28:49
type sUcSystemMasterVisit struct {
}

// NewUcSystemMasterVisit
//
// @Title NewUcSystemMasterVisit
// @Description: 创建一个新的 NewUcSystemMasterVisit 服务实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:29:01
// @return *sUcSystemMasterVisit
func NewUcSystemMasterVisit() *sUcSystemMasterVisit {
	return &sUcSystemMasterVisit{}
}

// CreateVisitCategory
//
// @Title CreateVisitCategory
// @Description: 新建访问类型分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:57:28
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterVisit) CreateVisitCategory(in dto.SystemMasterVisitCategoryInput) *system.SysResponse {
	err := repositories.NewUcSystemMasterVisit().Insert(in)
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

// ShowVisitCategory
//
// @Title ShowVisitCategory
// @Description: 访问类型分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 22:08:28
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterVisit) ShowVisitCategory(in dto.ListSystemMasterVisitCategoryInput) *system.SysResponse {
	out, err := repositories.NewUcSystemMasterVisit().GetList(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	totalRecords, err := repositories.NewUcSystemMasterVisit().CountRecords()
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
		Data: dto.ListSystemMasterVisitCategoryOutput{
			List:  out,
			Total: totalRecords,
		},
	}
}

// GetVisitCategoryByID
//
// @Title GetVisitCategoryByID
// @Description: 按ID获取访问类型分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 14:40:24
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sUcSystemMasterVisit) GetVisitCategoryByID(id uint32) *system.SysResponse {
	out, err := repositories.NewUcSystemMasterVisit().GetById(id)
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

// ChangeVisitCategoryByID
//
// @Title ChangeVisitCategoryByID
// @Description: 按ID更改访问类型分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 15:08:07
// @receiver svc
// @param id
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterVisit) ChangeVisitCategoryByID(id uint32, in dto.SystemMasterVisitCategoryInput) *system.SysResponse {
	err := repositories.NewUcSystemMasterVisit().UpdateByID(id, &model.UcSystemMasterVisitCategory{
		Title: in.Title,
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
