package services

import (
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
	"ice_flame_gin/internal/app/repositories"
	"ice_flame_gin/internal/system"
)

// sArticle
//
// @Description: Article 服务
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-09 23:57:29
type sArticle struct {
}

// NewArticleService
//
// @Title NewArticleService
// @Description: 创建一个新的 sArticle 服务实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-09 23:59:39
// @return *sArticle
func NewArticleService() *sArticle {
	return &sArticle{}
}

// CreateArticleCategory
//
// @Title CreateArticleCategory
// @Description: 新建文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-10 00:11:52
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sArticle) CreateArticleCategory(in *dto.ArticleCategoryInput) *system.SysResponse {
	err := repositories.NewArticleCategoryRepository().Insert(&model.ArticleCategory{
		Fid:    in.Fid,
		Name:   in.Name,
		Remark: in.Remark,
		Sort:   in.Sort,
		Status: in.Status,
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

// ShowArticleCategoryByFID
//
// @Title ShowArticleCategoryByFID
// @Description: 按FID获取
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-12 00:05:02
// @receiver svc
// @param fid
// @return *system.SysResponse
func (svc *sArticle) ShowArticleCategoryByFID(fid uint32) *system.SysResponse {
	output, err := repositories.NewArticleCategoryRepository().GetListByFID(fid)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	var data []*dto.SelectOptionOutput
	for _, v := range output {
		d := dto.SelectOptionOutput{
			Key:   v.ID,
			Value: v.Name,
		}

		data = append(data, &d)
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}
