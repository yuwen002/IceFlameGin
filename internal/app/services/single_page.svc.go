package services

import (
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
	"ice_flame_gin/internal/app/repositories"
	"ice_flame_gin/internal/system"
)

// sSinglePage
//
// @Description: 表示 SinglePage 服务
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 23:21:32
type sSinglePage struct {
}

// NewSinglePageService
//
// @Title NewSinglePageService
// @Description: 创建一个新的 sSinglePage 服务实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 23:22:04
// @return *sSinglePage
func NewSinglePageService() *sSinglePage {
	return &sSinglePage{}
}

// CreateSinglePage
//
// @Title CreateSinglePage
// @Description: 新建单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 23:23:33
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sSinglePage) CreateSinglePage(in dto.SinglePageInput) *system.SysResponse {
	err := repositories.NewSinglePageRepository().Insert(&model.SinglePage{
		Title:       in.Title,
		Description: in.Description,
		Keyword:     in.Keyword,
		Content:     in.Content,
		Thumbnail:   in.Thumbnail,
		Click:       in.Click,
		Status:      in.Status,
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
		Message: "success",
		Data:    nil,
	}
}

// ShowSinglePage
//
// @Title ShowSinglePage
// @Description: 单页信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-02 00:22:25
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sSinglePage) ShowSinglePage(in dto.ListSinglePageInput) *system.SysResponse {
	out, err := repositories.NewSinglePageRepository().GetList(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	totalRecords, err := repositories.NewSinglePageRepository().CountRecords()
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
		Data: dto.ListSinglePageOutput{
			List:  out,
			Total: totalRecords,
		},
	}
}

// GetSinglePageByID
//
// @Title GetSinglePageByID
// @Description: 按ID获取单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-03 15:32:54
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sSinglePage) GetSinglePageByID(id uint32) *system.SysResponse {
	out, err := repositories.NewSinglePageRepository().GetByID(id)
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

// ChangeSinglePageByID
//
// @Title ChangeSinglePageByID
// @Description: 按ID修改单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-04 23:37:35
// @receiver svc
// @param id
// @param in
// @return *system.SysResponse
func (svc *sSinglePage) ChangeSinglePageByID(id uint32, in dto.SinglePageInput) *system.SysResponse {
	err := repositories.NewSinglePageRepository().UpdateByID(id, &model.SinglePage{
		Title:       in.Title,
		Description: in.Description,
		Keyword:     in.Keyword,
		Content:     in.Content,
		Thumbnail:   in.Thumbnail,
		Click:       in.Click,
		Status:      in.Status,
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
