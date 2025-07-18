package dto

import (
	"ice_flame_gin/internal/app/models/association"
	"ice_flame_gin/internal/app/models/model"
	"time"
)

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

// ListArticleChannelInput
//
// @Description: 文章频道信息列表输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:01:28
type ListArticleChannelInput struct {
	Order  string
	Start  int
	Length int
}

// ListArticleChannelOutput
//
// @Description: 文章频道信息列表输出
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-20 01:02:08
type ListArticleChannelOutput struct {
	List  []*model.ArticleChannel
	Total int64
}

// ArticleTagInput
//
// @Description: 文章tag信息输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-22 23:53:34
type ArticleTagInput struct {
	Name   string
	Remark string
	Sort   uint32
	Status uint32
}

// ListArticleTagInput
//
// @Description: 文章tag信息列表输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-30 23:44:09
type ListArticleTagInput struct {
	Order  string
	Start  int
	Length int
}

// ListArticleTagOutput
//
// @Description: 文章tag信息列表输出
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-30 23:51:47
type ListArticleTagOutput struct {
	List  []*model.ArticleTag
	Total int64
}

type ArticleInput struct {
	CategoryId  uint32
	ChannelId   uint32
	Title       string
	Keywords    string
	Description string
	Content     string
	Link        string
	Author      string
	Tags        string
	PubDate     time.Time
	Thumbnail   string
	Summary     string
	Status      int
	Click       int
}

// ListArticleInput
//
// @Description: 文章信息列表输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2025-06-16 09:30:00
type ListArticleInput struct {
	Order    string
	Start    int
	Length   int
	Title    string // 支持按标题模糊搜索
	Category uint32 // 分类筛选
	Status   *int32 // 状态筛选，可为空
}

// ListArticleOutput
//
// @Description: 文章信息列表输出
// @Author liuxingyu <yuwen002@163.com>
// @Date 2025-06-16 09:30:00
type ListArticleOutput struct {
	List  []*association.Article
	Total int64
}
