package services

import (
	"ice_flame_gin/internal/app/dto"
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
	err := repositories.NewSinglePageRepository().Insert(in)
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
