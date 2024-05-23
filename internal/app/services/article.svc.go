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
		Status: &in.Status,
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
// @Description: 按FID获取文章分类信息
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

// ShowArticleCategory
//
// @Title ShowArticleCategory
// @Description: 文章分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:06:25
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sArticle) ShowArticleCategory(in dto.ListArticleCategoryInput) *system.SysResponse {
	out, err := repositories.NewArticleCategoryRepository().GetList(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	totalRecords, err := repositories.NewArticleCategoryRepository().CountRecords()
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
		Data: dto.ListArticleCategoryOutput{
			List:  out,
			Total: totalRecords,
		},
	}
}

// GetArticleCategoryByID
//
// @Title GetArticleCategoryByID
// @Description: 按ID查询文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 22:47:14
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sArticle) GetArticleCategoryByID(id uint32) *system.SysResponse {
	out, err := repositories.NewArticleCategoryRepository().GetByID(id)
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

// ChangeArticleCategoryByID
//
// @Title ChangeArticleCategoryByID
// @Description: 按ID更改文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 01:04:30
// @receiver svc
// @param id
// @param in
// @return *system.SysResponse
func (svc *sArticle) ChangeArticleCategoryByID(id uint32, in dto.ArticleCategoryInput) *system.SysResponse {
	err := repositories.NewArticleCategoryRepository().UpdateByID(id, &model.ArticleCategory{
		Fid:    in.Fid,
		Name:   in.Name,
		Remark: in.Remark,
		Sort:   in.Sort,
		Status: &in.Status,
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

// DeleteArticleCategoryByID
//
// @Title DeleteArticleCategoryByID
// @Description: 按ID删除文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-16 23:54:56
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sArticle) DeleteArticleCategoryByID(id uint32) *system.SysResponse {
	out, err := repositories.NewArticleCategoryRepository().GetByID(id)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if out == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "空记录",
			Data:    nil,
		}
	}

	if *out.Status == 1 {
		return &system.SysResponse{
			Code:    1,
			Message: "单页信息处于发布状态，不能删除",
			Data:    nil,
		}
	}

	//@todo 检查文章分类

	err = repositories.NewArticleCategoryRepository().DeleteByID(id)
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

// CreateArticleChannel
//
// @Title CreateArticleChannel
// @Description: 新建文章频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:08:10
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sArticle) CreateArticleChannel(in *dto.ArticleChannelInput) *system.SysResponse {
	err := repositories.NewArticleChannelRepository().Insert(&model.ArticleChannel{
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

// ShowArticleChannel
//
// @Title ShowArticleChannel
// @Description: 文章频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:20:54
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sArticle) ShowArticleChannel(in dto.ListArticleChannelInput) *system.SysResponse {
	out, err := repositories.NewArticleChannelRepository().GetList(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	totalRecords, err := repositories.NewArticleChannelRepository().CountRecords()
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
		Data: dto.ListArticleChannelOutput{
			List:  out,
			Total: totalRecords,
		},
	}
}

// GetArticleChannelByID
//
// @Title GetArticleChannelByID
// @Description: 按ID获取文章频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 22:36:08
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sArticle) GetArticleChannelByID(id uint32) *system.SysResponse {
	out, err := repositories.NewArticleChannelRepository().GetByID(id)
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

// ChangeArticleChannelByID
//
// @Title ChangeArticleChannelByID
// @Description: 按ID修改文章频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 17:07:24
// @receiver svc
// @param id
// @param in
// @return *system.SysResponse
func (svc *sArticle) ChangeArticleChannelByID(id uint32, in *dto.ArticleChannelInput) *system.SysResponse {
	err := repositories.NewArticleChannelRepository().UpdateByID(id, &model.ArticleChannel{
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

// DeleteArticleChannelByID
//
// @Title DeleteArticleChannelByID
// @Description: 按ID删除文章频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-21 22:42:55
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sArticle) DeleteArticleChannelByID(id uint32) *system.SysResponse {
	out, err := repositories.NewArticleChannelRepository().GetByID(id)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if out == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "空记录",
			Data:    nil,
		}
	}

	if out.Status == 1 {
		return &system.SysResponse{
			Code:    1,
			Message: "单页信息处于发布状态，不能删除",
			Data:    nil,
		}
	}

	//@todo 检查文章分类

	err = repositories.NewArticleChannelRepository().DeleteByID(id)
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

// GetArticleChannelAll
//
// @Title GetArticleChannelAll
// @Description: 获取所有频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-23 10:51:02
// @receiver svc
// @return *system.SysResponse
func (svc *sArticle) GetArticleChannelAll() *system.SysResponse {
	out, err := repositories.NewArticleChannelRepository().GetAll()
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
			Value: v.Name,
		}
		data = append(data, &d)
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    out,
	}
}

// CreateArticleTag
//
// @Title CreateArticleTag
// @Description: 新建文章tag信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-22 23:57:11
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sArticle) CreateArticleTag(in *dto.ArticleTagInput) *system.SysResponse {
	err := repositories.NewArticleTagRepository().Insert(&model.ArticleTag{
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

// ShowArticleTag
//
// @Title ShowArticleTag
// @Description: 文章tag信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-30 23:52:13
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sArticle) ShowArticleTag(in dto.ListArticleTagInput) *system.SysResponse {
	out, err := repositories.NewArticleTagRepository().GetList(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	totalRecords, err := repositories.NewArticleTagRepository().CountRecords()
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
		Data: dto.ListArticleTagOutput{
			List:  out,
			Total: totalRecords,
		},
	}
}

// GetArticleTagByID
//
// @Title GetArticleTagByID
// @Description: 按ID获取文章标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-02 00:03:08
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sArticle) GetArticleTagByID(id uint32) *system.SysResponse {
	out, err := repositories.NewArticleTagRepository().GetByID(id)
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

// ChangeArticleTagByID
//
// @Title ChangeArticleTagByID
// @Description: 按ID修改文章标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-02 00:15:45
// @receiver svc
// @param id
// @param in
// @return *system.SysResponse
func (svc *sArticle) ChangeArticleTagByID(id uint32, in *dto.ArticleTagInput) *system.SysResponse {
	err := repositories.NewArticleTagRepository().UpdateByID(id, &model.ArticleTag{
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

// DeleteArticleTagByID
//
// @Title DeleteArticleTagByID
// @Description: 按ID删除文章标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-11 17:55:14
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sArticle) DeleteArticleTagByID(id uint32) *system.SysResponse {
	out, err := repositories.NewArticleTagRepository().GetByID(id)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if out == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "空记录",
			Data:    nil,
		}
	}

	if out.Status == 1 {
		return &system.SysResponse{
			Code:    1,
			Message: "标签信息处于发布状态，不能删除",
			Data:    nil,
		}
	}

	//@todo 检查文章标签

	err = repositories.NewArticleTagRepository().DeleteByID(id)
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

// GetArticleTagAll
//
// @Title GetArticleTagAll
// @Description: 获取所有标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-05-23 14:50:41
// @receiver svc
// @return *system.SysResponse
func (svc *sArticle) GetArticleTagAll() *system.SysResponse {
	out, err := repositories.NewArticleTagRepository().GetAll()
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
			Value: v.Name,
		}
		data = append(data, &d)
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    out,
	}
}
