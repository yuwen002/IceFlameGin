package dto

import "ice_flame_gin/internal/app/models/model"

// ArticleCategoryInput
//
// @Description: 文章分类信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-10 00:02:56
type ArticleCategoryInput struct {
	Fid    uint32
	Name   string
	Remark string
	Sort   uint32
	Status uint32
}

// ListArticleCategoryInput
//
// @Description: 文章分类信息列表写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 00:37:06
type ListArticleCategoryInput struct {
	Order  string
	Start  int
	Length int
}

// ListArticleCategoryOutput
//
// @Description: 文章分类信息列表输出
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-14 00:39:41
type ListArticleCategoryOutput struct {
	List  []*model.ArticleCategory
	Total int64
}

// ArticleChannelInput
//
// @Description: 文章频道信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:27:54
type ArticleChannelInput struct {
	Name   string
	Remark string
	Sort   uint32
	Status uint32
}
