package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
	"ice_flame_gin/internal/app/repositories"
	"ice_flame_gin/internal/pkg/utils"
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

// ShowVisitCategoryAll
//
// @Title ShowVisitCategoryAll
// @Description: 全部访问类型分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 15:09:43
// @receiver svc
// @return *system.SysResponse
func (svc *sUcSystemMasterVisit) ShowVisitCategoryAll() *system.SysResponse {
	out, err := repositories.NewUcSystemMasterVisit().GetAll()
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	var data []*dto.SelectOptionOutput
	for _, v := range out {
		d := dto.SelectOptionOutput{
			Key:   v.ID,
			Value: v.Title,
		}
		data = append(data, &d)
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    data,
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

// CreateVisitorLogs
//
// @Title CreateVisitorLogs
// @Description: 管理员访问记录写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 21:37:29
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterVisit) CreateVisitorLogs(in dto.SystemMasterVisitorLogsInput) *system.SysResponse {
	err := repositories.NewUcSystemMasterVisitorLogs().Insert(in)

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

// ShowVisitorLogs
//
// @Title ShowVisitorLogs
// @Description: 管理员访问及记录列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-27 16:10:24
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMasterVisit) ShowVisitorLogs(in dto.ListSystemMasterVisitorLogsInput) *system.SysResponse {
	fmt.Println("in:", in)
	out, err := repositories.NewUcSystemMasterVisitorLogs().GetListByWhere(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	var totalRecords int64
	if in.Condition == "" {
		totalRecords, err = repositories.NewUcSystemMasterVisitorLogs().CountRecords()
	} else {
		totalRecords, err = repositories.NewUcSystemMasterVisitorLogs().CountWhereRecords(in)
	}
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
		Data: dto.ListSystemMasterVisitorLogsOutput{
			List:  out,
			Total: totalRecords,
		},
	}
}

// WriteSystemMasterVisitorLogs
//
// @Title SystemMasterVisitorLogs
// @Description: 写入用户访问记录
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 22:46:17
// @param accountID
// @param osCategory
// @param visitCategory
// @param unionID
// @param description
// @param
func (svc *sUcSystemMasterVisit) WriteSystemMasterVisitorLogs(c *gin.Context, osCategory uint32, visitCategory uint32, unionID uint32, description string) error {
	// 管理员AccountID
	accountID, err := system.GetMasterID(c)
	if err != nil {
		return err
	}

	// 获取管理员IP
	ip := c.ClientIP()
	ipLong, err := utils.IPToLong(ip)
	if err != nil {
		return err
	}

	ipLongString, err := utils.ToString(ipLong)
	if err != nil {
		return err
	}

	svc.CreateVisitorLogs(dto.SystemMasterVisitorLogsInput{
		AccountID:     accountID,
		OsCategory:    osCategory,
		VisitCategory: visitCategory,
		UnionID:       unionID,
		Description:   description,
		IPLong:        ipLongString,
		IP:            ip,
	})

	return nil
}
